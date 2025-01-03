package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"game/models"
	"game/pkg/auth"
	"game/pkg/utils"
	"game/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}

	token, err := h.Service.CreateUser(&user)
	if err != nil {
		log.Println("Internal Server while Creating User")
		http.Error(w, "Failed to create User", http.StatusInternalServerError)
		return
	}

	auth.SetJWTAsCookie(w, token)
	utils.WriteJson(w, "User created Successfully", http.StatusOK)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}
	//handle validation then send
	if (user.Email == ""  || user.Password == "" ){
		http.Error(w,"Invalid Input",http.StatusBadRequest)
	}
	token, err := h.Service.SignUser(&user)
	if err != nil {
		log.Println("Internal Server while Signing User")
		http.Error(w, "Failed to sign User", http.StatusInternalServerError)
		return
	}

	auth.SetJWTAsCookie(w, token)
	utils.WriteJson(w, "Signed In Successfully", http.StatusOK)
}