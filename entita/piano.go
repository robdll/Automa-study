package entita

import (
	"fmt"
	"strings"
)

// Struttura principale del piano
type Piano struct {
	Automi    map[string]*Automa
	Ostacoli  []Rettangolo
	Mappa     map[[2]int]interface{} // Mappa sparsa per posizioni
}

// Crea un nuovo piano vuoto
func NuovoPiano() *Piano {
	return &Piano{
		Automi:   make(map[string]*Automa),
		Ostacoli: []Rettangolo{},
		Mappa:    make(map[[2]int]interface{}),
	}
}

// Esegue un comando ricevuto
func (p *Piano) EseguiComando(comando string) {
	parts := strings.Fields(comando) // Divide il comando in parole
	if len(parts) == 0 {
		fmt.Println("Comando vuoto, ignorato.")
		return
	}

	switch parts[0] {
	case "c":
		p.Crea()
	case "s":
		if len(parts) < 3 {
			fmt.Println("Comando stato richiede coordinate.")
			return
		}
		fmt.Println(p.Stato(parts[1], parts[2]))
	case "S":
		p.Stampa()
	default:
		fmt.Printf("Comando sconosciuto: %s\n", parts[0])
	}
}

// Operazione "Crea": Resetta il piano
func (p *Piano) Crea() {
	p.Automi = make(map[string]*Automa)
	p.Ostacoli = []Rettangolo{}
	p.Mappa = make(map[[2]int]interface{})
	fmt.Println("Piano resettato.")
}

// Operazione "Stato": Controlla cosa si trova in una posizione
func (p *Piano) Stato(xStr, yStr string) string {
	var x, y int
	_, err := fmt.Sscanf(xStr, "%d", &x)
	if err != nil {
		return "Errore: Coordinata X non valida."
	}
	_, err = fmt.Sscanf(yStr, "%d", &y)
	if err != nil {
		return "Errore: Coordinata Y non valida."
	}

	key := [2]int{x, y}
	if entita, ok := p.Mappa[key]; ok {
		switch entita.(type) {
		case *Automa:
			return "A" // Automa
		case *Rettangolo:
			return "O" // Ostacolo
		}
	}
	return "E" // Vuoto
}

// Operazione "Stampa": Stampa lo stato del piano
func (p *Piano) Stampa() {
	fmt.Println("Automi:")
	for nome, automa := range p.Automi {
		fmt.Printf("%s: (%d, %d)\n", nome, automa.Posizione[0], automa.Posizione[1])
	}

	fmt.Println("Ostacoli:")
	for _, rettangolo := range p.Ostacoli {
		fmt.Printf("(%d, %d) -> (%d, %d)\n",
			rettangolo.AngoloBassoSinistro[0], rettangolo.AngoloBassoSinistro[1],
			rettangolo.AngoloAltoDestro[0], rettangolo.AngoloAltoDestro[1])
	}
}
