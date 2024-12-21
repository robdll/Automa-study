package main

import (
	"bufio"
	"os"
)

func main() {

	// Inizializza il piano
	p := piano.NuovoPiano()

	// Legge comandi dall'input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
			comando := scanner.Text()
			p.EseguiComando(comando)
	}

}