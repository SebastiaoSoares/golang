// 5. Aprovação de Aluno:
// Leia 3 notas, calcule a média.
// Se for >= 7 (Aprovado),
// entre 5 e 6.9 (Recuperação)
// e < 5 (Reprovado).
// Use switch ou if/else.

package exercicios

import "fmt"

func Ex5() {

	var (
		n1, n2, n3, media float64
		status            string
	)

	fmt.Println(
		"MÉDIA E STATUS ESCOLAR: Informe 3 notas separadas por vírgula:")
	fmt.Scanf("%f, %f, %f", &n1, &n2, &n3)

	media = (n1 + n2 + n3) / 3

	switch {
	case media < 5:
		status = "REPROVADO"
	case media < 7:
		status = "de RECUPERAÇÃO"
	default:
		status = "APROVADO"
	}

	fmt.Printf("\nA média é %f. Você está %s!", media, status)

}
