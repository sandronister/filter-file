# Filter File

Este projeto é uma aplicação em Go que permite buscar arquivos em um diretório com base em uma palavra-chave fornecida pelo usuário.

## Funcionalidades

- Solicita ao usuário um diretório inicial para listar arquivos.
- Permite ao usuário fornecer um novo diretório para busca.
- Realiza a busca de arquivos que contenham uma palavra-chave específica.
- Exibe mensagens de erro amigáveis em caso de falhas.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

filter-file/

    ├── cmd/ 
         └── main.go # Arquivo principal que executa a aplicação 

    ├── internal/ 
        ├── di/ # Contém a lógica de injeção de dependências 
        └── usecase/ # Contém os casos de uso da aplicação

    ├── pkg/ 
        ├── input_clear/ # Pacote para manipulação de entrada do usuário 
        └── fs_bussiness/ # Pacote para manipulação de arquivos PDF e diretórios

    └── go.mod # Arquivo de dependências do Go 
    ````
       
## Como Executar
    
1. Certifique-se de ter o Go instalado em sua máquina. Você pode verificar isso executando:
    
```go
go version
```

2. Clone este repositório:

````bash
git clone https://github.com/sandronister/filter-file.git
````

3. Execute o programa
```go
go run ./cmd/main.go <diretório-inicial>
```

Substitua <diretório-inicial> pelo caminho do diretório que deseja listar.

4. Siga as instruções no terminal para fornecer o novo diretório e a palavra-chave para busca.

###Exemplo de Uso

    $ go run ./cmd/main.go /home/user/documents
    Listing files in directory: /home/user/documents
    Enter the directory path: /home/user/projects
    Enter the keyword to search: report
    Searching for files with keyword: report

### Dependências
Este projeto utiliza os seguintes pacotes:

    github.com/sandronister/filter-file/internal/di: Gerencia a injeção de dependências.
    github.com/sandronister/filter-file/pkg/input_clear: Fornece utilitários para entrada de dados do usuário.

Certifique-se de instalar as dependências antes de executar o projeto:

```go
go mod tidy
```