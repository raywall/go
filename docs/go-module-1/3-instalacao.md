---
sidebar_position: 4
sidebar_label: Instalação, workspace e módulos
---

# Instalação, workspace e `go mod`

<!-- INSTALAÇÃO -->
<div className="row">
<div className="col">
### Instalação

1. Baixe o Go em [go.dev](https://go.dev/dl/) para o seu sistema operacional
2. Instale seguindo as instruções específicas:

- **Linux ou MAC**: Descompacte o arquivo e adicione ao `$PATH`
- **Windows**: Execute o instalador e verifique com `go version`
</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer.png').default} 
    style={{ marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>
:::tip Nota
No sistema operacional `MacOS` também é possível instalar o Go utilizando o `brew install go`
:::

4. Verifique a instalação utilizando o comando:

```bash
go version
```

<br />

## Workspace

<!-- WORKSPACE -->
<div className="row">
<div className="col">
Go utiliza uma estrutura de diretórios para organizar projetos. A partir do Go 1.11, o Go Modules substituiu o antigo $GOPATH como padrão.

### Diretórios principais (opcional, com Go Modules)

- `src/`: Código-fonte
- `bin/`: Binários compilados
- `pkg/`: Pacotes compartilhados

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-programmer-2.png').default} 
    style={{ transform:'scale(1.1)', marginTop:'20px' }}
    alt="A blue Go gopher" />
</div>
</div>

:::tip Nota
Com Go Modules, você pode trabalhar em qualquer diretório
:::

<br />

## Go Modules

Go Modules gerenciam dependências de forma semelhante ao Maven em Java.

1. Inicialize um módulo:

```bash
go mod init github.com/meu-usuario-github/meu-projeto
```

2. Adicione dependências:

```bash
go get github.com/exemplo/pacote
```

3. O arquivo `go.mod` armazena as dependências e versões

### Exemplo de `go.mod`

```bash
module github.com/meu-usuario-github/meu-projeto

go 1.21

require github.com/exemplo/pacote v1.0.0
```

:::info Caso de uso
Go Modules facilita a manutenção de projetos grandes, similar ao gerenciamento de dependências em Java com Maven ou Gradle.
:::
