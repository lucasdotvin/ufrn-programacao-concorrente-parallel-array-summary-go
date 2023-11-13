<h1 align="center">Parallel Array Summary - Go</h1>

<p align="center">
<a href="https://www.metropoledigital.ufrn.br/portal/"><img alt="UFRN - IMD" src="https://img.shields.io/badge/ufrn-imd-ufrn?style=for-the-badge&labelColor=%23164194&color=%230095DB&link=https%3A%2F%2Fwww.metropoledigital.ufrn.br%2Fportal%2F"></a>
<br>
<a href="https://github.com/gomods/athens"><img src="https://img.shields.io/github/go-mod/go-version/gomods/athens.svg" alt="GitHub go.mod Go version of a Go module"></a>
</p>

Projeto de demonstração de gerenciamento de concorrência ao lidar com leitura e escrita de um arquivo.

## Iniciando

Essas instruções lhe darão uma cópia do projeto e um caminho para executá-lo localmente para fins de desenvolvimento e teste.

### Pré-Requisitos

Você precisará do toolkit da linguagem Go em sua versão 1.21 instalado.

### Building

Para fazer build do projeto, basta executar:

```bash
make build
```

Isso criará um binário chamado `parallel-array-summary` na pasta `dist/`.

Para rodar os casos de teste pré-determinados, use:

```bash
make demonstrate
```

Os casos pré-determinados são:
1. `n = 5`, `t = 16`, `e = 0`
2. `n = 5`, `t = 16`, `e = 40`
3. `n = 5`, `t = 256`, `e = 40`
4. `n = 7`, `t = 256`, `e = 0`
5. `n = 7`, `t = 256`, `e = 40`

Os resultados serão salvos na pasta `resultados/`, enquanto os dados remanescentes serão salvos na pasta `data/`.

## Licença

Esse projeto é distribuído sob a Licença MIT. Leia o arquivo [LICENSE](LICENSE) para ter mais detalhes.
