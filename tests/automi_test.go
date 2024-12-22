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

	// Verifica che l'automa sia stato aggiunto nella posizione corretta
	automa, esiste := p.Automi["automa1"]
	if !esiste || automa.Posizione != [2]int{1, 1} {
		t.Errorf("Automa non trovato o posizione errata: %v", automa)
	}

	// Tenta di aggiungere un automa con un nome duplicato
	err = p.AggiungiAutoma(2, 2, "automa1")
	if err == nil {
		t.Errorf("Errore previsto per nome duplicato non generato")
	}

	// Tenta di aggiungere un automa in una posizione occupata
	err = p.AggiungiAutoma(1, 1, "automa2")
	if err == nil {
		t.Errorf("Errore previsto per posizione occupata non generato")
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

func TestMuoviAutoma(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un automa
	err := p.AggiungiAutoma(1, 1, "automa1")
	if err != nil {
		t.Errorf("Errore nell'aggiungere l'automa: %s", err)
	}

	// Muove l'automa a una nuova posizione
	err = p.MuoviAutoma("automa1", [2]int{2, 2})
	if err != nil {
		t.Errorf("Errore nel muovere l'automa: %s", err)
	}

	// Verifica la nuova posizione dell'automa
	automa, esiste := p.Automi["automa1"]
	if !esiste || automa.Posizione != [2]int{2, 2} {
		t.Errorf("Posizione errata dopo il movimento: %v", automa)
	}

	// Tenta di muovere un automa inesistente
	err = p.MuoviAutoma("automa2", [2]int{3, 3})
	if err == nil {
		t.Errorf("Errore previsto per automa inesistente non generato")
	}

	// Tenta di muovere l'automa in una posizione occupata
	p.AggiungiAutoma(3, 3, "automa2")
	err = p.MuoviAutoma("automa1", [2]int{3, 3})
	if err == nil {
		t.Errorf("Errore previsto per posizione occupata non generato")
	}
}
