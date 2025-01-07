package entita

import (
	"fmt"
	"sort"
	"strings"
)


type Piano struct {
	Automi    *map[string]*Automa
	Ostacoli  *[]Ostacolo
	Mappa     *map[[2]int][]interface{}
}

func NewPiano() Piano {
	ConditionalOutput("Piano creato.")
	automi := make(map[string]*Automa)
	ostacoli := []Ostacolo{}
	mappa := make(map[[2]int][]interface{})
	return Piano{
		Automi:   &automi,
		Ostacoli: &ostacoli,
		Mappa:    &mappa,
	}
}

func (p *Piano) ResettaPiano() {
	*p.Automi = make(map[string]*Automa)
	*p.Ostacoli = []Ostacolo{}
	*p.Mappa = make(map[[2]int][]interface{})
}

func (p *Piano) Stato(x, y int) {
	ConditionalOutput("Stato del piano in posizione: (",x,",",y, ")")
	key := [2]int{x, y}
	if entities, exists := (*p.Mappa)[key]; exists && len(entities) > 0 {
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
	if len(*p.Automi) > 0 {
		ConditionalOutput("Automi:")
	}
	fmt.Println("(")
	for _, automa := range *p.Automi {
		automa.Stampa()
	}
	fmt.Println(")")
	if len(*p.Ostacoli) > 0 {
		ConditionalOutput("Ostacoli:")
	}
	fmt.Println("[")
	for _, ostacolo := range *p.Ostacoli {
		ostacolo.Stampa()
	}
	fmt.Println("]")
}

func (p *Piano) StampaAutomiWithPrefix(prefix string) {
	fmt.Println("(")
	for _, automa := range *p.Automi {
		if strings.HasPrefix(automa.Nome, prefix) {
			automa.Stampa()
		}
	}
	fmt.Println(")")
}

func (p *Piano) Richiamo(x, y int, nome string) {
	ConditionalOutput("Richiamo", nome, "in posizione (", x,",", y, ")")

	// Chiave della posizione del richiamo
	key := [2]int{x, y}

	// Se è presente un ostacolo o un automa non fare nulla
	if entities, exists := (*p.Mappa)[key]; exists && len(entities) > 0 {
		ConditionalOutput("Richiamo", nome, "ignorato, posizione occupata")
		return
	}

	automasByDistance := make(map[int][]*Automa)
	distances := []int{}

	// Considera tutti gli automi del piano
	for _, automa := range *p.Automi {

		// Se l'automa non ha come prefisso il nome del segnale non far nulla 
		if !strings.HasPrefix(automa.Nome, nome) {
			continue
		}

		// Calcola la distanza Manhattan
		distance := GetManhattanDistance(automa.Posizione, key)
		
		// Check if the distance group exists
		if _, exists := automasByDistance[distance]; !exists {
			automasByDistance[distance] = []*Automa{}
			distances = append(distances, distance)
		}
		automasByDistance[distance] = append(automasByDistance[distance], automa)
		
	}
	
	// Ordina le distanze in ordine crescente O(nlogn)
	sort.Ints(distances)

	automaMoved := false
	// Considera gli automi con distanza in analisi, partendo dalla minima
	for i := 0; i < len(distances) && !automaMoved; i++ {
		automasGroup := automasByDistance[distances[i]]
		for _, automa := range automasGroup {
			if p.EsistePercorso(automa.Posizione, key) {
				p.PosizionaAutoma(x, y, automa.Nome)
				ConditionalOutput("Automa", automa.Nome, "spostato")
				automaMoved = true
			}
		}
	}
}

func (p *Piano) isOstacolo(key [2]int) bool {
	if entities, exists := (*p.Mappa)[key]; exists && len(entities) > 0 {
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
	automa, exists := (*p.Automi)[name]
	if !exists {
		return nil, fmt.Errorf("Automa '%s' non trovato", name)
	}
	return automa, nil
}