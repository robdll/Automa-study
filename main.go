package main

import (
	"progetto-algoritmi/entita"
)


func main() {
	// Crea un nuovo piano
	p := entita.NuovoPiano()

	// Aggiunge automi
	p.AggiungiAutoma(0, 0, "automa1")
	p.AggiungiAutoma(-1, -1, "automa2")

	// Aggiunge ostacoli
	p.AggiungiOstacolo(entita.Rettangolo{
		AngoloBassoSinistro: [2]int{1, 1},
		AngoloAltoDestro:    [2]int{3, 3},
	})

	// Stampa il piano graficamente
	p.StampaGrafica()
}