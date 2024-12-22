package main

import (
	"fmt"
	"progetto-algoritmi/entita"
)

func main() {
	// Crea un nuovo piano
	p := entita.NuovoPiano()

	// Aggiunge nuovi automi
	err := p.AggiungiAutoma(0, 0, "automa1")
	if err != nil {
		fmt.Println(err)
	}

	err = p.AggiungiAutoma(-1, -1, "automa2")
	if err != nil {
		fmt.Println(err)
	}

	// Tenta di aggiungere un automa con un nome duplicato
	err = p.AggiungiAutoma(1, 1, "automa1")
	if err != nil {
		fmt.Println(err)
	}

	// Muove un automa esistente
	err = p.MuoviAutoma("automa1", [2]int{2, 2})
	if err != nil {
		fmt.Println(err)
	}

	// Tenta di muovere un automa inesistente
	err = p.MuoviAutoma("automa3", [2]int{3, 3})
	if err != nil {
		fmt.Println(err)
	}

	// Elenca gli automi
	p.ElencaAutomi()

	// Stampa il piano graficamente
	p.StampaGrafica()
}
