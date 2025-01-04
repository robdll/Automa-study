package entita

import (
	"fmt"
	. "progetto-algoritmi/entita"
	"testing"
)

func TestAggiungiAutoma(t *testing.T) {
	p := Crea()
	
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
	p := Crea()
	
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
	p := Crea()

	// Aggiungi automi
	p.AggiungiAutoma(0, 0, "1011")        // Distanza: 10
	p.AggiungiAutoma(1, 1, "1010")        // Distanza: 8
	p.AggiungiAutoma(2, 2, "10100")       // Distanza: 8 (aggiunto per il nuovo caso)
	p.AggiungiAutoma(6, 6, "1111")        // Non risponde al segnale

	// Aggiungi un ostacolo
	err := p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{3, 3},
			AngoloAltoDestro:    [2]int{4, 4},
	})
	if err != nil {
			t.Fatalf("Errore nell'aggiungere l'ostacolo: %v", err)
	}

	// Stampa il piano prima del richiamo
	fmt.Println("Stato del piano prima del richiamo:")
	p.Stampa()

	// Effettua il richiamo
	p.Richiamo([2]int{5, 5}, "101")

	// Stampa il piano dopo il richiamo
	fmt.Println("Stato del piano dopo il richiamo:")
	p.Stampa()

	// Verifica che solo il primo automa con distanza 8 si sia spostato
	if p.Automi["1010"].Posizione != [2]int{5, 5} {
			t.Errorf("Automa '1010' non si è spostato correttamente verso la sorgente")
	}
	if p.Automi["10100"].Posizione == [2]int{5, 5} {
			t.Errorf("Automa '10100' non dovrebbe spostarsi verso una posizione già occupata")
	}

	// Verifica che gli altri automi non si siano mossi
	if p.Automi["1011"].Posizione != [2]int{0, 0} {
			t.Errorf("Automa '1011' non dovrebbe spostarsi")
	}
	if p.Automi["1111"].Posizione != [2]int{6, 6} {
			t.Errorf("Automa '1111' non dovrebbe rispondere al segnale e spostarsi")
	}
}

func TestExampleFromPDF(t *testing.T) {
	p := Crea()

	p.AggiungiAutoma(2, 1, "1")     // Automaton "1"
	p.AggiungiAutoma(2, 8, "10")    // Automaton "10"
	p.AggiungiAutoma(5, 11, "11")   // Automaton "11"
	p.AggiungiAutoma(8, 9, "100")   // Automaton "100"
	p.AggiungiAutoma(10, 6, "101")  // Automaton "101"
	p.AggiungiAutoma(13, 3, "110")  // Automaton "110"
	p.AggiungiAutoma(12, 2, "111")  // Automaton "111"

	// Add obstacles (with keyed fields)
	p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{3, 0},
			AngoloAltoDestro:    [2]int{15, 1},
	}) // obstacle1
	p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{6, 0},
			AngoloAltoDestro:    [2]int{10, 3},
	}) // obstacle2
	p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{3, 2},
			AngoloAltoDestro:    [2]int{4, 7},
	}) // obstacle3
	p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{1, 3},
			AngoloAltoDestro:    [2]int{4, 6},
	}) // obstacle4
	p.AggiungiOstacolo(Rettangolo{
			AngoloBassoSinistro: [2]int{6, 5},
			AngoloAltoDestro:    [2]int{7, 8},
	}) // obstacle5

	// Print the initial state of the plane
	fmt.Println("Stato del piano prima del richiamo:")
	p.Stampa()

	// Emit the signal
	p.Richiamo([2]int{5, 3}, "1")

	// Print the state of the plane after the signal
	fmt.Println("Stato del piano dopo il richiamo:")
	p.Stampa()

	// Expected positions after the signal
	expectedPositions := map[string][2]int{
			"1":    {2, 1},  // Can't reach the signal
			"10":   {2, 8},  // Stays because another automaton occupies (5, 3)
			"11":   {5, 11}, // Can't reach the signal
			"100":  {8, 9},  // Can't reach the signal
			"101":  {5, 3},  // Moves to (5, 3) because it has the minimum distance
			"110":  {13, 3}, // Can't reach the signal
			"111":  {12, 2}, // Can't reach the signal
	}

	// Verify the positions of automata
	for nome, posizioneAttesa := range expectedPositions {
			automa, esiste := p.Automi[nome]
			if !esiste {
					t.Errorf("Automa '%s' non trovato.", nome)
					continue
			}
			if automa.Posizione != posizioneAttesa {
					t.Errorf("Posizione di '%s' errata: attesa %v, trovata %v", nome, posizioneAttesa, automa.Posizione)
			}
	}
}
