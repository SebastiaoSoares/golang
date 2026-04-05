# Exercícios: Lógica de Programação em Go

Estes exercícios são iniciais, propostos por IA, focados em conceitos básicos e lógica de programação. O objetivo é acostumar com a sintaxe, aplicando conhecimentos adquiridos em estudos prévios.

## Fundamentos (Sintaxe, Variáveis e I/O)

1. Boas-vindas Interativo: Crie um programa que peça o nome e a idade do usuário no terminal e imprima uma mensagem formatada (ex: Olá, [Nome]! Você tem [Idade] anos.).

2. Calculadora Básica: Leia dois números reais (float64) e imprima a soma, subtração, multiplicação e divisão entre eles.

3. Conversor de Temperatura: Converta uma temperatura digitada em Celsius para Fahrenheit e Kelvin.

## Estruturas de Controle (Condicionais e Loops)

4. Par ou Ímpar: Leia um número inteiro e diga se ele é par ou ímpar.

5. Aprovação de Aluno: Leia 3 notas, calcule a média. Se for >= 7 (Aprovado), entre 5 e 6.9 (Recuperação) e < 5 (Reprovado). Use switch ou if/else.

6. Tabuada: Peça um número de 1 a 10 e imprima a tabuada dele usando um laço for.

7. Fibonacci: Imprima os primeiros N números da sequência de Fibonacci, onde N é informado pelo usuário.

## Estruturas de Dados (Slices e Maps)

8. Maior e Menor: Crie um slice com 10 números inteiros aleatórios. Itere sobre ele e encontre qual é o maior e o menor número.

9. Filtro de Pares: Receba um slice de números. Crie uma função que retorne um novo slice contendo apenas os números pares do slice original.

10. Contador de Palavras (Frequência): Receba uma frase (string). Use um map[string]int para contar quantas vezes cada palavra aparece na frase. (Dica: use strings.Fields()).

## Funções e Ponteiros

11. Troca de Valores (Swap): Crie uma função swap(a, b \*int) que receba dois ponteiros para inteiros e inverta os valores de a e b na memória. Imprima no main antes e depois de chamar a função para provar que mudou.

12. Busca Binária: Crie uma função que receba um slice de inteiros ordenados e um número alvo. Retorne a posição do número no slice (ou -1 se não existir) usando o algoritmo de busca binária.

13. Múltiplos Retornos: Crie uma função que receba um slice de floats e retorne três coisas de uma vez: a média, o valor máximo e o valor mínimo.

## Structs e Interfaces

14. Sistema de Banco: Crie um struct ContaBancaria com Titular e Saldo. Crie métodos (functions atrelados ao struct) para Depositar(valor) e Sacar(valor). O saque não pode deixar o saldo negativo (retorne um erro caso não haja fundos).

15. Formas Geométricas: \* Crie uma interface Forma com os métodos Area() float64 e Perimetro() float64.
    - Crie dois structs: Retangulo (base e altura) e Circulo (raio).

    - Implemente a interface para os dois structs e crie uma função ImprimirMedidas(f Forma) que funciona para ambos.

## Concorrência (Goroutines e Channels)

16. Cronômetros Concorrentes: Crie duas funções. Uma conta de 1 a 5 imprimindo "Processo A" a cada 500ms. A outra conta de 1 a 5 imprimindo "Processo B" a cada 800ms. Execute as duas ao mesmo tempo usando a palavra reservada go.

17. Ping Pong: Crie dois channels (chan string). Faça duas goroutines que ficam enviando "Ping" e "Pong" uma para a outra infinitamente através dos channels. Imprima o resultado no console com um time.Sleep para conseguir ler.

18. Worker Pool (Avançado): Você tem um slice com 100 números para calcular o quadrado. Crie 3 goroutines (workers) que ficam lendo esses números de um channel, calculando o quadrado e enviando para um channel de resultados.

## Projetos Práticos

19. **Gerenciador de Contatos com CSV (CLI)**

    Crie uma aplicação de linha de comando (CLI) simples para gerenciar uma agenda de contatos, no formato CSV.
    - O programa deve oferecer opções para adicionar um novo contato (Nome e Telefone) e listar todos os contatos salvos.

    - Os dados devem ser persistidos em um arquivo contatos.csv.

    _Dica: Pesquise sobre o pacote encoding/csv e o pacote os (especialmente as flags de abertura de arquivo como os.O_APPEND e os.O_CREATE para não sobrescrever os contatos anteriores ao adicionar um novo)._

20. **Buscador de CEP (Cliente HTTP)**

    Crie um programa que peça ao usuário para digitar um CEP no terminal. O programa deve fazer uma requisição GET para a API pública do ViaCEP (formato: https://viacep.com.br/ws/{cep}/json/).
    - O retorno da API será um JSON. Você deve decodificar esse JSON para uma struct em Go.

    - Por fim, imprima no terminal de forma formatada as informações do endereço (Logradouro, Bairro, Localidade e UF).

    _Dica: Utilize o pacote net/http (especificamente http.Get), encoding/json (para dar unmarshal na resposta) e io para ler o corpo da resposta._
