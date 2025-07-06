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
{label: 'Exercício', value: 'app'},
{label: 'Tarefa', value: 'task'},
]}>
<TabItem value="config">

## Configuração

1. Instale o Go e verifique com `go version`
2. Crie um diretório `lab3`

```bash
mkdir lab3
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab3
go mod init github.com/seu-usuario/lab3
```

</TabItem>
<TabItem value="app">

## Refatorar o CRUD Usando Interfaces para Repositórios

### Objetivo

Refatorar o CRUD do Módulo 2 para usar uma interface de repositório, abstraindo o armazenamento de dados e facilitando testes e manutenção.

### Passo a passo

1. Crie um arquivo `crud.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod3/lab/crud.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Execução

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Crie um segundo repositório (ex: `RepositorioMock`) que implemente `RepositorioProdutos` para testes, retornando dados fixos.
- Adicione uma função anônima no main para filtrar produtos com preço acima de um valor
- Use um método em Produto para calcular o preço com imposto (ex: 21%)

### Saída esperada

```bash
Lista de produtos:
ID: 1, Nome: Laptop, Preço: 999.99
ID: 2, Nome: Mouse, Preço: 29.99

Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
Produto deletado com sucesso

Lista de produtos:
ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

:::info Caso de uso prático
A interface RepositorioProdutos abstrai o armazenamento, permitindo trocar a implementação (ex: de memória para banco de dados) sem alterar o código principal, similar ao padrão Repository em Java.
:::
</TabItem>
</Tabs>
<br />
