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

func TestGetOneAuthor(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		want    string // dto.AuthorResponse
		author  *dto.AuthorResponse
		error   error
		wantErr bool
	}{
		// Success case
		{
			name:   "success case",
			status: 200,
			author: &dto.AuthorResponse{
				ID:   1,
				Name: "testing name",
				CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				},
				DeleteResponse: dto.DeleteResponse{
					DeletedAt: &createdAt,
					DeletedBy: nil,
				},
			},
			want:    `{"status":"ok","result":{"id":1,"name":"testing name","created_at":"2024-07-15T00:00:00Z","updated_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}}`,
			error:   nil,
			wantErr: false,
		},

		//Error case
		{
			name:   "error case",
			status: 400,
			error: &e.WrapError{
				ErrorCode: 400,
				Msg:       "Bad request",
				RootCause: errors.New("invalid request"),
			},
			want:    `{"status":"not ok","error":{"code":400,"message":"can't get single author","details":["invalid request"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/1", nil)
			res := httptest.NewRecorder()
			authorMock.On("GetAuthor", req).Once().Return(test.author, test.error)
			conn.GetOneAuthor(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}

}

func TestGetAllAuthors(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		author  *[]dto.AuthorResponse
		want    string //dto.AuthorResponse
		error   error
		wantErr bool
	}{
		//Success case
		{
			name:   "success case",
			status: 200,
			author: &[]dto.AuthorResponse{
				{
					ID:   2,
					Name: "testing 1",
					CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
						CreatedAt: createdAt,
						UpdatedAt: &updatedAt,
					},
					DeleteResponse: dto.DeleteResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
				{
					ID:   3,
					Name: "testing 2",
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
			want:    `{"status":"ok","result":[{"id":2,"name":"testing 1","created_at":"2024-07-15T00:00:00Z","updated_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"},{"id":3,"name":"testing 2","created_at":"2024-07-15T00:00:00Z","updated_at":"2024-07-15T00:00:00Z"}]}`,
			wantErr: false,
		},

		{
			name:   "error case",
			status: 500,
			author: nil,
			error: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't get all authors","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			res := httptest.NewRecorder()
			authorMock.On("GetAuthors").Once().Return(test.author, test.error)
			conn.GetaAllAuthors(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())

		})
	}

}

func TestUpdateAuthor(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedBy := 2
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		author  *dto.AuthorResponse
		want    string //dto.AuthoResponse
		error   error
		wantErr bool
	}{
		//Success case
		{
			name:   "success case",
			status: 200,
			author: &dto.AuthorResponse{
				ID:   3,
				Name: "updated author",
				CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
					CreatedAt: createdAt,
					UpdatedBy: &updatedBy,
					UpdatedAt: &updatedAt,
				},
				DeleteResponse: dto.DeleteResponse{
					DeletedBy: nil,
					DeletedAt: nil,
				},
			},
			want:    `{"status":"ok","result":"Author Updation Success"}`,
			wantErr: false,
		},
		{
			//error case
			name:   "error case",
			status: 500,
			error: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't update author","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/2", nil)
			res := httptest.NewRecorder()
			authorMock.On("UpdateAuthor", req).Once().Return(test.error)
			conn.UpdateAuthor(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}

func TestDeleteAuthor(t *testing.T) {
	createdAt := time.Date(2024, time.July, 16, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 16, 0, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2024, time.July, 17, 0, 0, 0, 0, time.UTC)
	deletedBy := 3
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		author  *dto.AuthorResponse
		want    string //dto.AuthorReaponse
		error   error
		wantErr bool
	}{
		{
			//Success case
			name:   "success case",
			status: 200,
			author: &dto.AuthorResponse{
				ID:   4,
				Name: "testing case1",
				CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				},
				DeleteResponse: dto.DeleteResponse{
					DeletedBy: &deletedBy,
					DeletedAt: &deletedAt,
				},
			},
			want:    `{"status":"ok","result":"Author deletion successfull"}`,
			wantErr: false,
		},

		{
			//error case
			name:   "testing case 2",
			status: 500,
			error: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't delete author","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/4", nil)
			res := httptest.NewRecorder()
			authorMock.On("DeleteAuthor", req).Once().Return(test.error)
			conn.DeleteAuthor(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}
