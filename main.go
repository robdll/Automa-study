package main

import (
	"bufio"
	"flag"
	"os"
	. "progetto-algoritmi/entita"
	"strings"
)

func main() {

	silent := flag.Bool("silent", false, "Suppress most standard output")
	flag.Parse()
	Silent = *silent

	p := Crea()
	scanner := bufio.NewScanner(os.Stdin)

	ConditionalOutput("## Pronto per la ricezione comandi. ('h' per la lista comandi, 'f' per terminare).")
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if line == "f" {
			ConditionalOutput("Esecuzione Terminata.")
			break
		}

		err := Esegui(p, line)
		if err != nil {
			ConditionalOutput(" ", err)
		}

		ConditionalOutput("## Attesa comando")
	}

}
