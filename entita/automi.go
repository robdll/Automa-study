package entita

import "fmt"

type Automa struct {
	Nome     string   // Nome binario dell'automa
	Posizione [2]int  // Coordinate (x, y)
}

// Aggiunge o sposta un automa sul piano
func (p *Piano) AggiungiAutoma(x, y int, nome string) error {
	key := [2]int{x, y}

	// Controlla se la posizione è occupata da un ostacolo
	if _, ok := p.Mappa[key].(*Rettangolo); ok {
		return fmt.Errorf("Impossibile posizionare l'automa: (%d, %d) è occupato da un ostacolo", x, y)
	}

	// Controlla se l'automa esiste già
	if automa, esiste := p.Automi[nome]; esiste {
		// Aggiorna la posizione dell'automa
		delete(p.Mappa, automa.Posizione) // Libera la vecchia posizione
		automa.Posizione = key
		p.Mappa[key] = automa
	} else {
		// Crea un nuovo automa
		automa := &Automa{
			Nome:     nome,
			Posizione: key,
		}
		p.Automi[nome] = automa
		p.Mappa[key] = automa
	}

	return nil
}

// Rimuove un automa dal piano
func (p *Piano) RimuoviAutoma(nome string) error {
	automa, esiste := p.Automi[nome]
	if !esiste {
		return fmt.Errorf("Automa '%s' non trovato", nome)
	}

	// Rimuove l'automa dal piano e dalla mappa
	delete(p.Mappa, automa.Posizione)
	delete(p.Automi, nome)

	return nil
}
