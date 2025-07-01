---
sidebar_position: 7
sidebar_label: Laboratório
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Laboratório

<div className="text--right" style={{ background:'#6eb6e6', borderBottom:'3px solid #007d9c', marginTop:'2rem', marginBottom:'5rem' }}>
<img src={require('@site/static/img/gophers/gopher-background.png').default} style={{ width:'20rem',padding:'10px 0' }} alt="" />
</div>

<Tabs
className="lab-tabs"
defaultValue="config"
values={[
{label: 'Configuração', value: 'config'},
{label: 'Exercício 1', value: 'hello'},
{label: 'Exercício 2', value: 'vars'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

## Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab1`:

```bash
mkdir lab1
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab1
go mod init github.com/seu-usuario/lab1
```

</TabItem>
<TabItem value="hello">

## Hello, World!

### Objetivo

Criar um programa simples em Go para praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

### Passo a passo

1. Crie o arquivo `hello.go` com o seguinte conteúdo:

<InteractiveCodeSnippet 
    src="code/mod1/lab/hello-world.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Execute

```bash
go run hello.go
```

</TabItem>
<TabItem value="vars">

## Manipulação de variáveis e tipos

### Objetivo

Criar um programa simples em Go para praticar a configuração do ambiente, estrutura básica, tipos, variáveis e funções.

### Passo a passo

1. Crie um arquivo `vars.go` com o seguinte conteúdo:

<InteractiveCodeSnippet 
    src="code/mod1/lab/vars.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Execute

```bash
    go run vars.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Modifique `vars.go` para incluir uma função que calcula o dobro de um número
- Adicione uma constante para o IVA (Imposto sobre Valor Agregado, ex: 0.21) e calcule o preço final de um produto

### Saída esperada

```bash
Nome: Raywall, Idade: 45, Versão: 1.0.0
Preço: 99.99, Ativo: true
Troca: Java, Go
```

:::info Caso de uso prático
Este lab simula a manipulação de dados em um sistema simples, como um cadastro de produtos, onde variáveis e constantes são usadas para armazenar informações fixas (versão) e dinâmicas (preço, status).
:::
</TabItem>
</Tabs>
