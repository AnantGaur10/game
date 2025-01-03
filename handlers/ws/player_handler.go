package handlers

import (
	"encoding/json"
	"fmt"
	"game/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"game/pkg/auth"
	"game/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			return false
		}
		if cookie.Value == "" {
			return false
		}
		_, err = auth.ValidateJWT(cookie.Value)
		
		return err == nil 
	},
}

type PlayerHandler struct {
	RoomService  *services.RoomService
	QueueService *services.QueueService
}

func NewPlayerHandler(rs *services.RoomService,qs *services.QueueService) *PlayerHandler {
	return &PlayerHandler{
		RoomService: rs,
		QueueService: qs,
	}
}

func (p *PlayerHandler) HandlePlayer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		fmt.Println("Failed to create connection",err)
		return
	}
	defer conn.Close()
	var player types.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}
	player.Conn = conn
	player.Game = types.PlayerData{}
	if !p.QueueService.Push(&player){
		var message types.Message[string] = types.Message[string]{
			Type : "Wait",
			Data : "Waiting for another player",
		}
		player.Conn.WriteJSON(message)
	}

	roomID,err := p.RoomService.CreateRoom()
	
	if err != nil {
		log.Println("Not Enough Players in the queue")	
	}
	p.RoomService.HandleRoom(roomID)

}