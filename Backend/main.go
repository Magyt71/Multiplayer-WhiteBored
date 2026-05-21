package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	initdb()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/register", registerHandler)
	mux.HandleFunc("POST /api/login", loginHandler)

	mux.HandleFunc("POST /api/rooms", func(w http.ResponseWriter, r *http.Request) {
		roomID := uuid.New().String()
		db.Exec("INSERT INTO rooms (id,name) VALUES (?,?)", roomID, "New Room")
		sendJSON(w, http.StatusCreated, map[string]string{"id": roomID})
	})

	mux.HandleFunc("GET /api/rooms/{id}/history", func(w http.ResponseWriter, r *http.Request) {
		roomID := r.PathValue("id")
		rows, _ := db.Query("SELECT data FROM strokes Where room_id=?", roomID)
		defer rows.Close()

		var history []json.RawMessage
		for rows.Next() {
			var data string
			rows.Scan(&data)
			history = append(history, json.RawMessage(data))
		}
		sendJSON(w, http.StatusOK, history)
	})

	mux.HandleFunc("/ws", wsHandler)

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST , GET ,OPTIONs")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	slog.Info("Server Working on port 8080")
	http.ListenAndServe(":8080", corsHandler(mux))

}
