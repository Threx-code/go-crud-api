package users

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
	"unicode"

	"github.com/Threx-code/go-api/config"
	"github.com/Threx-code/go-api/utils"
	"github.com/go-playground/validator/v10"
)

func (req *CreateRequest) CreateNewUser(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe)}
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(out)
			return
		}
	}

	if !IsAValidPassword(req.Password) {
		resErr := &ResError{
			Error: "invalid password it should contain at least one uppercase, lowercase, number and symbol",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resErr)
		return
	}

	config.Connect()
	db := config.GetDB()

	// check if email exist
	if IsEmailAvailable(req.Email) > 0 {
		resErr := &ResError{
			Error: "this email already taken",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resErr)
		return
	}

	query := "INSERT INTO users ( email, password, firstname, lastname ) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		// error that could not prepare query
		return
	}

	// hash the password

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		// could not hash the password
		return
	}

	row, err := stmt.ExecContext(ctx, req.Email, hashedPassword, req.Firstname, req.Lastname)

	if err != nil {
		// could not insert the password
		return
	}

	id, _ := row.LastInsertId()

	createdUser := &UserResponse{
		ID:        id,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		CreatedAt: time.Now(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email address"
	case "alpha":
		return "only alphabec allowed in this field"
	case "min":
		return "the minimum value is"
	}
	return fe.Error()
}

func IsAValidPassword(password string) bool {
	var (
		hasMinLength = false
		hasUppercase = false
		hasLowercase = false
		hasSymbol    = false
		hasNumeric   = false
	)

	if len(password) >= 6 {
		hasMinLength = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumeric = true
		case unicode.IsPunct(char):
			hasSymbol = true
		}
	}

	return hasMinLength && hasUppercase && hasLowercase && hasNumeric && hasSymbol
}

func IsEmailAvailable(email string) int {
	var num int
	config.Connect()
	db := config.GetDB()
	query := "SELECT count(id) FROM users WHERE email = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	row := stmt.QueryRowContext(ctx, email)
	row.Scan(&num)
	return num
}
