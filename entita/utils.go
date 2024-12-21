package entita

import "fmt"

// Stampa il piano graficamente con assi numerati
func (p *Piano) StampaGrafica() {
	// Determina i limiti del piano
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for pos := range p.Mappa {
		if pos[0] < minX {
			minX = pos[0]
		}
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] < minY {
			minY = pos[1]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}

	// Stampa il piano con l'asse Y numerato
	for y := maxY; y >= minY; y-- {
		fmt.Printf("%2d  ", y) // Etichetta dell'asse Y con meno spazi
		for x := minX; x <= maxX; x++ {
			key := [2]int{x, y}
			if entita, ok := p.Mappa[key]; ok {
				switch entita.(type) {
				case *Automa:
					fmt.Print("A  ") // Rappresentazione dell'automa
				case *Rettangolo:
					fmt.Print("O  ") // Rappresentazione dell'ostacolo
				}
			} else {
				fmt.Print(".  ") // Spazio vuoto
			}
		}
		fmt.Println()
	}

	// Stampa l'asse X sotto il piano
	fmt.Print("   ") // Spazi per allineare l'asse Y
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%2d", x) // Meno spazio rispetto al contenuto
		if x < maxX {
			fmt.Print(" ") // Spazi tra i valori dell'asse X
		}
	}
	fmt.Println()
}
