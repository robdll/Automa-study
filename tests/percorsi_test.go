package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

func TestEsistePercorso(t *testing.T) {
	p := NuovoPiano()

	// Aggiungi un ostacolo
	p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{1, 1},
		AngoloAltoDestro:    [2]int{2, 2},
	})

	// Test percorso valido
	if !p.EsistePercorso([2]int{0, 0}, [2]int{3, 3}) {
		t.Errorf("Percorso valido non rilevato")
	}

	// Test percorso bloccato
	if p.EsistePercorso([2]int{1, 1}, [2]int{3, 3}) {
		t.Errorf("Percorso non valido erroneamente rilevato come valido")
	}

	// Test inizio o fine su un ostacolo
	if p.EsistePercorso([2]int{1, 1}, [2]int{0, 0}) {
		t.Errorf("Percorso con inizio su un ostacolo erroneamente rilevato come valido")
	}
}
