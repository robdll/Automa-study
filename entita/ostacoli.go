package entita

type Rettangolo struct {
	AngoloBassoSinistro [2]int // Coordinate (x0, y0)
	AngoloAltoDestro    [2]int // Coordinate (x1, y1)
}

// Aggiunge un ostacolo al piano
func (p *Piano) AggiungiOstacolo(x0, y0, x1, y1 int) {
	// Logica per aggiungere un ostacolo
}
