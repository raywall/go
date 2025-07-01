---
sidebar_position: 1
sidebar_label: Pacotes, módulos e organização
---

# Pacotes, módulos e organização do código em Go

<div className="row">
<div className="col">

Este módulo explora a organização de código em Go, incluindo a estrutura idiomática de pacotes, convenções de projeto, gerenciamento de dependências com go mod, e boas práticas para engenheiros Java que estão aprendendo Go. O conteúdo é detalhado, mas objetivo, com exemplos e casos de uso para consulta futura.

O lab prático reorganiza o CRUD dos módulos anteriores em múltiplos pacotes, utilizando go mod para gerenciar dependências.

## Conteúdo

- [Estrutura de pacotes idiomática](1-estrutura-pacotes.md)
- [Convenções de projeto (`cmd`, `internal`, `pkg`)](2-convencoes-projeto.md)
- [`go mod` e versionamento](3-modulos-versoes.md)
- [Gerenciamento de dependências com `go get`, `replace`](4-gestao-dependencias.md)

📌 [Lab](5-laboratorio.md): Organizar o projeto CRUD em múltiplos pacotes com `go mod`.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-dependencies.png').default} 
    style={{ transform:'scale(0.65)', marginTop:'-65px' }}
    alt="A diaper brown gopher" />
</div>
</div>
