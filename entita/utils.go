package entita

import (
	"fmt"
	"strconv"
	"strings"
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


func Esegui(piano *Piano, input string) error {
	args := strings.Fields(input)
	
	cmd := args[0]
	switch cmd {
	case "h":
		PrintHelp()
	case "c":
		*piano = *Crea()
	case "s":
		if len(args) != 3 {
			return fmt.Errorf("@_@^ Il comando 's' deve essere seguito da due coordinate.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("@_@^ Coordinate non valide.")
		}
		piano.Stato(x, y)
	case "G":
		piano.StampaGrafica()
	case "S":
		piano.Stampa()
	case "a":
		if len(args) != 4 {
			return fmt.Errorf("@_@^ Il comando 'a' deve essere seguito da due coordinate e una una stringa in formato binario.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("@_@^ coordinate non valide.")
		}
		piano.PosizionaAutoma(x, y, args[3])
	case "o":
		if len(args) != 5 {
			return fmt.Errorf("@_@^ Il comando 'o' deve essere seguito da quattro coordinate.")
		}
		x0, _ := strconv.Atoi(args[1])
		y0, _ := strconv.Atoi(args[2])
		x1, _ := strconv.Atoi(args[3])
		y1, _ := strconv.Atoi(args[4])
		piano.AggiungiOstacolo(x0, y0, x1, y1)
	case "r":
		if len(args) != 4 {
			return fmt.Errorf("@_@^ Il comando 'r' deve essere seguito da 2 coordinate e una una stringa in formato binario.")
		}
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		piano.Richiamo(x, y, args[3])
	case "p":
		if len(args) != 2 {
			return fmt.Errorf("@_@^ Il comando 'p' deve essere seguito da una stringa in formato binario.")
		}
		piano.StampaAutomiWithPrefix(args[1])
	case "e":
		if len(args) != 4 {
			return fmt.Errorf("@_@^ Il comando 'e' deve essere seguito da due coordinate e una stringa in formato binario.")
		}
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		key := [2]int{x, y}
		target, err := piano.OttieniAutoma(args[3])
		if err != nil {
			return fmt.Errorf("@_@^ Automa non trovato.")
		}
		if piano.EsistePercorso(target.Posizione, key) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}
	case "f":
		return nil
	default:
		return fmt.Errorf("@_@^ Comando non riconosciuto.")
	}
	return nil
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
