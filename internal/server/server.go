package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server interface {
	Run() error
}

func New(addr string) Server {
	return &webSocketServer{
		addr:     fmt.Sprintf(":%s", addr),
		upgrader: websocket.Upgrader{},
	}
}

type webSocketServer struct {
	addr     string
	upgrader websocket.Upgrader
}

// Run implements Server.
func (s *webSocketServer) Run() error {
	http.HandleFunc("/ws", s.webSocketHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	return http.ListenAndServe(s.addr, nil)
}

func (s *webSocketServer) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
