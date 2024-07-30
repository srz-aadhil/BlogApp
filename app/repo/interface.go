package repo

import "database/sql"

// type Repo interface {
// 	Create(db *sql.DB) (lastInsertedID int64, err error)
// 	Update(db *sql.DB) (err error)
// 	Delete(db *sql.DB) (err error)
// 	GetOne(db *sql.DB) (result interface{}, err error)
// 	GetAll(db *sql.DB) (results []interface{}, err error)
// 	TableName() string //function for reuse table
// }

type Repo interface {
	Create(db *sql.DB) (lastInsertedID int64, err error)
	Update(db *sql.DB) (err error)
	Delete(db *sql.DB) (err error)
	GetOne(db *sql.DB) (result interface{}, err error)
	GetAll(db *sql.DB) (results []interface{}, err error)
	TableName() string //function for reuse table
}
