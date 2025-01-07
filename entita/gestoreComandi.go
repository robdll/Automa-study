package entita

import (
	"fmt"
	"os"
	"strings"
)

func Esegui(piano Piano, input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}
	cmd := args[0]
	switch cmd {
	case "h":
		PrintHelp()
	case "c":
		piano.ResettaPiano()
	case "s":
		piano.Stato(args[1], args[2])
	case "G":
		piano.StampaGrafica()
	case "S":
		piano.Stampa()
	case "a":
		piano.PosizionaAutoma(args[1], args[2], args[3])
	case "o":
		piano.AggiungiOstacolo(args[1], args[2], args[3], args[4])
	case "r":
		piano.Richiamo(args[1], args[2], args[3])
	case "p":
		piano.StampaAutomiWithPrefix(args[1])
	case "e":
		key := GetValuesFromKey(args[1] + "-" + args[2])
		target, _ := piano.OttieniAutoma(args[3])
		if piano.EsistePercorso(target.Posizione, key) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}
	case "f":
		os.Exit(0)
	default:
	}
}
