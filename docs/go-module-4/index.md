---
sidebar_position: 1
sidebar_label: Tratamento de erros
---

# Tratamento de erros em Go

<div className="row">
<div className="col">

Este módulo explora a abordagem de tratamento de erros em Go, que difere significativamente do modelo de exceções usado em Java. O foco está na filosofia de erros explícitos, no uso do tipo error, técnicas de `wrapping`/`unwrapping` e `logging estruturado`. O conteúdo é voltado para engenheiros Java, com exemplos práticos e casos de uso objetivos para consulta futura.

O lab prático implementa funções com tratamento de erros e logging estruturado.

## Conteúdo

- [Filosofia do Go: erros explícitos](1-filosofia.md)
- [Padrão error, `errors.New`, `fmt.Errorf`](2-padrao-error.md)
- [Wrapping e unwrapping com `errors.Is`, `errors.As`](3-wrapping.md)
- [Pacote `log` e `log/slog`](4-logs.md)

📌 [Lab](5-laboratorio.md): Criar funções com tratamento de erros e logging estruturado.

</div>
<div className="col col--4 text--center">
<img 
    src={require('@site/static/img/gophers/gopher-coffee.png').default} 
    style={{ transform:'scale(0.8)', marginTop:'-65px' }}
    alt="A diaper brown gopher" />
</div>
</div>
