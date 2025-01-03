package services

import (
	"fmt"
	"game/types"
	"sync"
	"time"

	"github.com/google/uuid"
)

type RoomService struct {
	rooms map[string]*types.Room
	mu    sync.Mutex
	qs    *QueueService
}

func NewRoomService(qs *QueueService) *RoomService {

	return &RoomService{
		rooms: make(map[string]*types.Room),
		mu:    sync.Mutex{},
		qs:    qs,
	}
}

func (rs *RoomService) generateRoomID() string {

	return uuid.New().String()
}

func (rs *RoomService) CreateRoom() (string,error) {

	Players, err := rs.qs.GetPlayersForRoom()

	if err != nil {
		fmt.Println("Not Enough Players")
		return "",err
	}

	var room types.Room
	var roomPlayers []*types.Player
	room.Mu = sync.Mutex{}
	room.ID = rs.generateRoomID()
	room.RoomToken = rs.generateRoomID()
	roomPlayers = append(roomPlayers, &Players[0], &Players[1])
	room.Players = roomPlayers
	room.State = &types.GameState{
		Round:          0,
		Players:        room.Players,
		RoundStartTime: time.Now(),
		RoundEnded:     false,
	}
	rs.mu.Lock()
	rs.rooms[room.ID] = &room
	rs.mu.Unlock()
	defer func() {
		rs.mu.Lock()
			delete(rs.rooms,room.ID)
		rs.mu.Unlock()
	}()
	for _, player := range Players {

		var message types.Message[string] = types.Message[string]{
			Type: "Start",
			Data: fmt.Sprintf("Match Starting in 3s %s", player.Name),
		}

		player.Conn.WriteJSON(message)
	}
	return room.ID,nil
	// go rs.HandleRoom(&room)
}

func(rs *RoomService) HandleRoom(roomID string) {
	

}