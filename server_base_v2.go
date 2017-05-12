package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

/**********************/
/* Structures         */
/**********************/

// Représente une position sur la map
type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Serpent
type Snake struct {
	Kind string `json:"kind"`

	Name  string `json:"name"`
	Color string `json:"color"`

	State string `json:"state"` // "alive" ou "dead"

	// Tableau de positions
	// La tête est le premier élement du tableau
	Body []Pos `json:"body"`

	// WebSocket du client qui le controle
	// `json:"-"` ça veut dire qu'on l'envoie/reçoit pas par le JSON
	WS *websocket.Conn `json:"-"`

	// Pour savoir si le serpent est déjà utilisé
	Used bool `json:"-"`
}

type Update struct {
	Kind string `json:"kind"`

	Snakes []Snake `json:"snakes"`
}

// Structure envoyé dés que le front JS se connecte
type Init struct {
	Kind        string `json:"kind"`
	PlayersSlot []int  `json:"players_slot"`
	StateGame   string `json:"state_game"` // “waiting” or “playing” or “ended”
	MapSize     int    `json:"map_size"`
}

// Va nous permettre d'extraire juste le "kind"
type KindOnly struct {
	Kind string `json:"kind"`
}

/**********************/
/* Variables globales */
/**********************/

// Sert à vérouiller les informations globales
var GeneralMutex sync.Mutex

// Etat du jeu
var StateGame = Init{
	Kind:        "init",
	StateGame:   "waiting",
	MapSize:     50,
	PlayersSlot: []int{1, 2},
}

// 1er joueur, avec une position prédéfinit, et une couleur/nom par défaut
var Player1 = Snake{
	Kind:  "snake",
	Name:  "p1",
	Color: "black",
	State: "alive",
	Body: []Pos{
		Pos{X: 1, Y: 3},
		Pos{X: 1, Y: 2},
		Pos{X: 1, Y: 1},
	},
}

var Move = Move {

}

// Pareil pour le 2ème joueur
var Player2 = Snake{
	Kind: "snake",
	Name:  "p2",
	Color: "black",
	State: "alive",
	Body: []Pos{
		Pos{X: 48, Y: 3},
		Pos{X: 48, Y: 2},
		Pos{X: 48, Y: 1},
	},
}

/**********************/
/* Fonctions          */
/**********************/

/* Main */

func main() {
	http.Handle("/", websocket.Handler(HandleClient))
	fmt.Println("Start on port 43219")
	err := http.ListenAndServe(":43219", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func HandleClient(ws *websocket.Conn) {

	// Dés qu'un client se connecte, on lui envoi l'état de la map
	ws.Write(getInitMessage())
	ws.Write(getUpdateMessage())

	for {
		/*
		** 1- Reception du message
		 */
		var content string
		err := websocket.Message.Receive(ws, &content)
		fmt.Println("Message:", string(content)) // Un peu de debug

		if err != nil {
			fmt.Println(err)
			return
		}

		/*
		** 2- Trouver le type du message
		 */

		var k KindOnly

		err = json.Unmarshal([]byte(content), &k) // JSON Texte -> Obj
		if err != nil {
			fmt.Println(err)
			return
		}

		kind := k.Kind
		fmt.Println("Kind=", kind)

		/*
		** 3- On envoie vers la bonne fonction d'interprétation
		 */

		// On vérouille avant que la fonction fasse une modification
		GeneralMutex.Lock()

		if kind == "move" {
			// Fonction "move"
		} else if kind == "connect" {
			// Fonction "connect"
		} else {
			fmt.Println("Kind inconnue !")
		}

		// On déverouille quand c'est fini
		GeneralMutex.Unlock()
	}
}

// "update" dans le protocole
func getUpdateMessage() []byte {
	var m Update

	m.Kind = "update"
	m.Snakes = []Snake{Player1, Player2}

	message, err := json.Marshal(m) // Transformation de l'objet "Update" en JSON
	if err != nil {
		fmt.Println("Something wrong with JSON Marshal map")
	}
	return message // (Json)
}

// "init" dans le protocole
func getInitMessage() []byte {
	// Transformation de l'objet "Init" en JSON
	message, err := json.Marshal(StateGame)
	if err != nil {
		fmt.Println("Something wrong with JSON Marshal init")
	}
	return message
}

func parseMove(content string) {
	bjson := []byte(content)

	var m Move

	err := json.Unmarshal(bjson, &m)
	if err != nil {
		fmt.Println(nil)
		return
	}
	fmt.Println("Key=", m.Key)
}