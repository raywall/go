---
sidebar_position: 1
sidebar_label: Web API em Go
---

# Web APIs com `net/http` e `Gin` em Go

<div className="row">
<div className="col">

Este módulo aborda a construção de `APIs RESTful` em Go, começando com o pacote padrão `net/http` e avançando para o framework `Gin`, que simplifica `roteamento`, `binding` e `validação`. O conteúdo é voltado para engenheiros Java, com comparações ao Spring Boot, e mantém-se objetivo com exemplos práticos e casos de uso para consulta futura.

## Conteúdo

- [Servidor HTTP com `net/http`](1-servidor-http.md)
- [Middlewares e handlers](2-middlewares-handlers.md)
- [Framework Gin: `roteamento`, `binding`, `validação`](3-framework-gin.md)
- [JSON, status codes e headers](4-json-status-headers.md)

📌 [Lab](5-laboratorio.md): Implementar uma API RESTful com Gin + validação.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-study.png').default} 
    style={{ transform:'scale(1.3)', marginTop:'-3rem' }}
    alt="A diaper brown gopher" />
</div>
</div>

O lab prático `implementa uma API RESTful` para o CRUD dos módulos anteriores, usando `Gin` com `validação de entrada`.
