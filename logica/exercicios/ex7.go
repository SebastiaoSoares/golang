// 7. Fibonacci:
// Imprima os primeiros N números da sequência
// de Fibonacci, onde N é informado pelo usuário.

package exercicios

import "fmt"

func Ex7() {

	var n int
	fib1 := 1
	fib2 := 1

	fmt.Print("FIBONACCI: Informe a quantidade de números da sequência: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Nenhum número a ser exibido.")
		return
	}

	fmt.Printf("%d ", fib1)

	for i := 1; i < n; i++ {
		fmt.Printf("%d ", fib2)
		fib1, fib2 = fib2, fib1+fib2
	}

}
