package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Books struct {
	ID        int       `json:"id,omitempty"`
	TITLE     string    `validate:"required,string"`
	CONTENT   string    `validate:"required,string"`
	USERID    string    `validate:"required,string"`
	CREATEDAT time.Time `json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at"`
}

// func (book *Books) CreateBook() (*Books, error) {
// 	validate := validator.New()
// 	err := validate.Struct(book)
// 	if err != nil {
// 		return nil, err
// 	}

// 	config.Connect()
// 	db := config.GetDB()

// 	query := "INSERT INTO books (title, content, user_id) VALUES (:title, :content, :user_id)"
// 	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancelfunc()

// 	stmt, err := db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := stmt.ExecContext(ctx, book.TITLE, book.CONTENT, book.USERID)

// }
