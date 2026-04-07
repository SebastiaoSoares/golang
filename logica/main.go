package main

import (
	"fmt"
	"hello-go/exercicios"
)

func main() {

	var ex int

	fmt.Print("Escolha o exercício a ser executado (de 1 a 8): ")
	fmt.Scan(&ex)

	switch ex {
	case 1:
		exercicios.Ex1()
	case 2:
		exercicios.Ex2()
	case 3:
		exercicios.Ex3()
	case 4:
		exercicios.Ex4()
	case 5:
		exercicios.Ex5()
	case 6:
		exercicios.Ex6()
	case 7:
		exercicios.Ex7()
	case 8:
		exercicios.Ex8()

	default:
		fmt.Print("Esse ainda não existe.")
	}
}
