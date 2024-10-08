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

func TestCreateBlog(t *testing.T) {
	blogMock := new(mocks.BlogService)
	conn := NewBlogController(blogMock)
	tests := []struct {
		name       string
		status     int
		blogCreate *dto.BlogCreateRequest
		blogID     int64
		err        error
		want       string // dto.BlogResponse
		wantErr    bool
	}{
		{
			//Success case
			name:   "blog test case",
			status: 200,
			blogCreate: &dto.BlogCreateRequest{
				Title:     "blog title",
				Content:   "blog content",
				AuthorID:  1,
				Status:    1,
				CreatedBy: 2,
			},
			blogID:  2,
			err:     nil,
			want:    `{"status":"ok","result":2}`,
			wantErr: false,
		},
		{
			//error case
			name:   "error",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server Error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't create the blog","details":["database error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/create", nil)
			res := httptest.NewRecorder()
			blogMock.On("CreateBlog", req).Once().Return(test.blogID, test.err)
			conn.CreateBlog(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}

func TestUpdateBlog(t *testing.T) {
	blogMock := new(mocks.BlogService)
	conn := NewBlogController(blogMock)
	tests := []struct {
		name       string
		status     int
		blogUpdate *dto.BlogUpdateRequest
		err        error
		want       string //dto.BlogResponse
		wantErr    bool
	}{
		{
			//Success case
			name:   "testing success case",
			status: 200,
			blogUpdate: &dto.BlogUpdateRequest{
				ID:        4,
				Title:     "blog title",
				Content:   "blog content",
				Status:    2,
				UpdatedBy: 1,
			},
			err:     nil,
			want:    `{"status":"ok","result":"blog updation successfully completed"}`,
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
			want:    `{"status":"not ok","error":{"code":500,"message":"can't update the blog","details":["database error"]}}`,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/4,", nil)
			res := httptest.NewRecorder()
			blogMock.On("UpdateBlog", req).Once().Return(test.err)
			conn.UpdateBlog(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}

func TestGetOneBlog(t *testing.T) {
	createdAt := time.Date(2024, time.August, 18, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.August, 19, 0, 0, 0, 0, time.UTC)
	blogMock := new(mocks.BlogService)
	conn := NewBlogController(blogMock)
	tests := []struct {
		name    string
		status  int
		blog    *dto.BlogResponse
		err     error
		want    string //dto.BlogResponse
		wantErr bool
	}{
		{
			//Success case
			name:   "success blog",
			status: 200,
			blog: &dto.BlogResponse{
				ID:       3,
				Title:    "blog title",
				Content:  "blog content",
				AuthorID: 1,
				Status:   2,
				CreatedUpdatedResponse: dto.CreatedUpdatedResponse{
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				},
			},
			err:     nil,
			want:    `{"status":"ok","result":{"id":3,"title":"blog title","content":"blog content","authorid":1,"status":2,"created_at":"2024-08-18T00:00:00Z","updated_at":"2024-08-19T00:00:00Z"}}`,
			wantErr: false,
		},
		{
			//error case
			name:   "error case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal Server error",
				RootCause: errors.New("database error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"can't get a single blog","details":["database error"]}}`,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/3", nil)
			res := httptest.NewRecorder()
			blogMock.On("GetBlog", req).Once().Return(test.blog, test.err)
			conn.GetOneBlog(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}

// func TestGetAllBlogs(t *testing.T){

// }