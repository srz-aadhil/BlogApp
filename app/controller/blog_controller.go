package controller

import (
	"blog/app/dto"
	"encoding/json"
	"log"
	"net/http"
)

type BlogController interface {
	GetAllBlogs(w http.ResponseWriter, r *http.Request)
	GetOneBlog(w http.ResponseWriter, r *http.Request)
}

var _ BlogController = (*blogControllerImpl)(nil)

type blogControllerImpl struct{}

func NewBlogController() BlogController {
	return &blogControllerImpl{}
}

func (c *blogControllerImpl) GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	var blogs []dto.BlogResponse

	blog1 := dto.BlogResponse{
		ID:      2,
		Title:   "Second blog",
		Content: "Second content",
	}
	blog2 := dto.BlogResponse{
		ID:      3,
		Title:   "Third blog",
		Content: "Third content",
	}

	blogs = append(blogs, blog1, blog2)

	jsonData, err := json.Marshal(blogs)
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

func (c *blogControllerImpl) GetOneBlog(w http.ResponseWriter, r *http.Request) {

	blog := dto.BlogResponse{
		ID:      1,
		Title:   "my blog",
		Content: "blog content",
	}
	jsonData, err := json.Marshal(blog)
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
