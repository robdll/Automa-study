package entita

import (
	"fmt"
)

type Automa struct {
	Nome     string
	Posizione [2]int
}

func (a *Automa) Stampa() {
	fmt.Printf("%s: %d,%d\n", a.Nome, a.Posizione[0], a.Posizione[1])
}

func (p *Piano) PosizionaAutoma(x, y string, nome string) {
	key := x + "_" + y

	// Controlla l'esistenza di un Ostacolo nella posizione (x, y)
	if p.isOstacolo(key) {
		return
	}

	// Inizializza p.Mappa[key] alla nuova posizione se necessario
	if _, exists := (*p.Mappa)[key]; !exists {
		(*p.Mappa)[key] = []interface{}{}
	}

	// Cerca tra gli automi del piano se ne esiste uno con il nome fornito
	automaToPlace, _ := p.OttieniAutoma(nome)
	if automaToPlace == nil {
		// Crea un nuovo Automa
		newAutoma := &Automa{
			Nome:	nome,
			Posizione: GetValuesFromKey(key),
		}
		// Aggiungi l'Automa alla mappa e alla lista degli automi
		(*p.Automi)[nome] = newAutoma
		(*p.Mappa)[key] = append((*p.Mappa)[key], newAutoma)
	} else {
		oldKey := GetKeyFromValues(automaToPlace.Posizione[0], automaToPlace.Posizione[1])
		for index, target := range (*p.Mappa)[oldKey] {
			if target == automaToPlace {
				// Rimuovi l'Automa dalla vecchia posizione
				(*p.Mappa)[oldKey] = append((*p.Mappa)[oldKey][:index], (*p.Mappa)[oldKey][index+1:]...)
				// Cancella la chiave se non ci sono più entità in quella posizione
				if len((*p.Mappa)[oldKey]) == 0 {
					delete((*p.Mappa), oldKey)
				}
				break
			}
		}
		// Aggiorna la posizione dell'Automa e la mappa
		automaToPlace.Posizione = GetValuesFromKey(key)
		(*p.Mappa)[key] = append((*p.Mappa)[key], automaToPlace)
	}
}
