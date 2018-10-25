package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/zett-8/crypchat/log"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	logger.Init()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
