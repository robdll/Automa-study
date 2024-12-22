package entita

import "fmt"

type Automa struct {
	Nome     string   // Nome binario dell'automa
	Posizione [2]int  // Coordinate (x, y)
}


// Aggiungi un nuovo automa sul piano (non sposta)
func (p *Piano) AggiungiAutoma(x, y int, nome string) error {
	key := [2]int{x, y}

	// Controlla se la posizione è occupata
	if _, ok := p.Mappa[key]; ok {
		return fmt.Errorf("Impossibile aggiungere l'automa: (%d, %d) è già occupato", x, y)
	}

	// Controlla se un automa con lo stesso nome esiste già
	if _, esiste := p.Automi[nome]; esiste {
		return fmt.Errorf("Un automa con il nome '%s' esiste già", nome)
	}

	// Crea un nuovo automa
	automa := &Automa{
		Nome:     nome,
		Posizione: key,
	}
	p.Automi[nome] = automa
	p.Mappa[key] = automa
	fmt.Printf("Automa '%s' aggiunto a (%d, %d)\n", nome, x, y)

	return nil
}

// Muove un automa esistente a una nuova posizione
func (p *Piano) MuoviAutoma(nome string, nuovaPosizione [2]int) error {
	automa, esiste := p.Automi[nome]
	if !esiste {
		return fmt.Errorf("Automa '%s' non trovato", nome)
	}

	// Controlla se la nuova posizione è occupata da un ostacolo o un altro automa
	if _, occupato := p.Mappa[nuovaPosizione]; occupato {
		return fmt.Errorf("La posizione (%d, %d) è occupata", nuovaPosizione[0], nuovaPosizione[1])
	}

	// Aggiorna la posizione dell'automa
	delete(p.Mappa, automa.Posizione) // Libera la vecchia posizione
	automa.Posizione = nuovaPosizione
	p.Mappa[nuovaPosizione] = automa

	fmt.Printf("Automa '%s' spostato a (%d, %d)\n", nome, nuovaPosizione[0], nuovaPosizione[1])
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

// Elenca tutti gli automi e le loro posizioni
func (p *Piano) ElencaAutomi() {
	fmt.Println("Automi nel piano:")
	for nome, automa := range p.Automi {
		fmt.Printf("Nome: %s, Posizione: (%d, %d)\n", nome, automa.Posizione[0], automa.Posizione[1])
	}
}
