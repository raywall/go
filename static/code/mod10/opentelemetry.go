package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	exporter, err := stdouttrace.New()
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)
	tracer := otel.Tracer("meu-app")
	_, span := tracer.Start(context.Background(), "operação-principal")
	defer span.End()

	log.Println("Operação rastreada")
}
