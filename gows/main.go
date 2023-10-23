package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var trusted_origin string

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var port string

func init() {
	pport := flag.String("port", "81", "Port to listen to")
	ptorigin := flag.String("origin", "http://127.0.0.1:8080", "Trusted origin")

	flag.Parse()

	port = ":" + *pport
	trusted_origin = *ptorigin

	log.Println("port: ", port)
	log.Println("trusted_origin: ", trusted_origin)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == trusted_origin
	}
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(port, nil))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrading error: %#v\n", err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hello client you've connected!"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

// func write(message string) {
// 	log.Println("Client Connected")
// 	err = ws.WriteMessage(1, []byte("Hello client you've connected!"))
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out incoming message
		fmt.Println("incoming message: " + string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
