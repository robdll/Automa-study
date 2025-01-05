package main

import (
	. "progetto-algoritmi/entita"
)

func main() {

	p := Crea()

	p.AggiungiOstacolo(1, 1, 3, 3)

	p.PosizionaAutoma(0, 0, "0101")

	p.StampaGrafica()

}
