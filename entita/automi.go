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
	fmt.Println("Posizionamento di un automa.")
	key := [2]int{x, y}

	// Controlla l'esistenza di un entità nella posizione (x, y)
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		
		// Se è un ostacolo, non fare nulla
		if _, ok := entities[0].(*Ostacolo); ok {
			fmt.Println("Ostacolo presente in quella posizione.")
			return
		} 

		// Cerca tra gli automi del piano se ne esiste uno con il nome fornito
		automi := p.ListaAutomi()
		for i := 0; i < len(automi); i++ {
			var target = entities[i].(*Automa)
			// Sposta l'automa se lo trovi
			if target.Nome == nome {
				target.Posizione = [2]int{x, y}
				fmt.Println("Spostamento effettuato.")
				return
			}
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
	fmt.Println("Creazione effettuata.")

}
