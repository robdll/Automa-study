package entita

import (
	. "progetto-algoritmi/entita"
	"testing"
)

func TestAggiungiOstacolo(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un ostacolo
	p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{1, 1},
		AngoloAltoDestro:    [2]int{3, 3},
	})

	// Verifica che tutte le posizioni del rettangolo siano occupate
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			if _, esiste := p.Mappa[[2]int{x, y}]; !esiste {
				t.Errorf("La posizione (%d, %d) non è occupata dall'ostacolo", x, y)
			}
		}
	}

	// Aggiunge un altro ostacolo sovrapposto
	p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{2, 2},
		AngoloAltoDestro:    [2]int{4, 4},
	})

	// Verifica che le nuove posizioni siano occupate
	for x := 2; x <= 4; x++ {
		for y := 2; y <= 4; y++ {
			if _, esiste := p.Mappa[[2]int{x, y}]; !esiste {
				t.Errorf("La posizione (%d, %d) non è occupata dall'ostacolo", x, y)
			}
		}
	}
}

func TestRimuoviOstacolo(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un ostacolo
	p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{1, 1},
		AngoloAltoDestro:    [2]int{3, 3},
	})

	// Rimuove l'ostacolo
	err := p.RimuoviOstacolo(2, 2)
	if err != nil {
		t.Errorf("Errore nel rimuovere l'ostacolo: %s", err)
	}

	// Verifica che tutte le posizioni siano state liberate
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			if _, esiste := p.Mappa[[2]int{x, y}]; esiste {
				t.Errorf("La posizione (%d, %d) non è stata liberata dopo la rimozione", x, y)
			}
		}
	}

	// Tenta di rimuovere un ostacolo inesistente
	err = p.RimuoviOstacolo(4, 4)
	if err == nil {
		t.Errorf("Errore previsto per ostacolo inesistente non generato")
	}
}

func TestCollisionDetection(t *testing.T) {
	p := NuovoPiano()

	// Aggiunge un ostacolo
	p.AggiungiOstacolo(Rettangolo{
		AngoloBassoSinistro: [2]int{1, 1},
		AngoloAltoDestro:    [2]int{3, 3},
	})

	// Tenta di aggiungere un automa sopra un ostacolo
	err := p.AggiungiAutoma(2, 2, "automa1")
	if err == nil {
		t.Errorf("Errore previsto per aggiunta di un automa sopra un ostacolo non generato")
	}

	// Aggiunge un automa in una posizione libera
	err = p.AggiungiAutoma(0, 0, "automa1")
	if err != nil {
		t.Errorf("Errore nell'aggiungere l'automa: %s", err)
	}

	// Tenta di muovere l'automa sopra un ostacolo
	err = p.MuoviAutoma("automa1", [2]int{1, 1})
	if err == nil {
		t.Errorf("Errore previsto per movimento sopra un ostacolo non generato")
	}
}
