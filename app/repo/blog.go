package repo

import (
	"blog/app/dto"
	"database/sql"
	"fmt"
	"time"
)

type BlogRepo interface {
	Create(blogReq *dto.BlogCreateRequest) (lastinsertedID int64, err error)
	Update(blogReq *dto.BlogUpdateRequest) error
	Delete(blogReq *dto.BlogDeleteRequest) error
	Getblog(blogReq *dto.BlogRequest) (blogResp *dto.BlogResponse, err error)
	GetBlogs() (blogResp *[]dto.BlogResponse, err error)
	TableName() string //function for reuse table
}

// Blog Model
type Blog struct {
	ID       uint16
	Title    string
	Content  string
	Status   int    // 1- Draft, 2 - Published, 3 - Deleted
	AuthorID uint16 // Author.ID
	Model
	DeleteInfo
}

type BlogRepoImpl struct {
	db *sql.DB
}

func NewBlogRepo(db *sql.DB) BlogRepo {
	return &BlogRepoImpl{
		db: db,
	}
}

var _ BlogRepo = (*BlogRepoImpl)(nil)

// Function for reuse table name
func (r *BlogRepoImpl) TableName() string {
	return " blogs "
}

func (r *BlogRepoImpl) Create(blogReq *dto.BlogCreateRequest) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(title,content,author_id,status,created_by)
			  VALUES ($1, $2,$3,$4,$5)
			  RETURNING id`

	if err := r.db.QueryRow(query, blogReq.Title, blogReq.Content, blogReq.AuthorID, blogReq.Status, blogReq.CreatedBy).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to : %w", err)
	}
	return lastInsertedID, nil
}

func (r *BlogRepoImpl) Update(blogReq *dto.BlogUpdateRequest) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET title=$1,content=$2,updated_at=$3,updated_by=$4
		WHERE id=$5
		AND status
		IN(1,2)
		`

	result, err := r.db.Exec(query, blogReq.Title, blogReq.Content, time.Now().UTC(), blogReq.UpdatedBy, blogReq.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to: %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no blogs with id=%d or status in 1 or 2", blogReq.ID)
	}

	return nil

}

func (r *BlogRepoImpl) Delete(blogReq *dto.BlogDeleteRequest) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_by=$1,deleted_at=$2,status=$3
	WHERE id=$4`
	// var blog Blog
	_, err = r.db.Exec(query, blogReq.DeletedBy, time.Now().UTC(), 3, blogReq.ID)
	if err != nil {
		return fmt.Errorf("delete query execution failed due to: %w", err)
	}
	return nil
}

func (r *BlogRepoImpl) Getblog(blogReq *dto.BlogRequest) (blogResp *dto.BlogResponse, err error) {
	query := `SELECT id,title,content,author_id,created_at,updated_at,updated_by,status FROM` + r.TableName() + `WHERE id=$1`
	blogResp = &dto.BlogResponse{}
	if err := r.db.QueryRow(query, blogReq.ID).Scan(&blogResp.ID, &blogResp.Title, &blogResp.Content, &blogResp.AuthorID, &blogResp.CreatedAt, &blogResp.UpdatedAt, &blogResp.UpdatedBy, &blogResp.Status); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return blogResp, nil
}

func (r *BlogRepoImpl) GetBlogs() (blogResp *[]dto.BlogResponse, err error) {
	query := `SELECT id,title,content,author_id,created_at,updated_at,updated_by
						FROM` + r.TableName() + `` //blogs

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %s", err)
	}
	defer rows.Close()
	var collection []dto.BlogResponse
	for rows.Next() {
		var blog dto.BlogResponse
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt, &blog.UpdatedBy); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		collection = append(collection, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return &collection, nil
}
