package entita

import "fmt"

type Rettangolo struct {
	AngoloBassoSinistro [2]int // Coordinate (x0, y0)
	AngoloAltoDestro    [2]int // Coordinate (x1, y1)
}

// Aggiunge un ostacolo al piano e lo popola nella mappa
func (p *Piano) AggiungiOstacolo(ostacolo Rettangolo) {
	// Aggiunge l'ostacolo alla lista
	p.Ostacoli = append(p.Ostacoli, ostacolo)

	// Popola tutti i punti del rettangolo nella mappa
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
		for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
			p.Mappa[[2]int{x, y}] = &ostacolo
		}
	}
}

// Rimuove un ostacolo dal piano
func (p *Piano) RimuoviOstacolo(x, y int) error {
	key := [2]int{x, y}

	// Trova l'ostacolo nella posizione specificata
	entita, esiste := p.Mappa[key]
	if !esiste {
		return fmt.Errorf("Nessun ostacolo trovato a (%d, %d)", x, y)
	}

	ostacolo, ok := entita.(*Rettangolo)
	if !ok {
		return fmt.Errorf("La posizione (%d, %d) non contiene un ostacolo", x, y)
	}

	// Rimuove tutti i punti occupati dall'ostacolo
	for i := ostacolo.AngoloBassoSinistro[0]; i <= ostacolo.AngoloAltoDestro[0]; i++ {
		for j := ostacolo.AngoloBassoSinistro[1]; j <= ostacolo.AngoloAltoDestro[1]; j++ {
			delete(p.Mappa, [2]int{i, j})
		}
	}

	// Rimuove l'ostacolo dalla lista
	for index, rect := range p.Ostacoli {
		if rect == *ostacolo {
			p.Ostacoli = append(p.Ostacoli[:index], p.Ostacoli[index+1:]...)
			break
		}
	}

	fmt.Printf("Ostacolo a (%d, %d) rimosso.\n", x, y)
	return nil
}
