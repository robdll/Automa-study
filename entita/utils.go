package entita

import "fmt"

// Stampa il piano graficamente sulla console
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

	// Stampa il piano riga per riga
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			key := [2]int{x, y}
			if entita, ok := p.Mappa[key]; ok {
				switch entita.(type) {
				case *Automa:
					fmt.Print("A ")
				case *Rettangolo:
					fmt.Print("O ")
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
