package entita

// func (p *Piano) EsistePercorso(start, end [2]int) bool {
// 	fmt.Printf("Verifica percorso libero per automa da %v a %v\n", start, end)

// 	// Determine the direction of movement
// 	xStep := 1
// 	if start[0] > end[0] {
// 			xStep = -1
// 	}
// 	yStep := 1
// 	if start[1] > end[1] {
// 			yStep = -1
// 	}

// 	// Traverse horizontally
// 	for x := start[0]; x != end[0]; x += xStep {
// 			if _, ok := p.Mappa[[2]int{x, start[1]}].(*Ostacolo); ok {
// 					fmt.Printf("Ostacolo trovato in posizione (%d, %d), percorso bloccato.\n", x, start[1])
// 					return false
// 			}
// 	}

// 	// Traverse vertically
// 	for y := start[1]; y != end[1]; y += yStep {
// 			if _, ok := p.Mappa[[2]int{end[0], y}].(*Ostacolo); ok {
// 					fmt.Printf("Ostacolo trovato in posizione (%d, %d), percorso bloccato.\n", end[0], y)
// 					return false
// 			}
// 	}

// 	fmt.Printf("Percorso trovato per automa da %v a %v\n", start, end)
// 	return true
// }

// GetDistanzaManhattan calcola la distanza Manhattan tra due punti
func GetDistanzaManhattan(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

// abs calcola il valore assoluto di un numero intero
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
