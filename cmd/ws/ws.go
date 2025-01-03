package wsserver

import (
	"net/http"

	"github.com/gorilla/mux"

	"game/db"
	handlers "game/handlers/ws"
	"game/pkg/auth"
	"game/repositories"
	"game/services"
)

var (
	Mux           *mux.Router             = mux.NewRouter()
	userRepository *repositories.UserRepository = repositories.NewUserRepository(db.Db)
	userService   *services.UserService   = services.NewUserService(userRepository)
	queueService  *services.QueueService  = services.NewQueueService()
	roomService   *services.RoomService   = services.NewRoomService(queueService)
	playerHandler *handlers.PlayerHandler = handlers.NewPlayerHandler(roomService, queueService,userService)
)

func InitRoutes() {
	Mux.Handle("/play", auth.JwtMiddleWare(http.HandlerFunc(playerHandler.HandlePlayer))).Methods(http.MethodGet)
}
