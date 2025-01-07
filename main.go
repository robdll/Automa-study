package main

import (
	"bufio"
	"flag"
	"os"
	. "progetto-algoritmi/entita"
)

func main() {

	verbose := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()
	Silent = !*verbose

	p := NewPiano()
	scanner := bufio.NewScanner(os.Stdin)

	ConditionalOutput("## Pronto per la ricezione comandi. ('h' per la lista comandi, 'f' per terminare).")
	for scanner.Scan() {
		line := scanner.Text()

		err := Esegui(p, line)
		if err != nil {
			ConditionalOutput(" ", err)
		}

		ConditionalOutput("## Attesa comando")
	}

}


func esegui(p Piano, s string) {
	Esegui(p, s)
}

func newPiano() Piano {
	return NewPiano()
}