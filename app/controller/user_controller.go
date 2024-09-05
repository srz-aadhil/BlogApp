package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (c *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := c.userService.CreateUser(r)
	if err != nil {
		log.Fatal("User creation failed due to : ", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, userID)
}

func (c *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		log.Fatal("User updation failed due to : ", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User updation successfull")
}

func (c *UserControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := c.userService.GetAllUsers()
	if err != nil {
		log.Fatal("Fetching all users failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, allUsers)

}

func (c *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := c.userService.GetUser(r)
	if err != nil {
		log.Fatal("Fetching single user failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, user)
}

func (c *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		log.Fatal("User deletion failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User deletion successfully completed")
}
