# Informações sobre o desafio 

Nesse desafio você deve criar um servidor que encurte urls e faça redirecionamento. 

## Como é o funcionamento

O usuário envia uma URL para o encurtador de URL, depois ele retorna um código de 6 dígitos que é um código que pode ser utilizado para receber a URL original. 

## Requisitos

- Criar um servidor http que contenha dois endopoins
- POST - recebe uma URL e retorna um código único
- GET/code - utiliza o code para redirecionamento para a url original 
- O code é um código único, a mesma url enviada várias vezes gera código diferentes 
- O code tem o tamanho de até 6 caracteres 

## Como rodar o projeto

1. Clonar o projeto
2. Rodar o seguinte comando para fazer o download das dependëncias do projeto
```
go mod tidy 
```
3. Rodar o projeto
```
go run cmd/encurtador-url/main.go
```