package main

import (
	"bufio"
	"fmt"
	"os"
	. "progetto-algoritmi/entita"
	"strings"
)

func main() {
	p := Crea()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("## Pronto per la ricezione comandi. ('h' per la lista comandi, 'f' per terminare).")
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if line == "f" {
			fmt.Println("Esecuzione Terminata.")
			break
		}

		err := Esegui(p, line)
		if err != nil {
			fmt.Println(" ", err)
		}

		fmt.Println("## Attesa comando")
	}

}