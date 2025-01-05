package main

import (
	"bufio"
	"fmt"
	"os"
	. "progetto-algoritmi/entita"
	"strconv"
	"strings"
)

func main() {
	p := Crea()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Pronto per la ricezione comandi. 'f' per terminare, 'h' per visualizzare la lista comandi.")
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if line == "f" {
			fmt.Println("Esecuzione Terminata.")
			break
		}

		err := esegui(p, line)
		if err != nil {
			fmt.Println("Comando non valido: ", err)
		}

		fmt.Println("Pronto per la ricezione comandi. 'f' per terminare, 'h' per visualizzare la lista comandi.")
	}

}

func esegui(piano *Piano, input string) error {
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
	case "S":
		piano.Stampa()
	case "a":
		if len(args) != 4 {
			return fmt.Errorf("il comando 's' deve essere seguito da due coordinate e un nome binario.")
		}
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("coordinate non valide.")
		}
		piano.PosizionaAutoma(x, y, args[3])
	case "o":
		if len(args) != 5 {
			return fmt.Errorf("invalid arguments for 'o'")
		}
		x0, _ := strconv.Atoi(args[1])
		y0, _ := strconv.Atoi(args[2])
		x1, _ := strconv.Atoi(args[3])
		y1, _ := strconv.Atoi(args[4])
		piano.AggiungiOstacolo(x0, y0, x1, y1)
	case "r":
		if len(args) != 4 {
			return fmt.Errorf("invalid arguments for 'r'")
		}
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		piano.Richiamo(x, y, args[3])
	case "p":
		// if len(args) != 2 {
		// 	return fmt.Errorf("invalid arguments for 'p'")
		// }
		// piano.Posizioni(args[1]) // Placeholder for posizioni function
	case "e":
		if len(args) != 4 {
			return fmt.Errorf("invalid arguments for 'e'")
		}
		// Add esistePercorso handler
	case "f":
		// Termination is handled in the main loop
		return nil
	default:
		return fmt.Errorf("unsupported command: %s", cmd)
	}

	return nil
}
