package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"errors"
	"log"
	"net/http"
)

var AllRooms RoomMap

func CreateOrJoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.FindOrCreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client    *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func Broadcaster() {
	for {
		msg := <- broadcast

		for _, client := range AllRooms.Map[msg.RoomID] {
			if(client.Conn != msg.Client) {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Println(err)
					client.Conn.Close()
				}
			}
		}
	}
}

func RemoveFromRoom(roomID string, ws *websocket.Conn) error {
	AllRooms.Mutex.Lock()
	defer AllRooms.Mutex.Unlock()

	participants, ok := AllRooms.Map[roomID]
	if !ok {
		return errors.New("room not found")
	}

	for i, participant := range participants {
		if participant.Conn == ws {
			AllRooms.Map[roomID] = append(participants[:i], participants[i+1:]...)
			// Check if the room is empty and delete it if so
			if len(AllRooms.Map[roomID]) == 0 {
				delete(AllRooms.Map, roomID)
			}
			return nil
		}
	}

	return errors.New("participant not found")
}


// JoinRoomRequestHandler will join the client in a particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("roomID missing in URL Parameters")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Web Socket Upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomID[0], false, ws)

	defer func() {
		err := RemoveFromRoom(roomID[0], ws)
		if err != nil {
			log.Println("Error removing participant:", err)
		}
	}()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Read Error: ", err)
		}

		msg.Client = ws
		msg.RoomID = roomID[0]

		log.Println(msg.Message)

		broadcast <- msg
	}
}
