package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func setRandom() int64 {

	numeroMaximo := big.NewInt(100)
	numeroAleatorio, _ := rand.Int(rand.Reader, numeroMaximo)

	numeroAleatorio.Add(numeroAleatorio, big.NewInt(1))

	return numeroAleatorio.Int64()

}

func main() {

	fmt.Printf("Bem-vindo ao jogo da adivinhação!\n\n-------------------------\n\n")

	vezesJogadas_tentativas := (make(map[int]int))
	vezesJogadas := 1
	vezesJogadas_tentativas[vezesJogadas]++

	numeroAleatorioInt := setRandom()

	var valorDigitado int64

	for {

		fmt.Printf("Digite um número entre 1 e 100: ")
		fmt.Scanln(&valorDigitado)

		if valorDigitado > numeroAleatorioInt {

			vezesJogadas_tentativas[vezesJogadas]++
			fmt.Printf("O número é menor que %d.\n", valorDigitado)

		} else if valorDigitado < numeroAleatorioInt {

			vezesJogadas_tentativas[vezesJogadas]++
			fmt.Printf("O número é maior que %d.\n", valorDigitado)

		} else {

			fmt.Printf("Parabéns, você acertou!\nVocê deseja jogar novamente? (s/n): ")
			repeat := ""
			fmt.Scanln(&repeat)

			if repeat == "n" {

				for i := 1; i <= len(vezesJogadas_tentativas); i++ {

					fmt.Printf("Você utilizou %d tentativas na %dª jogada.\n", vezesJogadas_tentativas[i], i)

				}

				break

			} else {

				numeroAleatorioInt = setRandom()
				vezesJogadas++
				vezesJogadas_tentativas[vezesJogadas]++

			}

		}

	}

}
