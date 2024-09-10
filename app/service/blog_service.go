package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
)

type BlogService interface {
	CreateBlog(r *http.Request) (int64, error)
	UpdateBlog(r *http.Request) error
	DeleteBlog(r *http.Request) error
	GetBlog(r *http.Request) (*dto.BlogResponse, error)
	GetBlogs() (*[]dto.BlogResponse, error)
}

var _ BlogService = (*BlogServiceImpl)(nil)

type BlogServiceImpl struct {
	blogRepo repo.BlogRepo
}

func NewBlogService(blogRepo repo.BlogRepo) BlogService {
	return &BlogServiceImpl{
		blogRepo: blogRepo,
	}
}

func (s *BlogServiceImpl) CreateBlog(r *http.Request) (int64, error) {
	body := &dto.BlogCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, err
	}

	if err := body.Validate(); err != nil {
		return 0, err
	}

	blogID, err := s.blogRepo.Create(body)
	if err != nil {
		return 0, err
	}
	return blogID, nil

}

func (s *BlogServiceImpl) UpdateBlog(r *http.Request) error {
	body := &dto.BlogUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}

	if err := body.Validate(); err != nil {
		return err
	}

	if err := s.blogRepo.Update(body); err != nil {
		return err
	}
	return nil
}

func (s *BlogServiceImpl) DeleteBlog(r *http.Request) error {
	req := &dto.BlogDeleteRequest{}
	if err := req.Parse(r); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return err
	}

	if err := s.blogRepo.Delete(req); err != nil {
		return err
	}
	return nil
}

func (s *BlogServiceImpl) GetBlog(r *http.Request) (*dto.BlogResponse, error) {
	body := &dto.BlogRequest{}
	if err := body.Parse(r); err != nil {
		return nil, err
	}

	if err := body.Validate(); err != nil {
		return nil, err
	}

	blog, err := s.blogRepo.Getblog(body)
	if err != nil {
		return nil, err
	}

	var blogResp dto.BlogResponse

	blogResp.ID = blog.ID
	blogResp.Title = blog.Title
	blogResp.Content = blog.Content
	blogResp.AuthorID = blog.AuthorID
	blogResp.Status = blog.Status
	blogResp.CreatedBy = blog.CreatedBy
	blogResp.CreatedAt = blog.CreatedAt
	blogResp.UpdatedBy = blog.UpdatedBy
	blogResp.UpdatedAt = blog.UpdatedAt
	blogResp.DeletedAt = blog.DeletedAt
	blogResp.DeletedBy = blog.DeletedBy

	return &blogResp, nil
}

func (s *BlogServiceImpl) GetBlogs() (*[]dto.BlogResponse, error) {
	results, err := s.blogRepo.GetBlogs()
	if err != nil {
		return nil, err
	}

	var blogList []dto.BlogResponse
	for _, val := range *results {

		var blogResp dto.BlogResponse
		blogResp.ID = val.ID
		blogResp.Title = val.Title
		blogResp.Content = val.Content
		blogResp.AuthorID = val.AuthorID
		blogResp.CreatedBy = val.CreatedBy
		blogResp.CreatedAt = val.CreatedAt
		blogResp.UpdatedBy = val.UpdatedBy
		blogResp.UpdatedAt = val.UpdatedAt
		blogResp.DeletedAt = val.DeletedAt
		blogResp.DeletedBy = val.DeletedBy
		blogList = append(blogList, blogResp)
	}
	return &blogList, nil
}
