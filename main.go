package main

import (
	"bufio"
	"os"
	. "progetto-algoritmi/entita"
)

func main() {

	p := NewPiano()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		Esegui(p, line)
	}

}

func esegui(p Piano, s string) {
	Esegui(p, s)
}

func newPiano() Piano {
	return NewPiano()
}