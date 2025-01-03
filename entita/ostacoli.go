package entita

import "fmt"

type Rettangolo struct {
	AngoloBassoSinistro [2]int // Coordinate (x0, y0)
	AngoloAltoDestro    [2]int // Coordinate (x1, y1)
}

func (p *Piano) AggiungiOstacolo(ostacolo Rettangolo) error {
	fmt.Printf("Tentativo di aggiungere ostacolo: (%d, %d) -> (%d, %d)\n",
			ostacolo.AngoloBassoSinistro[0], ostacolo.AngoloBassoSinistro[1],
			ostacolo.AngoloAltoDestro[0], ostacolo.AngoloAltoDestro[1])

	// Controlla se esiste un automa nella posizione del rettangolo
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
			for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
					key := [2]int{x, y}
					if entita, ok := p.Mappa[key]; ok {
							switch entita.(type) {
							case *Automa:
									fmt.Printf("Errore: automa presente alla posizione (%d, %d), impossibile aggiungere l'ostacolo\n", x, y)
									return fmt.Errorf("impossibile aggiungere ostacolo: esiste un automa in (%d, %d)", x, y)
							case *Rettangolo:
									fmt.Printf("Nota: la posizione (%d, %d) è già parte di un altro ostacolo\n", x, y)
							}
					}
			}
	}

	// Aggiunge l'ostacolo alla lista
	p.Ostacoli = append(p.Ostacoli, ostacolo)
	fmt.Printf("Ostacolo aggiunto alla lista: (%d, %d) -> (%d, %d)\n",
			ostacolo.AngoloBassoSinistro[0], ostacolo.AngoloBassoSinistro[1],
			ostacolo.AngoloAltoDestro[0], ostacolo.AngoloAltoDestro[1])

	// Popola tutti i punti del rettangolo nella mappa
	for x := ostacolo.AngoloBassoSinistro[0]; x <= ostacolo.AngoloAltoDestro[0]; x++ {
			for y := ostacolo.AngoloBassoSinistro[1]; y <= ostacolo.AngoloAltoDestro[1]; y++ {
					key := [2]int{x, y}
					p.Mappa[key] = &ostacolo
					fmt.Printf("Aggiunto ostacolo in posizione (%d, %d)\n", x, y)
			}
	}

	fmt.Println("Ostacolo aggiunto correttamente.")
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

