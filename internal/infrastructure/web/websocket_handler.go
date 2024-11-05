package web

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// HandleConnections gestiona las conexiones websocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error al actualizar la conexión:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error al leer mensaje:", err)
            break
        }
        fmt.Printf("Mensaje recibido: %s\n", msg)

        err = conn.WriteMessage(messageType, []byte("Respuesta desde el servidor"))
        if err != nil {
            fmt.Println("Error al enviar mensaje:", err)
            break
        }
    }
}
