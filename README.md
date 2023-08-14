# Desafio Padawan Go

Este é um projeto de desafio que implementa uma API de conversão de moedas usando a linguagem de programação Go e o framework Gin.

#Para 

## Requisitos

Certifique-se de ter instalado o seguinte software:

- Go (1.16+)
- Docker para o banco MySQL 

## Tecnologias utilizadas 

- GORM para programação com banco de dados
- Gin-GONIC para Requisições HTTP 
- MySQL 
- Docker 
- Testes unitários

## Instalação

1. Clone o repositório para sua máquina:

   ```bash
    git clone https://github.com/seu-usuario/desafio-padawan-go.git
    ```
2. Baixe as dependencias do go 
   
    ```go 
    go mod tidy 
    ```
3. Crie a imagem do container MySQL
    ```bash
    docker-compose up --build
    ```

4. Para fechar o docker 
    ```bash 
    docker-compose down 
    ```


## Execução

Com o servidor docker MySQL iniciado navegue até a pasta abaixo 

```bash
lcslima45@LAPTOP-IR1IKESJ:/mnt/c/users/user/go/src/github.com/lcslima45$ cd desafio-padawan-go
lcslima45@LAPTOP-IR1IKESJ:/mnt/c/users/user/go/src/github.com/lcslima45/desafio-padawan-go$ go run main/main.go
```

## Rotas

A API possui a seguinte rota para conversão de moedas:

- GET /exchange/:amount/:from/:to/:rate: Converte um valor de uma moeda para outra com uma taxa de conversão específica.
Exemplo de uso:

## Exemplo de uso
Com o comando curl faça a conversão

```
curl http://localhost:8000/exchange/100/USD/EUR/1.2
```

Resposta esperada
```json
{
  "valorConvertido": "120.00",
  "simboloMoeda": "EUR"
}
```

## Testes

Adotamos testes unitários para a camada:
- de Handlers na pasta controllers
- de Serviço na pasta converter 
- de Repositorio na pasta repository
No terminal execute 

```go
go test ./...
```


### Descrição dos Diretórios

- `controllers/`: Contém os controladores responsáveis por manipular as requisições HTTP e chamar os serviços apropriados.
- `converter/`: Inclui os componentes relacionados à conversão de moedas, como o serviço de conversão.
- `models/`: Armazena as definições das estruturas de dados, como o modelo de conversão de moedas.
- `repository/`: Contém a camada de repositório, que lida com operações de banco de dados.
- `routes/`: Define as rotas da API e associa-as aos controladores correspondentes.
- `main/`: Contém o arquivo principal da aplicação, que configura o servidor e inicia a execução.
- `README.md`: Documentação do projeto, que fornece informações sobre como configurar, executar e testar a aplicação.

