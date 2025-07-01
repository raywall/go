package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Estrutura do corpo do nosso request
type RequestBody struct {
	Code string `json:"code"`
}

// Estrutura da resposta da API do Go Playground
type PlaygroundResponse struct {
	Errors      string            `json:"Errors"`
	Events      []PlaygroundEvent `json:"Events"`
	Status      int               `json:"Status"`
	IsTest      bool              `json:"IsTest"`
	TestsFailed int               `json:"TestsFailed"`
}

type PlaygroundEvent struct {
	Message string `json:"Message"`
	Kind    string `json:"Kind"` // "stdout" ou "stderr"
	Delay   int    `json:"Delay"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
		"Content-Type":                 "application/json",
	}

	if request.RequestContext.HTTP.Method == "OPTIONS" {
		return events.APIGatewayV2HTTPResponse{StatusCode: 204, Headers: headers}, nil
	}

	var reqBody RequestBody
	if err := json.Unmarshal([]byte(request.Body), &reqBody); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Headers:    headers,
			Body:       `{"error": "invalid request body"}`,
		}, nil
	}

	// 1. Preparar a requisição para o Go Playground
	playgroundURL := "https://play.golang.org/compile"
	data := url.Values{}
	data.Set("version", "2")
	data.Set("body", reqBody.Code)

	// 2. Fazer a chamada HTTP para o Playground
	resp, err := http.Post(playgroundURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 502,
			Headers:    headers,
			Body:       `{"error": "failed to contact go playground"}`,
		}, nil
	}
	defer resp.Body.Close()

	// 3. Ler a resposta do Playground
	playgroundBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Headers:    headers,
			Body:       `{"error": "failed to read playground response"}`,
		}, nil
	}

	// 4. Processar a resposta do Playground para o nosso formato
	var playgroundResp PlaygroundResponse
	_ = json.Unmarshal(playgroundBody, &playgroundResp)

	finalOutput := make(map[string]string)
	if playgroundResp.Errors != "" {
		finalOutput["error"] = playgroundResp.Errors
	} else {
		var outputBuilder strings.Builder
		for _, event := range playgroundResp.Events {
			outputBuilder.WriteString(event.Message)
		}
		finalOutput["output"] = outputBuilder.String()
	}

	responseBody, _ := json.Marshal(finalOutput)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
