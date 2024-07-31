package repo

import (
	"blog/pkg/salthash"
	"database/sql"
	"fmt"
	"time"
)

// User Model
type User struct {
	ID        uint16
	UserName  string
	Password  string
	Salt      string
	IsDeleted bool
	Model
	DeleteInfo
}

var _ Repo = (*User)(nil)

var user User

func (r *User) TableName() string {
	return " users "
}
func (r *User) Create(db *sql.DB) (lastInsertID int64, err error) {
	// Generate Salt
	salt, err := salthash.GenerateSalt(10)
	if err != nil {
		return 0, fmt.Errorf("error generating salt: %w", err)
	}
	// Hash Password
	PasswordString := salthash.HashPassword(r.Password, salt)

	query := `INSERT INTO` + r.TableName() + `(username,password,salt)
						VALUES ($1,$2,$3)
						RETURNING id`

	if err := db.QueryRow(query, r.UserName, PasswordString, salt).Scan(&lastInsertID); err != nil {
		return 0, fmt.Errorf("query execution failed due to : %w", err)
	}

	return lastInsertID, nil
}

func (r *User) Update(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET username=$1,password=$2,updated_at=$3
			  			WHERE id=$4`

	result, err := db.Exec(query, r.UserName, salthash.HashPassword(r.Password, r.Salt), time.Now().UTC(), r.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to : %w", err)
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d", r.ID)
	}
	return nil
}

func (r *User) Delete(db *sql.DB) (err error) {
	query := `UPDATE` + r.TableName() +
		`SET is_deleted=$1,deleted_at=$2
				WHERE id = $3`

	_, err = db.Exec(query, true, time.Now().UTC(), r.ID)
	if err != nil {
		return fmt.Errorf("query execution failed due to: %w", err)
	}
	return nil
}
func (r *User) GetAll(db *sql.DB) (results []interface{}, err error) {
	query := `SELECT id,username,password,created_at,updated_at
			  FROM` + r.TableName() + `WHERE is_deleted=false`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		// var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		results = append(results, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return results, nil
}

func (r *User) GetOne(db *sql.DB) (result interface{}, err error) {
	query := `SELECT id,username,password,created_at,updated_at
						FROM` + r.TableName() +
		`WHERE id=$1
						AND
						is_deleted=false`
	// var user User
	if err := db.QueryRow(query, r.ID).Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("query execution failed due to : %w", err)
	}
	return user, nil

}
