## README.md - API REST de Conversão de Moedas

Descrição

Esta é uma API REST simples, desenvolvida em Go, que permite a conversão de valores entre diferentes moedas, com taxas de câmbio fixas. A API aceita requisições em JSON e retorna o valor convertido de acordo com as moedas fornecidas.

## Estrutura do Projeto

- main.go: Inicia o servidor e define o ponto de entrada para a aplicação
- controllers.go: Contém a lógica da conversão de moedas e o tratamento das requisições.
- routes.go: Define as rotas da API

## Como Usar

1.Clonar o Repositório

-git clone https://github.com/BodeXX/api-rest-go-git
-cd api-rest-go

2.Executar a API (Certifique-se de ter o Go instalado.)

-go run main.go

A API estará disponivel na porta 8080.

3.Testando a API

Usando o Postman

1.Abra o Postman.

2.Crie uma nova requisição do tipo POST.

3.Insira a URL: http://localhost:8080/convert.

4.No corpo da requisição, selecione JSON e insira o seguinte exemplo:

{
    "moedaOrigem": "USD",
    "moedaDestino": "BRL",
    "valor": 100
}

5.Envie a requisição.

Usando o Terminal com curl: Você também pode testar a API diretamente no terminal usando curl. Execute o seguinte comando:

-curl -X POST http://localhost:8080/convert -H "Content-Type: application/json" -d '{"moedaOrigem":"USD","moedaDestino":"BRL","valor":100}'

## Endpoints Disponiveis

POST /convert

Converte um valor entre duas moedas.

-Corpo da requisição (JSON):

{
    "moedaOrigem": "USD",
    "moedaDestino": "BRL",
    "valor": 100
}

-Resposta (JSON):
{
    "moedaOrigem": "USD",
    "moedaDestino": "BRL",
    "valorOriginal": 100,
    "valorConvertido": 546
}

## Moedas Suportadas

As seguintes moedas são suportadas:

-USD: Dólar americano
-EUR: Euro
-BRL: Real brasileiro
-JPY: iene japonês
-CHF: Franco suiço

## Validações

-O valor a ser convertido deve ser amior que zero.
-A moeda de origem e a moeda de destino devem ser válidas e suportadas.