package main

import (
	"video-chat-app/server"
	"log"
	"net/http"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/createOrJoin", server.CreateOrJoinRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	go server.Broadcaster()

	log.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
