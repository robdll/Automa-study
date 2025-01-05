package entita

import (
	"fmt"
)

func (p *Piano) EsistePercorso(pointA, pointB [2]int) bool {
	fmt.Printf("Verifica percorso libero per automa da %v a %v\n", pointA, pointB)

	start := pointA
	end := pointB
	if shouldSwap(pointA, pointB) {
		start, end = pointB, pointA
	}

	// Coda per BFS
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true

	// Direzioni possibili
	directions := [][2]int{
		{0, 1}, // Alto
	}

	if start[0] < end[0] {
		directions = append(directions, [2]int{1, 0}) // Destra
	} else {
		directions = append(directions, [2]int{-1, 0}) // Sinistra
	}

	for len(queue) > 0 {
		// Prendi il primo elemento della coda
		current := queue[0]
		// Rimuovi il primo elemento dalla coda
		queue = queue[1:]

		// Termina se raggiungi il punto finale 
		if current == end {
			return true
		}

		// Esplora i vicini
		for _, dir := range directions {
			neighbor := [2]int{current[0] + dir[0], current[1] + dir[1]}

			outsideHorizontalLimit := false
			if  start[0] < end[0] && neighbor[0] > end[0] {
				outsideHorizontalLimit = true
			}
			if  start[0] > end[0] && neighbor[0] < end[0] {
				outsideHorizontalLimit = true
			}
			
			/* Salta nei seguente casi:
			* 1) è già stato visitato,
			* 2) è un ostacolo,
			* 3) il vicino è fuori dai limiti del movimento. */
			if visited[neighbor] || p.isOstacolo(neighbor) || (outsideHorizontalLimit || neighbor[1] > end[1])  {
				continue
			}

			// Marca il vicino come visitato e aggiungilo alla coda
			visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}
	return false
}

// Restituisce vero se B è più in basso o alla stessa alteza ma più a sinistra di A
func shouldSwap(a, b [2]int) bool {
	return b[1] < a[1]
}
