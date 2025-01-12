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

func (p *Piano) Stato(x, y string) {
	key := [2]int{GetInt(x), GetInt(y)}
	if entities, exists := (*p.Mappa)[key]; exists {
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
	fmt.Println("(")
	for _, automa := range *p.Automi {
		automa.Stampa()
	}
	fmt.Println(")")
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

func (p *Piano) Richiamo(x, y string, nome string) {

	// Chiave e posizione del richiamo
	key := [2]int{GetInt(x), GetInt(y)}

	// Se è presente un ostacolo o un automa alla posizione di arrivo, allora non serve fare altro
	if entities, exists := (*p.Mappa)[key]; exists && len(entities) > 0 {
		return
	}
	
	automasByDistance := make(map[int][]*Automa)

	// Considera tutti gli automi del piano
	for _, automa := range *p.Automi {

		// Ignora gli che non hanno come prefisso il nome del segnale
		if !strings.HasPrefix(automa.Nome, nome) {
			continue
		}

		// Calcola la distanza Manhattan dell'automa corrente
		distance := GetManhattanDistance(automa.Posizione, key)
		
		// Crea la slice se non esistono ancora automi con la stessa distanza di Manhattan
		if _, exists := automasByDistance[distance]; !exists {
			automasByDistance[distance] = []*Automa{}
		}

		// Aggiungi l'automa alla slice
		automasByDistance[distance] = append(automasByDistance[distance], automa)
		
	}
	
	// Ordina le distanze in ordine crescente
	distances := []int{}
	for distance := range automasByDistance {
		distances = append(distances, distance)
	}
	sort.Ints(distances)

	automaMoved := false
	for i := 0; i < len(distances) && !automaMoved; i++ {
		automasGroup := automasByDistance[distances[i]]
		for _, automa := range automasGroup {
			if p.EsistePercorso(automa.Posizione, key) {
				p.PosizionaAutoma(x, y, automa.Nome)
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

	// Valuta le direzioni possibili (Alto e Sinitra, o Alto e Destra)
	directions := [][2]int{
		{0, 1}, // Alto
	}
	if start[0] < end[0] {
		directions = append(directions, [2]int{1, 0}) // Destra
	} else {
		directions = append(directions, [2]int{-1, 0}) // Sinistra
	}

	// Fino a che ci sono elementi nella coda
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
			horizontalOffset := dir[0]
			verticalOffset := dir[1]
			neighbour := [2]int{current[0] + horizontalOffset, current[1] + verticalOffset}

			outsideVerticalLimit := neighbour[1] > end[1]
			outsideHorizontalLimit := false
			if  start[0] < end[0] && neighbour[0] > end[0] {
				outsideHorizontalLimit = true
			}
			if  start[0] > end[0] && neighbour[0] < end[0] {
				outsideHorizontalLimit = true
			}
			
			/* Salta nei seguente casi:
			* 1) è già stato visitato,
			* 2) è un ostacolo,
			* 3) il vicino è fuori dai limiti del movimento. */
			if visited[neighbour] || p.isOstacolo(neighbour) || (outsideHorizontalLimit || outsideVerticalLimit)  {
				continue
			}

			// Marca il vicino come visitato e aggiungilo alla coda
			visited[neighbour] = true
			queue = append(queue, neighbour)
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