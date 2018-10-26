package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Zett-8/crypchat/log"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	logfile := logger.Init()
	defer logfile.Close()

	// io.MultiWriteで、
	// 標準出力とファイルの両方を束ねて、
	// logの出力先に設定する
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

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
	log.Println("connect")
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
