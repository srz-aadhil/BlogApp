package controller

import (
	"blog/app/dto"
	"blog/app/service/mocks"
	"blog/pkg/e"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)
	tests := []struct {
		name    string
		status  int
		user    *dto.UserCreateRequest
		userId  int64
		want    string
		err     error
		wantErr bool
	}{
		{
			//Success case
			name:   "success test case",
			status: 200,
			user: &dto.UserCreateRequest{
				UserName: "my name",
				Password: "qwerty",
			},
			userId:  2,
			want:    `{"status":"ok","result":2}`,
			err:     nil,
			wantErr: false,
		},

		{
			//error case
			name:   "error case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't create the user","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/create", nil)
			res := httptest.NewRecorder()
			userMock.On("CreateUser", req).Once().Return(test.userId, test.err)
			conn.CreateUser(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}

}

func TestGetUser(t *testing.T) {
	createdAt := time.Date(2024, time.August, 18, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.August, 19, 0, 0, 0, 0, time.UTC)
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)
	tests := []struct {
		name    string
		status  int
		user    *dto.UserResponse
		want    string
		err     error
		wantErr bool
	}{
		{
			//Success case
			name:   "success test case",
			status: 200,
			user: &dto.UserResponse{
				ID:        1,
				UserName:  "monu",
				Password:  "possu",
				Salt:      "1kj1n232w3",
				IsDeleted: false,
				CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				},
				DeleteResponse: dto.DeleteResponse{
					DeletedBy: nil,
					DeletedAt: nil,
				},
			},
			want:    `{"status":"ok","result":{"id":1,"username":"monu","password":"possu","salt":"1kj1n232w3","is_deleted":false,"created_at":"2024-08-18T00:00:00Z","updated_at":"2024-08-19T00:00:00Z"}}`,
			err:     nil,
			wantErr: false,
		},
		{
			//error case
			name:   "error test case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't get a single user","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/1", nil)
			res := httptest.NewRecorder()
			userMock.On("GetUser", req).Once().Return(test.user, test.err)
			conn.GetUser(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())

		})
	}
}

func TestGetAllUsers(t *testing.T) {
	createdAt := time.Date(2024, time.August, 18, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.August, 19, 0, 0, 0, 0, time.UTC)
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)
	tests := []struct {
		name    string
		status  int
		user    *[]dto.UserResponse
		want    string
		err     error
		wantErr bool
	}{
		{
			//success case
			name:   "success case",
			status: 200,
			user: &[]dto.UserResponse{
				{
					ID:        1,
					UserName:  "asad",
					Password:  "1234dddd",
					Salt:      "1234dsdsds",
					IsDeleted: false,
					CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
						CreatedAt: createdAt,
						UpdatedAt: &updatedAt,
					},
					DeleteResponse: dto.DeleteResponse{
						DeletedBy: nil,
						DeletedAt: nil,
					},
				},
				{
					//error case
					ID:        2,
					UserName:  "sadam",
					Password:  "1234",
					Salt:      "123456",
					IsDeleted: false,
					CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
						CreatedAt: createdAt,
						UpdatedAt: &updatedAt,
					},
					DeleteResponse: dto.DeleteResponse{
						DeletedBy: nil,
						DeletedAt: nil,
					},
				},
			},
			want:    `{"status":"ok","result":[{"id":1,"username":"asad","password":"1234dddd","salt":"1234dsdsds","is_deleted":false,"created_at":"2024-08-18T00:00:00Z","updated_at":"2024-08-19T00:00:00Z"},{"id":2,"username":"sadam","password":"1234","salt":"123456","is_deleted":false,"created_at":"2024-08-18T00:00:00Z","updated_at":"2024-08-19T00:00:00Z"}]}`,
			wantErr: false,
			err:     nil,
		},
		{
			//error case
			name:   "error case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal server error",
				RootCause: errors.New("database error"),
			},
			want: `{"status":"not ok","error":{"code":500,"message":"can't get all users","details":["database error"]}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			res := httptest.NewRecorder()
			userMock.On("GetAllUsers").Once().Return(test.user, test.err)
			conn.GetAllUsers(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}
