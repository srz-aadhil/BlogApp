package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"blog/pkg/e"
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
		httpErr := e.NewAPIError(err, "can't create the user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, userID)
}

func (c *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		httpErr := e.NewAPIError(err, "can't update the user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User updation successfull")
}

func (c *UserControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := c.userService.GetAllUsers()
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get all users")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, allUsers)

}

func (c *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := c.userService.GetUser(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get a single author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, user)
}

func (c *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		httpErr := e.NewAPIError(err, "can't delete the user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User deletion successfully completed")
}
