package entita

import "fmt"

type Rettangolo struct {
	AngoloBassoSinistro [2]int // Coordinate (x0, y0)
	AngoloAltoDestro    [2]int // Coordinate (x1, y1)
}

// Aggiunge un ostacolo al piano e lo popola nella mappa
func (p *Piano) AggiungiOstacolo(ostacolo Rettangolo) error {
	// Controlla se esiste un automa dove si vuole posizionare il rettangolo
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
		for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
			if _, ok := p.Mappa[[2]int{x, y}].(*Automa); ok {
				return fmt.Errorf("Impossibile aggiungere ostacolo: esiste un automa in (%d, %d)", x, y)
			}
		}
	}

	// Aggiunge l'ostacolo alla lista
	p.Ostacoli = append(p.Ostacoli, ostacolo)

	// Popola tutti i punti del rettangolo nella mappa
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
		for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
			p.Mappa[[2]int{x, y}] = &ostacolo
		}
	}
	return nil
}


func (p *Piano) RimuoviOstacolo(ostacolo Rettangolo) error {
	// Find the obstacle in the list
	index := -1
	for i, o := range p.Ostacoli {
		if o == ostacolo {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Ostacolo non trovato")
	}

	// Rimuove l'ostacolo dalla lista
	p.Ostacoli = append(p.Ostacoli[:index], p.Ostacoli[index+1:]...)

	// Clear the obstacle's area from the map
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
		for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
			if _, ok := p.Mappa[[2]int{x, y}].(*Rettangolo); ok {
				delete(p.Mappa, [2]int{x, y})
			}
		}
	}
	return nil
}

