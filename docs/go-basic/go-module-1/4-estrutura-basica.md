---
sidebar_position: 5
sidebar_label: Estrutura básica de um programa
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Estrutura básica de um programa

Um programa Go segue uma estrutura simples, com o pacote main como ponto de entrada.

## Exemplo

<div className="row">
<div className="col">

<InteractiveCodeSnippet 
    src="code/mod1/lab/hello-world.go" 
    allowExecute={true} 
    allowEdit={false} />

<br />

</div>
<div className="col col--4 text--center">
<img
    src={require('@site/static/img/gophers/gopher-hello-world.png').default}
    style={{ transform:'scale(1)', marginTop:'-70px' }}
    alt="A blue Go gopher" />
</div>
</div>

- **package main**: Define o pacote principal, que gera um executável.
- **import "fmt"**: Importa o pacote fmt para formatação e saída.
- **func main()**: Função de entrada, equivalente ao public static void main em Java.

<br />

## Compilação e execução

```bash
go run hello.go     # Executa diretamente
go build hello.go   # Compila para um binário
```

<br />

## Comparação com Java

- Em Java, classes e métodos estáticos são obrigatórios. Em Go, a função main é suficiente.
- Go não usa ponto e vírgula (;) ao final das linhas.
