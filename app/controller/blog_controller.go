package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
)

type BlogController interface {
	CreateBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)
	GetAllBlogs(w http.ResponseWriter, r *http.Request)
	GetOneBlog(w http.ResponseWriter, r *http.Request)
}

var _ BlogController = (*blogControllerImpl)(nil)

type blogControllerImpl struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) BlogController {
	return &blogControllerImpl{
		blogService: blogService,
	}
}

func (c *blogControllerImpl) CreateBlog(w http.ResponseWriter, r *http.Request) {
	blogID, err := c.blogService.CreateBlog(r)
	if err != nil {
		log.Fatal("blog creation failed due  to : ", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, blogID)
}

func (c *blogControllerImpl) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.UpdateBlog(r); err != nil {
		log.Fatal("blog updation failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "blog updation successfully completed")

}

func (c *blogControllerImpl) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.DeleteBlog(r); err != nil {
		log.Fatal("Blog deletion failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Blog deletion successfully completed")
}

func (c *blogControllerImpl) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	result, err := c.blogService.GetBlogs()
	if err != nil {
		log.Fatal("fetching all blogs failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)

}

func (c *blogControllerImpl) GetOneBlog(w http.ResponseWriter, r *http.Request) {
	result, err := c.blogService.GetBlog(r)
	if err != nil {
		log.Fatal("fetching single blog failed due to :", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}
