# Progetto Algoritmi: Automi e Segnali

Questo progetto, sviluppato come parte del corso di Algoritmi e Strutture Dati, implementa un sistema per la gestione di automi puntiformi che si muovono su un piano discreto. Gli automi rispondono a segnali e navigano se possono evitare gli ostacoli.

## Caratteristiche Principali

- Modellazione del Piano: Una struttura dati rappresenta il piano bidimensionale, con supporto per automi e ostacoli.

- Gestione degli Automi: Creazione, posizionamento, e movimento degli automi in risposta a segnali.

- Ostacoli: Aggiunta di ostacoli che bloccano il movimento degli automi.

- Percorsi Liberi: Verifica dell’esistenza di percorsi privi di ostacoli per automi.

- Interazione tramite Input Testuale: Interfaccia a linea di comando per interagire con il programma.

## Struttura del Repository

```
progetto-algoritmi
|_ entita
   |_ automi.go      # Gestione degli automi
   |_ ostacoli.go    # Definizione e gestione degli ostacoli
   |_ piano.go       # Struttura dati del piano e operazioni
   |_ utils.go       # Funzioni di utilità
|_ tests
   |_ base-in       # file testuale di input per i test
   |_ base-out      # file testuale di output per i test
   |_ lib_test.go
   |_ utils_test.go
|_ README.md         
|_ solution.go       # Punto di ingresso del programma
|_ go.mod            # Gestione dei moduli Go
|_ progetto-gen2025-v3-5gen2025.pdf  # Specifiche del progetto
```

## Requisiti

- Go: Versione 1.23 o superiore.

- Sistema operativo compatibile con Go.

## Installazione

- Clonare il repository:

`git clone git@github.com:robdll/Automa-study.git`

- Aprire il terminale e spostarsi nella cartella del progetto

`cd progetto-algoritmi`

- Inizializzare i moduli Go:

`go mod init`

- Compilare il programma:

`go build -o solution solution.go`

## Utilizzo

Eseguire il programma compilato:

./solution [-verbose]

## Comandi Disponibili

- `c`: Crea un nuovo piano.

- `s <x> <y>`: Mostra lo stato del piano nel punto specificato.

- `S`: Stampa l’elenco degli automi e degli ostacoli.

- `a <x> <y> <n>`: Posiziona o muove un automa di nome n nel punto (x, y).

- `o <x0> <y0> <x1> <y1>`: Aggiunge un ostacolo tra due coordinate specificate.

- `r <x> <y> <n>`: Emette un segnale da (x, y) con prefisso n.

- `p <n>`: Stampa le posizioni degli automi il cui nome inizia con n.

- `e <x> <y> <n>`: Verifica se un automa n può raggiungere (x, y).

- `G`: Stampa grafica della mappa.

- `f`: Termina il programma.

- `h`: Mostra l’elenco dei comandi disponibili.

## Esempio di Input/Output

### Input

```
c
a 1 2 1
o 0 0 2 2
s 1 2
S
f
```

### Output

```
A
(
automa1: 1,2
)
[
(0,0)(2,2)
]
```

## Testing

Per eseguire i test:

```
cd tests
go test -v
```


Autore

Progetto sviluppato da Roberto Di Lillo. Matricola: [908918].