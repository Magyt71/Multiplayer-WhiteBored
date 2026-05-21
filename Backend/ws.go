package main

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]string)
var clientsmu sync.Mutex

func wsHandler(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("room")
	if roomId == "" {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Failed to Upgrade the server", "error", err)
		return
	}
	defer conn.Close()

	clientsmu.Lock()
	clients[conn] = roomId
	clientsmu.Unlock()

	defer func() {
		clientsmu.Lock()
		delete(clients, conn)
		clientsmu.Unlock()
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		db.Exec("INSERT INTO strokes (room_id, data) VALUE (?, ?)", roomId, string(msg))

		clientsmu.Lock()
		for clientConn, clientRoom := range clients {
			if clientRoom == roomId && clientConn != conn {
				clientConn.WriteMessage(msgType, msg)
			}
		}
		clientsmu.Unlock()
	}
}
