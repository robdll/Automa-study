package entita

import (
	"fmt"
)


var Silent = true

func ConditionalOutput(message ...interface{}) {
	if !Silent {
		fmt.Println(message...)
	}
}

func GetManhattanDistance(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *Piano) isOstacolo(key [2]int) bool {
	return p.Stato(key[0], key[1], false) == "O"
}

func (p *Piano) isAutoma(key [2]int) bool {
	return p.Stato(key[0], key[1], false) == "A"
}

func (p *Piano) isEmpty(key [2]int) bool {
	return p.Stato(key[0], key[1], false) == "E"
}

func PrintHelp() {
	fmt.Println("Elenco Comandi:")
	fmt.Println("  c                   		- Crea un nuovo piano.")
	fmt.Println("  s <x> <y>           		- Stampa lo stato del piano nel punto (<x>, <y>):")
	fmt.Println("                        	'A' se presente almeno automa, 'O' per gli ostacoli obstacle, 'E' se vuoto.")
	fmt.Println("  S                   		- Stampa tutti gli automi seguiti da tutti gli ostacoli.")
	fmt.Println("  G                   		- Stampa grafica del piano.")
	fmt.Println("  a <x> <y> <n>       		- Posiziona o muove un automa con nome n nel punto (<x>, <y>).")
	fmt.Println("  o <x0> <y0> <x1> <y1>  	- Aggiunge un ostacolo con angoli (<x0>, <y0>) e (<x1>, <y1>).")
	fmt.Println("  r <x> <y> <n>          	- Emette un segnale con nome <n> dalla posizione (<x>, <y>).")
	fmt.Println("  p <n>                 	- Stampa la posizione deglie automi con nome che inizia con prefisso n.")
	fmt.Println("  e <x> <y> <n>          	- Controlla se l'automa n può raggiungere (x, y). Outputs 'SI' or 'NO'.")
	fmt.Println("  f                   		- Termina il programma.")
	fmt.Println("  h                   		- Stampa questo messaggio.")
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
			if p.Stato(key[0], key[1], false) != "E" {
				fmt.Print(p.Stato(key[0], key[1], false), " ")
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
