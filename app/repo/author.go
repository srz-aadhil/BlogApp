package repo

import (
	"database/sql"
	"fmt"
	"time"
)

// Author Model
type Author struct {
	ID   uint16
	Name string
	Model
	DeleteInfo
}

var _ Repo = (*Author)(nil)

func (r *Author) TableName() string {
	return " authors "
}

func (r *Author) Create(db *sql.DB) (lastInsertedID int64, err error) {

	query := `INSERT INTO` + r.TableName() + `(name)
						VALUES ($1)
						RETURNING id`

	if err := db.QueryRow(query, r.Name).Scan(&lastInsertedID); err != nil {
		return 0, fmt.Errorf("couldn't get last inserted id due to: %w", err)
	}
	return lastInsertedID, nil
}

func (r *Author) Update(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET name= $1,updated_at=$2,updated_by=$3
				WHERE id=$4`

	result, err := db.Exec(query, r.Name, time.Now().UTC(), r.UpdatedBy, r.ID)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d ", r.ID)
	}
	return nil
}

func (r *Author) Delete(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET deleted_at=$1,deleted_by=$2
		      WHERE id=$3`

	_, err = db.Exec(query, time.Now().UTC(), r.DeletedBy, r.ID)
	if err != nil {
		return fmt.Errorf("update query failed due to : %w", err)
	}
	return nil
}

var author Author

func (r *Author) GetOne(db *sql.DB) (result interface{}, err error) {
	query := `SELECT id,name,created_at,updated_at
			  FROM` + r.TableName() +
		`WHERE id=$1`

	if err := db.QueryRow(query, r.ID).Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt); err != nil {
		return nil, fmt.Errorf("query failed due to : %w", err)
	}
	return author, nil
}

func (r *Author) GetAll(db *sql.DB) (results []interface{}, err error) {
	query := `SELECT id,name,created_at,updated_at
	         FROM ` + r.TableName() + ``

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		// var author Author
		if err := rows.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		results = append(results, author)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return results, nil
}
