// 6. Tabuada:
// Peça um número de 1 a 10 e imprima
// a tabuada dele usando um laço for.

package exercicios

import "fmt"

func Ex6() {

	var n int

	for n < 1 || n > 10 {
		fmt.Print("TABUADA: Informe um número de 1 a 10: ")
		fmt.Scan(&n)
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d + %d = %d\n", n, i, n+i)
		if i == 10 {
			fmt.Print("\n")
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d - %d = %d\n", n, i, n-i)
		if i == 10 {
			fmt.Print("\n")
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
		if i == 10 {
			fmt.Print("\n")
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d / %d = %.2f\n", n, i, float64(n)/float64(i))
		if i == 10 {
			fmt.Print("\n")
		}
	}

}
