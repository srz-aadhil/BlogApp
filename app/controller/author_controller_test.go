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

func TestGetAllAuthors(t *testing.T)
