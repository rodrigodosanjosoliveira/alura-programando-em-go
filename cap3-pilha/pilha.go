package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha()
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")
	fmt.Println("Tamanho após empilhar 4 valores: ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando ", v)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia? ", pilha.Vazia())
	}

	_, err := 
}