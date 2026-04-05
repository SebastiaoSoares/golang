// 3. Conversor de Temperatura:
// Converta uma temperatura digitada em Celsius
// para Fahrenheit e Kelvin.

package exercicios

import "fmt"

func Ex3() {

	var temp_c, temp_f, temp_k float64

	fmt.Println("CONVERSOR DE TEMPERATURA: Informe a temperatura em graus Celsius")
	fmt.Scanln(&temp_c)

	temp_f = temp_c*1.8 + 32
	temp_k = temp_c + 273.15

	fmt.Printf(
		"Fahrenheit: %.2f\n"+
			"Kelvin: %.2f\n",
		temp_f, temp_k)

}
