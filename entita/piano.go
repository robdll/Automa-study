package piano

import "fmt"

type Piano struct {
	Automi    map[string]*Automa
	Ostacoli  []Rettangolo
	Mappa     map[[2]int]interface{} // Mappa sparsa per posizioni
}

// Funzione per creare un nuovo piano
func NuovoPiano() *Piano {
	return &Piano{
		Automi:   make(map[string]*Automa),
		Ostacoli: []Rettangolo{},
		Mappa:    make(map[[2]int]interface{}),
	}
}

// Metodo per eseguire un comando
func (p *Piano) EseguiComando(comando string) {
	fmt.Println("Eseguendo comando:", comando)
	// Logica di parsing e gestione dei comandi
}
