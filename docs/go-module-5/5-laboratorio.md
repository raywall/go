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
2. Crie o diretório `lab5`:

```bash
mkdir lab5
```

3. Acesse o diretório e inicialize um módulo:

```bash
cd lab5
go mod init github.om/seu-usuario/lab5
```

</TabItem>

<TabItem value="app">

## Criar um worker pool para processamento concorrente de tarefas

### Objetivo

Implementar um worker pool que processa tarefas de cálculo (ex.: calcular o quadrado de números) usando goroutines e channels, com sincronização via `sync.WaitGroup` e `logging` estruturado.

### Passo a passo

1. Crie um arquivo `workerpool.go` com o seguinte código:

<InteractiveCodeSnippet 
    src="code/mod5/lab/workerpool.go" 
    allowExecute={true} 
    allowEdit={false} />

#### Execução

```bash
go run workerpool.go
```

</TabItem>
<TabItem value="task">

## Tarefa

- Adicione um `context.Context` para suportar cancelamento do worker pool após um timeout.
- Implemente um mecanismo para lidar com erros em uma tarefa (ex: valor inválido).
- Use `select` para processar resultados com um timeout por tarefa.

### Saída esperada

#### Console

```bash
Tarefa 1: Quadrado = 1
Tarefa 2: Quadrado = 4
Tarefa 3: Quadrado = 9
Tarefa 4: Quadrado = 16
Tarefa 5: Quadrado = 25
```

#### Logs JSON (exemplo)

```json
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Tarefa enviada","tarefa_id":1,"valor":1}
{"time":"2025-06-12T01:07:00Z","level":"INFO","msg":"Processando tarefa","worker":1,"tarefa_id":1,"valor":1}
...
{"time":"2025-06-12T01:07:01Z","level":"INFO","msg":"Processamento concluído","total_tarefas":5}
```

:::info Caso de uso prático
O `worker pool` simula cenários reais, como processamento de filas de mensagens (ex: `RabbitMQ`) ou tarefas em lote (ex: redimensionamento de imagens).
:::

</TabItem>
</Tabs>
