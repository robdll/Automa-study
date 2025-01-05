package entita

import (
	"fmt"
	"strconv"
	"strings"
)

func Esegui(piano *Piano, input string) error {
	args := strings.Fields(input)
	
	cmd := args[0]
	switch cmd {
	case "h":
		PrintHelp()
	case "c":
		piano = NewPiano()
	case "s":
		if len(args) != 3 {
			return fmt.Errorf("@_@^ Il comando 's' deve essere seguito da due coordinate.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("@_@^ Coordinate non valide.")
		}
		piano.Stato(x, y, true)
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
