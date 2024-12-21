package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

func TestAggiungiAutoma(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un nuovo automa
	err := p.AggiungiAutoma(1, 1, "automa1")
	if err != nil {
		t.Errorf("Errore nell'aggiungere l'automa: %s", err)
	}

	// Verifica che l'automa sia nella posizione corretta
	automa, esiste := p.Automi["automa1"]
	if !esiste || automa.Posizione != [2]int{1, 1} {
		t.Errorf("Automa non trovato o posizione errata: %v", automa)
	}

	// Sposta l'automa
	err = p.AggiungiAutoma(2, 2, "automa1")
	if err != nil {
		t.Errorf("Errore nello spostare l'automa: %s", err)
	}

	// Verifica la nuova posizione
	automa, _ = p.Automi["automa1"]
	if automa.Posizione != [2]int{2, 2} {
		t.Errorf("Posizione errata dopo lo spostamento: %v", automa.Posizione)
	}
}

func TestRimuoviAutoma(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un automa
	p.AggiungiAutoma(1, 1, "automa1")

	// Rimuove l'automa
	err := p.RimuoviAutoma("automa1")
	if err != nil {
		t.Errorf("Errore nel rimuovere l'automa: %s", err)
	}

	// Verifica che l'automa sia stato rimosso
	if _, esiste := p.Automi["automa1"]; esiste {
		t.Errorf("Automa non rimosso correttamente")
	}
}
