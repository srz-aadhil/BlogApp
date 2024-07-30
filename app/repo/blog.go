package repo

import (
	"database/sql"
	"fmt"
	"time"
)

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

var _ Repo = (*Blog)(nil)

// Function for reuse table name
func (r *Blog) TableName() string {
	return " blogs "
}

var blog Blog

func (r *Blog) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(title,content,author_id,status,created_by)
			  VALUES ($1, $2,$3,$4,$5)
			  RETURNING id`

	if err := db.QueryRow(query, r.Title, r.Content, r.AuthorID, r.Status, r.CreatedBy).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to : %w", err)
	}
	return lastInsertedID, nil
}

func (r *Blog) Update(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET title=$1,content=$2,updated_at=$3,updated_at=$3,updated_by=$4
		WHERE id=$5
		AND status
		IN(1,2)
		`

	result, err := db.Exec(query, r.Title, r.Content, time.Now().UTC(), r.UpdatedBy, r.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to: %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no blogs with id=%d or status in 1 or 2", r.ID)
	}

	return nil

}

func (r *Blog) Delete(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_by=$1,deleted_at=$2,status=$3
	WHERE id=$4`

	_, err = db.Exec(query, r.DeletedBy, time.Now().UTC(), 3, r.ID)
	if err != nil {
		return fmt.Errorf("delete query execution failed due to: %w", err)
	}
	return nil
}

func (r *Blog) GetOne(db *sql.DB) (result interface{}, err error) {
	query := `SELECT id,title,content,author_id,created_at,updated_at FROM` + r.TableName() + `WHERE id=$1 AND status=2`
	//  var blog Blog
	if err := db.QueryRow(query, r.ID).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return blog, nil
}

func (r *Blog) GetAll(db *sql.DB) (results []interface{}, err error) {
	query := `SELECT id,title,content,author_id,created_at,updated_at
						FROM` + r.TableName() + `` //blogs

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		// 		var blog Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		results = append(results, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return results, nil
}
