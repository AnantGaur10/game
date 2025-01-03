package wsserver

import (
	"net/http"

	"github.com/gorilla/mux"

	handlers "game/handlers/ws"
	"game/services"
)

var (
	Mux           *mux.Router             = mux.NewRouter()
	queueService  *services.QueueService  = services.NewQueueService()
	roomService   *services.RoomService   = services.NewRoomService(queueService)
	playerHandler *handlers.PlayerHandler = handlers.NewPlayerHandler(roomService, queueService)
)

func InitRoutes() {
	Mux.HandleFunc("/play", playerHandler.HandlePlayer).Methods(http.MethodGet)
}
