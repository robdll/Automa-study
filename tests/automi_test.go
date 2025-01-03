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

func TestRichiamo(t *testing.T) {
	p := NuovoPiano()

	// Aggiungi automi
	p.AggiungiAutoma(0, 0, "1011")
	p.AggiungiAutoma(1, 1, "1010")
	p.AggiungiAutoma(6, 6, "1111")

	// Aggiungi un ostacolo ora che la posizione (3,3) è libera
	err := p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{2, 2},
		AngoloAltoDestro:    [2]int{4, 4},
	})
	if err != nil {
			t.Fatalf("Errore nell'aggiungere l'ostacolo: %v", err)
	}

	// Stampa il piano per verificare lo stato
	p.Stampa()

	// Test del richiamo
	p.Richiamo([2]int{5, 5}, "101")

	// Verifica la posizione degli automi
	if p.Automi["1011"].Posizione != [2]int{5, 5} {
			t.Errorf("Automa '1011' non si è spostato correttamente verso la sorgente")
	}

	if p.Automi["1010"].Posizione == [2]int{5, 5} {
			t.Errorf("Automa '1010' non dovrebbe spostarsi verso una posizione già occupata")
	}

	// Verifica che l'automa '1111' sia rimasto nella sua nuova posizione
	if p.Automi["1111"].Posizione != [2]int{6, 6} {
			t.Errorf("Automa '1111' non dovrebbe rispondere al segnale e spostarsi")
	}

	p.Stampa()
}
