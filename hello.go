package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Print("Monitorando...")
	case 2:
		fmt.Print("Exibindo Logs...")
	case 0:
		os.Exit(0)
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

func exibeMenu() {
	fmt.Println("1 - Inicar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	return comando
}
