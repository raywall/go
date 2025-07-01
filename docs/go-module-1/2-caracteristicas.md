---
sidebar_position: 3
sidebar_label: Características da linguagem
---

# Características da linguagem

<!-- SIMPLICIDADE -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher.png').default} 
    style={{ transform:'scalex(-1) scale(0.7)', marginTop:'0px' }}
    alt="A blue Go gopher" />
</div>
<div className="col">

## Simplicidade

- Sintaxe reduzida, com poucas palavras-chave (25, contra 50+ em Java)
- Formatação automática com `go fmt`
- Ausência de recursos complexos como herança ou sobrecarga de métodos
</div>
</div>

<!-- PERFORMANCE -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-flash.png').default} 
    style={{ marginTop:'20px' }}
    alt="A gopher dressed in flash and running" />
</div>
<div className="col">

## Performance

- Compilação para binário nativo, eliminando a necessidade de uma máquina virtual (diferente de Java)
- Garbage collector otimizado para baixa latência
- Execução rápida, ideal para serviços de alta carga
</div>
</div>

<!-- CONCORRENCIA -->
<div className="row">
<div className="col col--3 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-goroutines.png').default} 
    style={{ transform:'scale(1)', marginTop:'5px' }}
    alt="A super strong gopher with a tattoo gourthines written on the chest" />
</div>
<div className="col">

## Concorrência

- **Goroutines**: Funções leves que executam concorrência de forma eficiente
- **Channels**: Mecanismo para comunicação segura entre goroutines
- Diferente do modelo de threads em Java, Go abstrai a complexidade do gerenciamento de concorrência
</div>
</div>

:::info Caso de uso
APIs RESTful e microsserviços, onde Go oferece alta performance e concorrência para lidar com múltiplas requisições simultâneas.
:::
