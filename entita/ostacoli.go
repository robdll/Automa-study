package entita

import "fmt"

type Ostacolo struct {
	AngoloBassoSinistro [2]int // Coordinate (x0, y0)
	AngoloAltoDestro    [2]int // Coordinate (x1, y1)
}

func (r *Ostacolo) Stampa() {
	fmt.Printf(
		"(%d,%d)(%d,%d)\n",
		r.AngoloBassoSinistro[0],
		r.AngoloBassoSinistro[1],
		r.AngoloAltoDestro[0],
		r.AngoloAltoDestro[1],
	)
}


func (p *Piano) AggiungiOstacolo(x0, y0, x1, y1 int) {
	// Controlla che non ci siano automi nell'area del rettangolo
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			if p.isAutoma([2]int{x, y}) {
				ConditionalOutput("Impossibile posizionare ostacolo in quella posizione.")
				return
			}
		}
	}

	// Crea un nuovo ostacolo
	newOstacolo := Ostacolo{
		AngoloBassoSinistro: [2]int{x0, y0},
		AngoloAltoDestro:    [2]int{x1, y1},
	}

	// Aggiungi l'ostacolo alla lista degli ostacoli
	p.Ostacoli = append(p.Ostacoli, newOstacolo)

	// Aggiorna la mappa del piano
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			key := [2]int{x, y}
			if _, exists := p.Mappa[key]; !exists {
				p.Mappa[key] = []interface{}{}
			}
			p.Mappa[key] = append(p.Mappa[key], &newOstacolo)
		}
	}
	ConditionalOutput(
		"Ostacolo creato: (", x0, ",",y0,") -> (", x1, ",",y1,")",
	)
}