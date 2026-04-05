// 1. Boas-vindas Interativo:
// Crie um programa que peça o nome e a idade do usuário
// no terminal e imprima uma mensagem formatada
// (ex: Olá, [Nome]! Você tem [Idade] anos.).

package exercicios

import "fmt"

func Ex1() {
	var nome string
	var idade int

	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)

	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)

	fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
}
