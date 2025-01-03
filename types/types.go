package types

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	ID        string
	State     *GameState         // Room-specific game state
	Players   []*Player // Players in the room
	RoomToken string
	Mu sync.Mutex

}

type GameState struct {
	Round          int
	Players        []*Player
	RoundStartTime time.Time
	RoundEnded     bool
}

type Event struct {
	RoomID   string      // Room where the event originated
	PlayerID string      // Player who triggered the event
	Action   string      // Action type (e.g., "move", "shoot")
	Data     interface{} // Generic data type
}

type Message[T any] struct {
	Type string `json:"type"` // e.g., "update", "move", "shoot"
	Data T      `json:"data"` // Generic data type
}

type PlayerData struct {
	HP      int
	Bullets int
}

type Player struct {
	Name string
	Conn *websocket.Conn
	Game PlayerData
}
