package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"

	"game/db"
	handlers "game/handlers/http"
	"game/repositories"
	"game/services"
)

var(
	 Router *mux.Router = mux.NewRouter()
userRepository *repositories.UserRepository = repositories.NewUserRepository(db.Db)
userService *services.UserService= services.NewUserService(userRepository)
userHandler *handlers.UserHandler = handlers.NewUserHandler(userService)

)
func InitRoutes() {
	Router.HandleFunc("/signup", userHandler.SignUp).Methods(http.MethodPost)
	Router.HandleFunc("/signin",userHandler.SignIn).Methods(http.MethodPost)
}