# GOCEPQUERY

This project is a simple program written in Go that performs ZIP Code (Postal Address Code) queries in two different APIs: ViaCep and BrasilAPI. The program returns the result of the first API it responds to.

## What I've learned

During the development of this project, I learned several important things about the Go language and about programming in general:
1. **Goroutines**: I learned how to use goroutines to run functions in parallel. This is useful for making multiple HTTP requests at the same time and improving the efficiency of the program.
2. **Channels**: I learned how to use channels for communication between goroutines. This allows goroutines to send and receive data in a secure and synchronized manner.
3. **Select**: I've learned how to use the `select` statement to handle multiple channels. This is useful for waiting until one of the channels receives a value and running the corresponding case.

## How to use

To use this program, you need to have Go installed on your computer. Once you've cloned the repository, you can run the program with the 'go run main.go' command. The program will ask you to enter the zip code you want to consult.

----

Este projeto é um programa simples escrito em Go que realiza consultas de CEP (Código de Endereçamento Postal) em duas APIs diferentes: ViaCep e BrasilAPI. O programa retorna o resultado da primeira API que responder.

## O que eu aprendi

Durante o desenvolvimento deste projeto, aprendi várias coisas importantes sobre a linguagem Go e sobre a programação em geral:
1. **Goroutines**: Aprendi como usar goroutines para executar funções em paralelo. Isso é útil para fazer várias requisições HTTP ao mesmo tempo e melhorar a eficiência do programa.
2. **Canais**: Aprendi a usar canais para comunicação entre goroutines. Isso permite que as goroutines enviem e recebam dados de forma segura e sincronizada.
3. **Select**: Aprendi a usar a declaração `select` para lidar com múltiplos canais. Isso é útil para esperar até que um dos canais receba um valor e executar o caso correspondente.

## Como usar

Para usar este programa, você precisa ter Go instalado em seu computador. Depois de clonar o repositório, você pode executar o programa com o comando `go run main.go`. O programa irá solicitar que você digite o CEP que deseja consultar.
