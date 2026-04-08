// 9. Filtro de Pares:
// Receba um slice_original de números. Crie uma função
// que retorne um novo slice_original contendo apenas
// os números pares do slice_original original.

package exercicios

import (
	"fmt"
)

func Ex9() {

	var (
		tam_slice, num int
		slice_original, slice_par []int
	)

	fmt.Printf("Tamanho do Slice: ")
	fmt.Scan(&tam_slice)

	for i := range tam_slice {
		fmt.Printf("Informe o %dº elemento: ", i+1)
		fmt.Scan(&num)
		slice_original = append(slice_original, num)
	}

	for _, n := range slice_original {
		if n % 2 == 0 {
			slice_par = append(slice_par, n)
		}
	}

	fmt.Printf("\nO Slice apenas com os números pares é o seguinte: ")
	fmt.Print(slice_par)

}
