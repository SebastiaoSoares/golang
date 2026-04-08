// 9. Filtro de Pares:
// Receba um slice de números. Crie uma função
// que retorne um novo slice contendo apenas
// os números pares do slice original.

package exercicios

import "fmt"

func Ex9() {

	var (
		tamSlice, num int
		slice []int
	)

	fmt.Printf("Tamanho do Slice: ")
	fmt.Scan(&tamSlice)

	for i := range tamSlice {
		fmt.Printf("Informe o %dº elemento: ", i+1)
		fmt.Scan(&num)
		slice = append(slice, num)
	}

	fmt.Printf("\nO Slice apenas com os números pares é o seguinte: ")
	fmt.Print(coletaPares(slice))

}

func coletaPares(sliceOriginal []int) []int {
	var slicePar []int

	for _, n := range sliceOriginal {
		if n % 2 == 0 {
			slicePar = append(slicePar, n)
		}
	}

	return slicePar
}
