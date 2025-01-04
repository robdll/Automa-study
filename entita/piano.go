package entita

import (
	"fmt"
)

type Piano struct {
	Automi    map[string]*Automa
	Ostacoli  []Ostacolo
	Mappa     map[[2]int][]interface{}
}

func Crea() *Piano {
	fmt.Println("Creazione di un nuovo piano.")
	return &Piano{
		Automi:   make(map[string]*Automa),
		Ostacoli: []Ostacolo{},
		Mappa:    make(map[[2]int][]interface{}),
	}
}

func (p *Piano) Stato(x, y int) {
	fmt.Printf("Stato del piano in posizione: (%d, %d)\n", x, y)
	key := [2]int{x, y}
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		switch entities[0].(type) {
		case *Automa:
			fmt.Println("A")
		case *Ostacolo:
			fmt.Println("O")
		}
	} else {
		fmt.Println("E")
	}
}

func (p *Piano) Stampa(x, y int) {
	fmt.Println("Stampa elenco automi:")
	for _, automa := range p.Automi {
		automa.Stampa()
	}
	for _, ostacolo := range p.Ostacoli {
		ostacolo.Stampa()
	}
}

func (p *Piano) ListaAutomi() []*Automa {
	automi := make([]*Automa, 0, len(p.Automi))
	for _, automa := range p.Automi {
		automi = append(automi, automa)
	}
	return automi
}

// func (p *Piano) Richiamo(sorgente [2]int, segnale string) {
// 	// Mappa per tracciare gli automi che rispondono e la distanza minima
// 	automiRispondenti := make(map[string]int)
// 	distanzaMinima := -1

// 	fmt.Printf("Richiamo emesso dalla sorgente: %v con segnale: %s\n", sorgente, segnale)

// 	// Primo passaggio: trova la distanza minima tra gli automi che rispondono
// 	for nome, automa := range p.Automi {
// 			// Controlla se l'automa risponde al segnale
// 			if !strings.HasPrefix(automa.Nome, segnale) {
// 					fmt.Printf("Automa '%s' non risponde al segnale.\n", nome)
// 					continue
// 			}

// 			// Calcola la distanza Manhattan
// 			distanza := GetDistanzaManhattan(automa.Posizione, sorgente)
// 			fmt.Printf("Automa '%s' ha distanza %d dalla sorgente.\n", nome, distanza)

// 			// Verifica se esiste un percorso libero
// 			if !p.EsistePercorso(automa.Posizione, sorgente) {
// 					fmt.Printf("Automa '%s' non può raggiungere la sorgente: %v\n", nome, sorgente)
// 					continue
// 			}

// 			// Aggiorna la distanza minima e registra l'automa
// 			if distanzaMinima == -1 || distanza < distanzaMinima {
// 					distanzaMinima = distanza
// 					automiRispondenti = map[string]int{nome: distanza}
// 					fmt.Printf("Nuova distanza minima trovata: %d per automa '%s'.\n", distanza, nome)
// 			} else if distanza == distanzaMinima {
// 					automiRispondenti[nome] = distanza
// 					fmt.Printf("Automa '%s' aggiunto con distanza minima %d.\n", nome, distanza)
// 			}
// 	}

// 	fmt.Printf("Distanza minima: %d. Automi rispondenti: %v\n", distanzaMinima, automiRispondenti)

// 	// Secondo passaggio: sposta gli automi con distanza minima
// 	for nome := range automiRispondenti {
// 			if automiRispondenti[nome] == distanzaMinima {
// 					// Sposta l'automa
// 					err := p.MuoviAutoma(nome, sorgente)
// 					if err != nil {
// 							fmt.Printf("Errore nel movimento dell'automa '%s': %v\n", nome, err)
// 					} else {
// 							fmt.Printf("Automa '%s' si è spostato a %v\n", nome, sorgente)
// 							break // Solo un automa si deve spostare
// 					}
// 			}
// 	}
// }