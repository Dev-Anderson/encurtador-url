# Informações sobre o desafio 

Nesse desafio você deve criar um servidor que encurte urls e faça redirecionamento. O usuário deverá repassar uma URL e depois você deve enviar um código para o usuário, com esse código o usuário deverá conseguir a URL original

## Como é o funcionamento

O usuário envia uma URL para o encurtador de URL, depois ele retorna um código de 6 dígitos que é um código que pode ser utilizado para receber a URL original. 

## Requisitos

- Criar um servidor http que contenha dois endopoins
- POST - recebe uma URL e retorna um código único
- GET/code - utiliza o code para redirecionamento para a url original 
- O código é um código único, a mesma url enviada várias vezes gera código diferentes 
- O código tem o tamanho de até 6 caracteres 

## Stacks utilizadas 
- Golang
- 