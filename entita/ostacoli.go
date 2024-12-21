package entita

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