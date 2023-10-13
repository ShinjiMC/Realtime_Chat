package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Declaramos upgrader definiendo el tamaño de Read and Write con un validador de origen que en este caso será general
var upgrader = websocket.Upgrader {
    ReadBufferSize:  1024,
	WriteBufferSize: 1024,
  	CheckOrigin: func(r *http.Request) bool { return true },
}

//Definimos reader que recibira los mensajes que se envian al websocket
func reader(conn *websocket.Conn) {
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(string(p))
        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

//"Actualiza" una conexión HTTP estándar a una conexión WebSocket
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
	}
    reader(ws)
}

//Configura las rutas para el servidor HTTP.
func setupRoutes() {
  	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola Mundo, es Server!!")
	})
    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Hola Mundo, Es Terminal?")
    setupRoutes()
    http.ListenAndServe(":3000", nil)
}