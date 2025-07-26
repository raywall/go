---
sidebar_position: 1
sidebar_label: Testes e qualidade de código
---

# Testes e qualidade de código em Go

<div className="row">
<div className="col">

Este módulo aborda testes e qualidade de código em Go, incluindo `testes unitários` e `testes de integração` com o pacote `testing`, uso de `mocks` com bibliotecas como `testify`, `benchmarks`, `profiling` e ferramentas de `análise estática`. O conteúdo é voltado para engenheiros Java, comparando com práticas como `JUnit` e `Mockito`, e mantendo-se objetivo para consulta futura.

## Conteúdo

- [Testes com testing](./1-testing.md)
- [Testes de unidade e integração](./2-testes-unitarios.md)
- [Testes com mocks (`testify`, `gomock`)](./3-testes-mock.md)
- [Benchmarks e profiling](./4-benchmarks-profiling.md)
- [Ferramentas: `go vet`, `golint`, `staticcheck`](./5-ferramentas.md)

📌 [Lab](./6-laboratorio.md): Criar testes unitários e de integração para o CRUD com cobertura de erro.

</div>
<div className="col col--4 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-schema.png').default} 
    style={{ transform:'scale(0.8)', marginTop:'-1.1rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

O lab prático implementa testes unitários e de integração para o CRUD dos módulos anteriores, com foco em `cobertura de erros`.
