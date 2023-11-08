package main

import (
	"Realtime_Chat/websocket"
	"fmt"
	"net/http"
)

//"Actualiza" una conexión HTTP estándar a una conexión WebSocket
func serveWs(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
    go websocket.Writer(ws)
    websocket.Reader(ws)
}

//Configura las rutas para el servidor HTTP.
func setupRoutes() {
    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":3000", nil)
}