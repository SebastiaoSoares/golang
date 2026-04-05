// 4. Par ou Ímpar:
// Leia um número inteiro e diga se ele
// é par ou ímpar.

package exercicios

import "fmt"

func Ex4() {

	var n int

	fmt.Print("VERIFICADOR DE PARIDADE: Informe um número inteiro: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("É Par!")
	} else {
		fmt.Println("É Ímpar!")
	}

}
