---
sidebar_position: 2
sidebar_label: História e propósito
---

# História e propósito do Go

## História

Go, também conhecido como Golang, começou a ser desenvolvida em 2007 pelos engenheiros `Robert Griesemer`, `Rob Pike` e `Ken Thompson`, tendo sido oficialmente anuncidada pela Google em 2009. A linguagem surgiu para atender às necessidades de `desenvolvimento em larga escala`, com `foco em sistemas distribuídos`, `infraestrutura de servidores` e `ferramentas modernas`. Inspirada por C, Pascal e outras linguagens, `Go combina simplicidade com eficiência`.

### Timeline

A jornada do Go, desde sua concepção no Google até se tornar uma das linguagens mais populares para o desenvolvimento de sistemas concorrentes e de alta performance, é marcada por lançamentos estratégicos e a introdução de funcionalidades cruciais.

#### Os primórdios (2007-2009)

- **2007:** Em um cenário de crescente complexidade de software no Google, os engenheiros Robert Griesemer, Rob Pike e Ken Thompson iniciam o esboço de uma nova linguagem. O objetivo era criar uma linguagem que combinasse a performance de linguagens compiladas como C++ com a simplicidade e a produtividade de linguagens dinâmicas, além de facilitar a programação concorrente.
- **2009:** A linguagem Go é oficialmente anunciada e seu código-fonte é aberto para a comunidade. Esta data é frequentemente citada como a "criação" oficial do Go, marcando o início de sua vida como um projeto de código aberto.

#### A consolidação e a promessa de estabilidade (2010-2012)

- **2010:** O Google realiza o primeiro lançamento oficial e estável da linguagem, disponibilizando-a para um público mais amplo e incentivando sua adoção inicial.
- **2012:** O lançamento do **Go 1.0** representa um marco fundamental. Com ele, a equipe de desenvolvimento assume um compromisso com a compatibilidade retroativa, garantindo que programas escritos em Go 1.0 continuariam a compilar e rodar com futuras versões da linguagem (Go 1.x). Esta versão trouxe estabilidade e um conjunto de ferramentas aprimoradas que solidificaram sua base de usuários.

#### Evolução e amadurecimento (2015-2021)

- **2015 - Go 1.5:** Ao contrário do que o resumo inicial indicava, o Go 1.5 **não introduziu Generics**. Esta versão foi notável por reescrever o compilador e o runtime inteiramente em Go (anteriormente eram em C), um marco de autossuficiência para o projeto. Também introduziu o coletor de lixo (garbage collector) concorrente de baixa latência.
- **2018 - Go 1.11:** Esta versão introduziu o suporte experimental para **Módulos (Modules)**, o sistema de gerenciamento de dependências que se tornaria o padrão, resolvendo um dos maiores desafios do ecossistema Go até então.

#### A era dos generics e melhorias contínuas (2022-2025)

- **2022 - Go 1.18:** Um dos lançamentos mais aguardados, o **Go 1.18 introduziu o suporte a Generics (parâmetros de tipo)**. Esta funcionalidade, solicitada pela comunidade por anos, permitiu a escrita de código mais flexível e reutilizável sem sacrificar a performance e a segurança de tipos.
- **2022 - Go 1.19 (Agosto):** Focou em refinar a implementação dos Generics e trouxe melhorias de segurança e desempenho.
- **2023 - Go 1.20 (Fevereiro):** Apresentou melhorias no compilador, otimizações de Profile-Guided Optimization (PGO) e a adição da função `errors.Join` para envolver múltiplos erros.
- **2023 - Go 1.21 (Agosto):** Continuou a aprimorar o desempenho, incluiu novas funções built-in (`min`, `max`, `clear`), e melhorou a compatibilidade com versões anteriores através da diretiva `GODEBUG`. O suporte a WebAssembly (Wasm) também foi aprimorado.
- **2024 - Go 1.22 (Fevereiro):** Introduziu uma mudança significativa no comportamento das variáveis de laços `for`, corrigindo uma fonte comum de bugs para iniciantes. Também adicionou a capacidade de iterar sobre inteiros em laços `range`.
- **2024 - Go 1.23 (Agosto):** Oficializou o recurso de "range over function", permitindo a iteração sobre funções, e introduziu um sistema de telemetria opcional para coletar dados de uso e ajudar a guiar o desenvolvimento futuro da linguagem.
- **2025 - Go 1.24 (Fevereiro):** Trouxe melhorias no ferramental, com um novo mecanismo para rastrear dependências de ferramentas via `go.mod`. Adicionou um novo analisador no `go vet` para detectar erros comuns em testes e expandiu a biblioteca padrão com pacotes de criptografia e um novo tipo `os.Root` para operações de sistema de arquivos isoladas.

## Propósito

Go foi projetado para:

- **Simplicidade**: Sintaxe clara e minimalista, reduzindo a complexidade
- **Performance**: Compilação rápida para binários nativos, com desempenho próximo ao de C/C++
- **Concorrência**: Suporte nativo para programação concorrente via `goroutines` e `channels`
- **Produtividade**: Ferramentas integradas (formatação, testes, documentação) para agilizar o desenvolvimento

:::info Caso de uso
Go é amplamente usado em projetos como Kubernetes, Docker e sistemas de alta performance no Google, onde a escalabilidade e a manutenção são críticas.
:::
