package entita

// EsistePercorso verifica se esiste un percorso libero tra due punti
func (p *Piano) EsistePercorso(start, end [2]int) bool {
	// Se la posizione iniziale o finale è occupata da un ostacolo, ritorna false
	if _, ok := p.Mappa[start].(*Rettangolo); ok {
		return false
	}
	if _, ok := p.Mappa[end].(*Rettangolo); ok {
		return false
	}

	// Direzioni per i movimenti (sinistra, destra, su, giù)
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	// Mappa per tenere traccia delle posizioni visitate
	visited := make(map[[2]int]bool)
	visited[start] = true

	// Coda per BFS
	queue := [][2]int{start}

	// BFS loop
	for len(queue) > 0 {
		// Estrae il primo elemento dalla coda
		current := queue[0]
		queue = queue[1:]

		// Controlla se siamo arrivati alla fine
		if current == end {
			return true
		}

		// Esplora le posizioni vicine
		for _, dir := range directions {
			next := [2]int{current[0] + dir[0], current[1] + dir[1]}

			// Salta le posizioni già visitate o occupate da ostacoli
			if visited[next] {
				continue
			}
			if _, ok := p.Mappa[next].(*Rettangolo); ok {
				continue
			}

			// Marca la posizione come visitata e aggiungila alla coda
			visited[next] = true
			queue = append(queue, next)
		}
	}

	// Nessun percorso trovato
	return false
}
