package entita

import (
	"fmt"
	"strconv"
	"strings"
)

func GetManhattanDistance(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func GetKeyFromValues(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func GetValuesFromKey(key string) [2]int {
	split := strings.Split(key, "_")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return [2]int{x, y}
}


// Stampa il piano graficamente sulla console
func (p *Piano) StampaGrafica() {
	var minX, maxX, minY, maxY int
	undefinedValue := true
	for key := range *p.Mappa {
		pos := GetValuesFromKey(key)
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
			key := GetKeyFromValues(x, y)
			if entities, ok := (*p.Mappa)[key]; ok {
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
