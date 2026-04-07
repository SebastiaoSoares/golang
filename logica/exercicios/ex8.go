// 8. Maior e Menor:
// Crie um slice com 10 números inteiros aleatórios.
// Itere sobre ele e encontre qual é o maior e o
// menor número.

// Optei por limitar de 1 a 100 para conseguir enxergar as comparações.

package exercicios

import (
	"fmt"
	"math/rand"
)

func Ex8() {

	var (
		numeros []int
		maior, menor int
	)

	fmt.Print("Meu Slice de número aleatórios (de 1 a 100): ")

	for i := 0; i <= 10; i++ {
		rand := rand.Intn(100)+1
		numeros = append(numeros, rand)
		fmt.Printf("%d ", rand)
	}

	for i, n := range numeros {
		if i != 0 {
			if maior < n {
				maior = n
			}
			if menor > n {
				menor = n
			}
		} else {
			maior = n
			menor = n
		}
	}

	fmt.Printf("\nO maior deles é %d, enquato o menor é %d.", maior, menor)

}