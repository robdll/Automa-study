package entita

import (
	"fmt"
	"sort"
	"strings"
)

type Piano struct {
	Automi    map[string]*Automa
	Ostacoli  []Ostacolo
	Mappa     map[[2]int][]interface{}
}

func Crea() *Piano {
	fmt.Println("Piano creato.")
	return &Piano{
		Automi:   make(map[string]*Automa),
		Ostacoli: []Ostacolo{},
		Mappa:    make(map[[2]int][]interface{}),
	}
}

func (p *Piano) Stato(x, y int) {
	fmt.Printf("Stato del piano in posizione: (%d, %d)\n", x, y)
	key := [2]int{x, y}
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		switch entities[0].(type) {
		case *Automa:
			fmt.Println("A")
		case *Ostacolo:
			fmt.Println("O")
		}
	} else {
		fmt.Println("E")
	}
}

func (p *Piano) Stampa() {
	if len(p.Automi) > 0 {
		fmt.Println("Automi:")
	}
	for _, automa := range p.Automi {
		automa.Stampa()
	}
	if len(p.Ostacoli) > 0 {
		fmt.Println("Ostacoli:")
	}
	for _, ostacolo := range p.Ostacoli {
		ostacolo.Stampa()
	}
}

func (p *Piano) StampaAutomiWithPrefix(prefix string) {
	for _, automa := range p.Automi {
		if strings.HasPrefix(automa.Nome, prefix) {
			automa.Stampa()
		}
	}
}

func (p *Piano) ListaAutomi() []*Automa {
	automi := make([]*Automa, 0, len(p.Automi))
	for _, automa := range p.Automi {
		automi = append(automi, automa)
	}
	return automi
}

// Stampa il piano graficamente sulla console
func (p *Piano) StampaGrafica() {
	var minX, maxX, minY, maxY int
	undefinedValue := true
	for pos := range p.Mappa {
		if undefinedValue {
			minX, maxX, minY, maxY = pos[0], pos[0], pos[1], pos[1]
			undefinedValue = false
			continue
		}
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

	// Stampa il piano punto per punto dall'alto verso il basso
	for y := maxY; y >= minY; y-- {
		fmt.Printf("%2d | ", y)
		for x := minX; x <= maxX; x++ {
			key := [2]int{x, y}
			if entities, ok := p.Mappa[key]; ok {
				switch entities[0].(type) {
				case *Automa:
					fmt.Print("A  ")
				case *Ostacolo:
					fmt.Print("O  ")
				}
			} else {
				fmt.Print(".  ")
			}
		}
		fmt.Println()
	}

	fmt.Print("    ")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("---")
	}
	fmt.Println()
	fmt.Print("    ")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%2d ", x)
	}
	fmt.Println()
}


func (p *Piano) Richiamo(x, y int, nome string) {
	fmt.Printf("Richiamo %s in posizione (%d, %d)\n", nome, x, y)

	// Chiave della posizione del richiamo
	key := [2]int{x, y}

	// Se è presente un ostacolo o un automa non fare nulla
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		fmt.Printf("Richiamo \"%s\" ignorato, posizione occupata\n", nome)
		return
	}

	automasByDistance := make(map[int]*AutomaGroup)
	distances := []int{}

	// Considera tutti gli automi del piano
	for _, automa := range p.Automi {

		// Se l'automa non ha come prefisso il nome del segnale non far nulla 
		if !strings.HasPrefix(automa.Nome, nome) {
			continue
		}

		// Calcola la distanza Manhattan
		distance := GetManhattanDistance(automa.Posizione, key)
		
		// Check if the distance group exists
		if _, exists := automasByDistance[distance]; !exists {
			automasByDistance[distance] = &AutomaGroup{
				Distanza: distance,
				Automi:   []*Automa{},
			}
			distances = append(distances, distance)
		}
		automasByDistance[distance].Automi = append(automasByDistance[distance].Automi, automa)
		
	}
	
	// Ordina le distanze in ordine crescente O(nlogn)
	sort.Ints(distances)

	automaMoved := false
	// Considera gli automi con distanza in analisi, partendo dalla minima
	for i := 0; i < len(distances) && !automaMoved; i++ {
		automaGroup := automasByDistance[distances[i]]
		for _, automa := range automaGroup.Automi {
			if p.EsistePercorso(automa.Posizione, key) {

				// Rimuovi l'automa dalla mappa precedente
				oldPos := automa.Posizione
    		entities := p.Mappa[oldPos]
        for idx, entity := range entities {
					if entity == automa {
						p.Mappa[oldPos] = append(entities[:idx], entities[idx+1:]...)
						break
					}
        }

        // Se la slice risultante è vuota, rimuovi completamente la chiave
        if len(p.Mappa[oldPos]) == 0 {
					delete(p.Mappa, oldPos)
        }

				// Sposta l'automa alla nuova posizione
				automa.Posizione = key

				// Aggiungi l'automa alla mappa
				if _, exists := p.Mappa[key]; exists {
					p.Mappa[key] = append(p.Mappa[key], automa)
				} else {
					p.Mappa[key] = []interface{}{automa}
				}

				fmt.Printf("Automa %s spostato\n", automa.Nome)

				automaMoved = true
			}
		}
	}
}

func (p *Piano) isOstacolo(key [2]int) bool {
	if entities, exists := p.Mappa[key]; exists && len(entities) > 0 {
		if _, ok := entities[0].(*Ostacolo); ok {
			return true
		}
	}
	return false
}

func (p *Piano) EsistePercorso(pointA, pointB [2]int) bool {

	start := pointA
	end := pointB
	
	// Scambia punti se B è sotto A
	if pointB[1] < pointA[1] {
		start, end = pointB, pointA
	}

	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true

	// Direzioni possibili: Alto, e Destra o Sinistra
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

func (p *Piano) OttieniAutoma(name string) (*Automa, error) {
	automa, exists := p.Automi[name]
	if !exists {
		return nil, fmt.Errorf("Automa '%s' non trovato", name)
	}
	return automa, nil
}