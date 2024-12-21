package main

import (
	"bufio"
	"os"
	"progetto-algoritmi/entita"
)

func main() {

	// Inizializza il piano
	p := entita.NuovoPiano()

	// Legge comandi dall'input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
			comando := scanner.Text()
			p.EseguiComando(comando)
	}

}