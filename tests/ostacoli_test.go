package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

func TestAggiungiOstacolo(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un Automa
	p.AggiungiAutoma(1, 1, "automa1")

	// Aggiunge un ostacolo dove c'è l'Automa
	err := p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{0, 0},
		AngoloAltoDestro:    [2]int{2, 2},
	})
	if err == nil {
		t.Errorf("Ostacolo sovrapposto aggiunto erroneamente")
	}

	// Aggiunge un altro ostacolo
	err = p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{3, 3},
		AngoloAltoDestro:    [2]int{4, 4},
	})
	if err != nil {
		t.Errorf("Errore nell'aggiungere ostacolo: %v", err)
	}
}

func TestRimuoviOstacolo(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un ostacolo
	ostacolo := Rettangolo{
		AngoloBassoSinistro: [2]int{3, 3},
		AngoloAltoDestro:    [2]int{4, 4},
	}
	p.AggiungiOstacolo(ostacolo)

	// Rimuove l'ostacolo
	err := p.RimuoviOstacolo(ostacolo)
	if err != nil {
		t.Errorf("Errore nel rimuovere ostacolo: %v", err)
	}

	// Verifica che tutte le posizioni siano state liberate
	for x := 3; x <= 4; x++ {
		for y := 3; y <= 4; y++ {
			if _, ok := p.Mappa[[2]int{x, y}].(*Rettangolo); ok {
				t.Errorf("Ostacolo non rimosso correttamente da (%d, %d)", x, y)
			}
		}
	}
}
