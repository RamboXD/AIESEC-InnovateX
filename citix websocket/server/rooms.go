package server

import (
	"github.com/gorilla/websocket"
	"math/rand"
	"sync"
	"time"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

func (r *RoomMap) FindOrCreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	availableRooms := make([]string, 0)

	for roomID, participants := range r.Map {
		if len(participants) < 2 {
			availableRooms = append(availableRooms, roomID)
		}
	}

	if len(availableRooms) > 0 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(availableRooms))
		return availableRooms[randomIndex]
	}

	// Create a new room if no available room is found
	return r.CreateRoom()
}

func (r *RoomMap) CreateRoom() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID
}

func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}
	r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
