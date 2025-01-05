package entita

import (
	"fmt"
	"strconv"
	"strings"
)

type AutomaGroup struct {
	Distanza int
	Automi   []*Automa
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
	fmt.Println("  S                   		- Stampa grafica del piano.")
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
			return fmt.Errorf("il comando 's' deve essere seguito da due coordinate.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("coordinate non valide.")
		}
		piano.Stato(x, y)
	case "G":
		piano.StampaGrafica()
	case "S":
		piano.Stampa()
	case "a":
		if len(args) != 4 {
			return fmt.Errorf("il comando 'a' deve essere seguito da due coordinate e una una stringa in formato binario.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("coordinate non valide.")
		}
		piano.PosizionaAutoma(x, y, args[3])
	case "o":
		if len(args) != 5 {
			return fmt.Errorf("il comando 'o' deve essere seguito da quattro coordinate.")
		}
		x0, _ := strconv.Atoi(args[1])
		y0, _ := strconv.Atoi(args[2])
		x1, _ := strconv.Atoi(args[3])
		y1, _ := strconv.Atoi(args[4])
		piano.AggiungiOstacolo(x0, y0, x1, y1)
	case "r":
		if len(args) != 4 {
			return fmt.Errorf("il comando 'r' deve essere seguito da 2 coordinate e una una stringa in formato binario.")
		}
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		piano.Richiamo(x, y, args[3])
	case "p":
		if len(args) != 2 {
			return fmt.Errorf("il comando 'p' deve essere seguito da una stringa in formato binario.")
		}
		piano.StampaAutomiWithPrefix(args[1])
	case "e":
		if len(args) != 4 {
			return fmt.Errorf("il comando 'e' deve essere seguito da due coordinate e una stringa in formato binario.")
		}
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		key := [2]int{x, y}
		target, err := piano.OttieniAutoma(args[3])
		if err != nil || !piano.EsistePercorso(target.Posizione, key) {
			fmt.Println("NO")
		} else {
			fmt.Println("NO")
		}
	case "f":
		return nil
	default:
		return fmt.Errorf("Comando non valido")
	}

	return nil
}
