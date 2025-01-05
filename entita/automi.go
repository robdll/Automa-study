package entita

import (
	"fmt"
)

type Automa struct {
	Nome     string
	Posizione [2]int
}

func (a *Automa) Stampa() {
	fmt.Printf("%s: (%d, %d)\n", a.Nome, a.Posizione[0], a.Posizione[1])
}

func (p *Piano) PosizionaAutoma(x, y int, nome string) {
	key := [2]int{x, y}

	// Controlla l'esistenza di un Ostacolo nella posizione (x, y)
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		// Se è un ostacolo, non fare nulla
		if _, ok := entities[0].(*Ostacolo); ok {
			fmt.Println("Impossibile posizionare automa in quella posizione.")
			return
		} 
	}

	// Cerca tra gli automi del piano se ne esiste uno con il nome fornito
	automi := p.ListaAutomi()
	for i := 0; i < len(automi); i++ {
		// Sposta l'automa se lo trovi
		if automi[i].Nome == nome {
			automi[i].Posizione = [2]int{x, y}
			fmt.Println("Automa spostato.")
			return
		}
	}

	// Crea un nuovo automa dopo aver inizializzato p.Mappa[key] se necessario
	if _, exists := p.Mappa[key]; !exists {
		p.Mappa[key] = []interface{}{}
	}
	newAutoma := &Automa{
		Nome:	nome,
		Posizione: key,
	}
	p.Automi[nome] = newAutoma
	p.Mappa[key] = append(p.Mappa[key], newAutoma)
	fmt.Println("Automa creato.")

}
