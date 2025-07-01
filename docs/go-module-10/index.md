---
sidebar_position: 1
sidebar_label: Módulo 10
---

# Deploy, observabilidade e boas ppráticas em Go

<div className="row">
<div className="col">

Este módulo aborda o _deploy de aplicações Go_, `observabilidade com logging`, `tracing` e `métricas`, e `boas práticas` para produção. Focado em engenheiros Java, compara com práticas do Spring Boot e ferramentas como `Prometheus`, com exemplos práticos e objetivos para consulta futura.

O lab prático containeriza a API CRUD dos módulos anteriores, configurando logging estruturado, tracing com `OpenTelemetry`, e exposição de métricas.

## Conteúdo

- [Build com `go build`, cross-compilation](1-build.md)
- [Docker com Go](2-docker.md)
- [Logging estruturado (`slog`, `zap`)](3-logging-estruturado.md)
- [Tracing com OpenTelemetry](4-tracing-otel.md)
- [`Linter`, cobertura, documentação automática com `godoc`](5-linter-cobertura-doc.md)

📌 [Lab](6-laboratorio.md): Containerizar o serviço e expor métricas/trace/logs.

</div>
<div className="col col--3 text--left">
<img 
    src={require('@site/static/img/gophers/gopher-inspect.png').default} 
    style={{ transform:'scalex(-1) scale(0.9)', marginTop:'-1rem' }}
    alt="A diaper brown gopher" />
</div>
</div>
