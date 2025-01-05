package entita

import "fmt"

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
	fmt.Println("  a <x> <y> <n>       		- Posiziona o muove un automa con nome n nel punto (<x>, <y>).")
	fmt.Println("  o <x0> <y0> <x1> <y1>  	- Aggiunge un ostacolo con angoli (<x0>, <y0>) e (<x1>, <y1>).")
	fmt.Println("  r <x> <y> <n>          	- Emette un segnale con nome <n> dalla posizione (<x>, <y>).")
	fmt.Println("  p <n>                 	- Stampa la posizione deglie automi con nome che inizia con prefisso n.")
	fmt.Println("  e <x> <y> <n>          	- Controlla se l'automa n può raggiungere (x, y). Outputs 'SI' or 'NO'.")
	fmt.Println("  f                   		- Termina il programma.")
	fmt.Println("  h                   		- Stampa questo messaggio.")
}