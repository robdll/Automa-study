package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

func TestAggiungiOstacolo(t *testing.T) {
	p := Crea()

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
	p := Crea()

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


func TestOstacoliSovrapposti(t *testing.T) {
	p := Crea()

	// Aggiungi un primo ostacolo
	err := p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{1, 1},
			AngoloAltoDestro:    [2]int{3, 3},
	})
	if err != nil {
			t.Fatalf("Errore nell'aggiungere il primo ostacolo: %v", err)
	}

	// Aggiungi un secondo ostacolo che si sovrappone al primo
	err = p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{2, 2},
			AngoloAltoDestro:    [2]int{4, 4},
	})
	if err != nil {
			t.Fatalf("Errore nell'aggiungere il secondo ostacolo: %v", err)
	}

	// Stampa il piano per verifica manuale
	p.Stampa()

	// Verifica che le posizioni comuni siano ancora rappresentate come ostacoli
	for x := 2; x <= 3; x++ {
			for y := 2; y <= 3; y++ {
					if _, ok := p.Mappa[[2]int{x, y}].(*Rettangolo); !ok {
							t.Errorf("Posizione (%d, %d) non rappresenta un ostacolo sovrapposto", x, y)
					}
			}
	}
}
