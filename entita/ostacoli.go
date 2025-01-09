package entita

import (
	"fmt"
	"strconv"
)

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

func (p *Piano) AggiungiOstacolo(a, b, c, d string) {
	// Converte le stringhe in interi
	x0, _ := strconv.Atoi(a)
	y0, _ := strconv.Atoi(b)
	x1, _ := strconv.Atoi(c)
	y1, _ := strconv.Atoi(d)

	// Non procedere oltre se esiste un automa dentro i limiti dell'ostacolo
	for _, automa := range *p.Automi {
		if automa.Posizione[0] >= x0 && automa.Posizione[0] <= x1 && automa.Posizione[1] >= y0 && automa.Posizione[1] <= y1 {
			return
		}
	}

	// Crea un nuovo ostacolo
	newOstacolo := Ostacolo{
		AngoloBassoSinistro: [2]int{x0, y0},
		AngoloAltoDestro:    [2]int{x1, y1},
	}

	// Aggiungi l'ostacolo alla lista degli ostacoli
	*p.Ostacoli = append(*p.Ostacoli, newOstacolo)

}