package main

import (
	"fmt"
	"github.com/ntekim/FinChat/chat-service/pkg/websocket"
	"net/http"
)

//define our websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
	fmt.Println(r.Host)

	//go websocket.Writer(ws)
	//websocket.Reader(ws)
}

func setupRoute() {
	pool := websocket.NewPool()
	go pool.Start()
	//map our '/ws' endpoint to the `sweverWs` function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoute()
	http.ListenAndServe(":8088", nil)

}
