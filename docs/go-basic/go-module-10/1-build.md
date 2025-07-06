---
sidebar_position: 2
sidebar_label: Build e cross-compilation
---

import InteractiveCodeSnippet from '@theme/InteractiveCodeSnippet';

# Build com `go build` e cross-compilation

## `go build`

- Compila código Go em executáveis.
- Simples, com suporte a flags para otimização.

### Exemplo

<InteractiveCodeSnippet 
    src="code/mod10/build.go" 
    allowExecute={true} 
    allowEdit={false} />

### Executar

```bash
go build -o app
./app # Saída: Hello, World!
```

## Cross-compilation

- _Go suporta compilação para diferentes sistemas operacionais e arquiteturas_, como:

  - `Linux`
  - `Windows`
  - `ARM`

- Definido com variáveis de ambiente `GOOS` e `GOARCH`

### Exemplo

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux
GOOS=windows GOARCH=amd64 go build -o app-windows.exe
```

### Comparação com Java

#### Java

- Compila para bytecode (JVM), com dependências de runtime.

#### Go

- Gera **binários nativos**, _sem dependências externas_.

:::info Caso de uso
Cross-compilation é ideal para criar binários para containers ou dispositivos embarcados.
:::
