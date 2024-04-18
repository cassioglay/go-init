package main

import (
	"fmt"
)

func main() {

	exibeIntroducao()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Print("Monitorando...")
	case 2:
		fmt.Print("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa.")
	default:
		fmt.Print("Não conheço este comando!")
	}

}

func exibeIntroducao() {
	nome := "Cássio Glay"
	versao := 1.1
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}
