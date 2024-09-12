package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"blog/pkg/e"
	"net/http"
)

type UserService interface {
	CreateUser(r *http.Request) (int64, error)
	UpdateUser(r *http.Request) error
	GetUser(r *http.Request) (*dto.UserResponse, error)
	GetAllUsers() (*[]dto.UserResponse, error)
	DeleteUser(r *http.Request) error
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(r *http.Request) (int64, error) {
	req := &dto.UserCreateRequest{}
	if err := req.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "user create parse error", err)
	}

	if err := req.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "user create validation error", err)
	}

	userID, err := s.userRepo.Create(req)
	if err != nil {
		return 0, e.NewError(e.ErrInvalidRequest, "user creation error", err)
	}

	return userID, nil

}

func (s *UserServiceImpl) UpdateUser(r *http.Request) error {
	req := &dto.UserUpdateRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "user update request decode error", err)
	}

	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "user update validation failed", err)
	}

	if err := s.userRepo.Update(req); err != nil {
		return e.NewError(e.ErrInternalServer, "user updation error", err)
	}

	return nil
}

func (s *UserServiceImpl) GetUser(r *http.Request) (*dto.UserResponse, error) {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "user request parse error", err)
	}

	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "user request validation failed", err)
	}

	user, err := s.userRepo.GetOne(req.ID)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "not found user with id", err)
	}

	return user, nil
}

func (s *UserServiceImpl) GetAllUsers() (*[]dto.UserResponse, error) {
	result, err := s.userRepo.GetAll()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "All users request error", err)
	}

	var usersList []dto.UserResponse

	for _, val := range *result {

		var user dto.UserResponse
		user.ID = val.ID
		user.UserName = val.UserName
		user.Password = val.Password
		user.Salt = val.Salt
		user.CreatedBy = val.CreatedBy
		user.CreatedAt = val.CreatedAt
		user.UpdatedBy = val.UpdatedBy
		user.UpdatedAt = val.UpdatedAt
		user.DeletedBy = val.DeletedBy
		user.DeletedAt = val.DeletedAt

		usersList = append(usersList, user)

	}

	return &usersList, nil
}

func (s *UserServiceImpl) DeleteUser(r *http.Request) error {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "user delete parse error", err)
	}

	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrInvalidRequest, "user deletion validate error", err)
	}

	if err := s.userRepo.Delete(req.ID); err != nil {
		return e.NewError(e.ErrResourceNotFound, "user not found with id", err)
	}

	return nil
}
