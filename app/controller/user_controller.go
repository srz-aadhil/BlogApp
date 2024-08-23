package controller

import (
	"blog/app/dto"
	"encoding/json"
	"log"
	"net/http"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetOneUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct{}

func NewUserController() UserController {
	return &UserControllerImpl{}
}

func (c *UserControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []dto.UserResponse

	user1 := dto.UserResponse{
		ID:       1,
		Username: "Adam",
	}

	user2 := dto.UserResponse{
		ID:       2,
		Username: "Ahmad",
	}

	users = append(users, user1, user2)

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Printf("error due to : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
func (c *UserControllerImpl) GetOneUser(w http.ResponseWriter, r *http.Request) {

	user3 := dto.UserResponse{
		ID:       3,
		Username: "Anandhu",
	}

	jsonData, err := json.Marshal(user3)
	if err != nil {
		log.Printf("error due to :%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}

	w.Header().Set("Content_Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
