// 2. Calculadora Básica:
// Leia dois números reais (float64) e imprima a soma,
// subtração, multiplicação e divisão entre eles.

package exercicios

import "fmt"

func Ex2() {

	var a, b float64

	fmt.Println("CALCULADORA: Informe dois valores separados por um espaço:")
	fmt.Scanf("%f %f", &a, &b)

	fmt.Printf(
		"Soma: %.2f\n"+
		"Substração: %.2f\n"+
		"Multiplicação: %.2f\n"+
		"Divisão: %.2f\n",
		a+b,
		a-b,
		a*b,
		a/b)
}
