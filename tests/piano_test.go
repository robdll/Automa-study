package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

// Test per la creazione di un nuovo piano
func TestNuovoPiano(t *testing.T) {
	p := NuovoPiano()
	if len(p.Automi) != 0 || len(p.Ostacoli) != 0 || len(p.Mappa) != 0 {
		t.Errorf("NuovoPiano non crea un piano vuoto.")
	}
}

// Test per l'operazione "Crea"
func TestCrea(t *testing.T) {
	p := NuovoPiano()

	// Aggiungiamo dati al piano per testare il reset
	p.Automi["test"] = &Automa{Nome: "test", Posizione: [2]int{1, 1}}
	p.Mappa[[2]int{1, 1}] = p.Automi["test"]
	p.Ostacoli = append(p.Ostacoli, Rettangolo{
		AngoloBassoSinistro: [2]int{0, 0},
		AngoloAltoDestro:    [2]int{2, 2},
	})

	// Resetta il piano
	p.Crea()
	if len(p.Automi) != 0 || len(p.Ostacoli) != 0 || len(p.Mappa) != 0 {
		t.Errorf("Crea non resetta il piano correttamente.")
	}
}

// Test per l'operazione "Stato"
func TestStato(t *testing.T) {
	p := NuovoPiano()

	// Aggiungiamo un automa e un ostacolo
	p.Automi["test"] = &Automa{Nome: "test", Posizione: [2]int{1, 1}}
	p.Mappa[[2]int{1, 1}] = p.Automi["test"]

	p.Ostacoli = append(p.Ostacoli, Rettangolo{
		AngoloBassoSinistro: [2]int{0, 0},
		AngoloAltoDestro:    [2]int{2, 2},
	})
	p.Mappa[[2]int{0, 0}] = &p.Ostacoli[0]

	// Test stato per automa
	if stato := p.Stato("1", "1"); stato != "A" {
		t.Errorf("Stato per (1,1) dovrebbe essere 'A', ma è %s", stato)
	}

	// Test stato per ostacolo
	if stato := p.Stato("0", "0"); stato != "O" {
		t.Errorf("Stato per (0,0) dovrebbe essere 'O', ma è %s", stato)
	}

	// Test stato per posizione vuota
	if stato := p.Stato("2", "2"); stato != "E" {
		t.Errorf("Stato per (2,2) dovrebbe essere 'E', ma è %s", stato)
	}
}

// Test per l'operazione "Stampa"
func TestStampa(t *testing.T) {
	p := NuovoPiano()

	// Aggiungiamo un automa
	p.Automi["test"] = &Automa{Nome: "test", Posizione: [2]int{1, 1}}
	p.Mappa[[2]int{1, 1}] = p.Automi["test"]

	// Aggiungiamo un ostacolo
	p.Ostacoli = append(p.Ostacoli, Rettangolo{
		AngoloBassoSinistro: [2]int{0, 0},
		AngoloAltoDestro:    [2]int{2, 2},
	})
	p.Mappa[[2]int{0, 0}] = &p.Ostacoli[0]

	// Verifica manuale con stampa
	p.Stampa() // Dovrebbe stampare automi e ostacoli
}
