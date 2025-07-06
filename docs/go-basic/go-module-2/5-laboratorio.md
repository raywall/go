---
sidebar_position: 6
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
2. Crie um diretório `lab2`:

```bash
mkdir lab2
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab2
go mod init github.com/seu-usuario/lab2
```

</TabItem>
<TabItem value="app">

## Implementar as funções de um CRUD (Create, Read, Update e Delete) em memória

### Objetivo

Criar um CRUD capaz de gerenciar uma lista de produtos em memória, usando `slices`, `maps`, `structs` e `estruturas de controle`. O lab simula um CRUD (`Create`, `Read`, `Update`, `Delete`).

### Passo a passo

1. Crie um arquivo `crud.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod2/lab/crud.go" 
    allowExecute={true} 
    allowEdit={false} />

### Execute

```bash
go run crud.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione validação para impedir preços negativos em criarProduto e atualizarProduto
- Implemente uma função que liste produtos com preço acima de um valor específico
- Use um map para armazenar produtos por ID, em vez de um slice, e compare o desempenho

### Saída esperada

```bash
    Lista de produtos:
    ID: 1, Nome: Laptop, Preço: 999.99
    ID: 2, Nome: Mouse, Preço: 29.99

    Produto encontrado: {ID:1 Nome:Laptop Preco:999.99}
    Produto atualizado: {ID:1 Nome:Laptop Pro Preco:1299.99}
    Produto deletado com sucesso

    Lista final:
    ID: 1, Nome: Laptop Pro, Preço: 1299.99
```

:::info Caso de uso prático
Este lab simula uma API simples de gerenciamento de produtos, semelhante a um backend de e-commerce. Slices e maps são usados para manipulação de dados em memória, enquanto structs modelam entidades.
:::
</TabItem>
</Tabs>
