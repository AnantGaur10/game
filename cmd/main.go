package main

import (
	httpserver "game/cmd/http"
	wsserver "game/cmd/ws"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	mainRouter := mux.NewRouter()
	
	httpserver.InitRoutes()
	wsserver.InitRoutes()

	mainRouter.Handle("/signup", httpserver.Router).Methods(http.MethodPost) 
	mainRouter.Handle("/signin",httpserver.Router).Methods(http.MethodPost)
	mainRouter.Handle("/play", wsserver.Mux).Methods(http.MethodGet) 

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mainRouter))

}