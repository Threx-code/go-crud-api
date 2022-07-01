package users

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Threx-code/go-api/config"
	"github.com/Threx-code/go-api/middlewares"
	"github.com/Threx-code/go-api/utils"
	"github.com/go-playground/validator/v10"
)

func (req *LoginRequest) LoginUser(w http.ResponseWriter) {

	validate := validator.New()
	err := validate.Struct(req)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgFortag(fe)}
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(out)
			return
		}
	}

	config.Connect()
	db := config.GetDB()
	user := &Users{}

	query := "SELECT id, email, password, firstname, lastname, created_at, updated_at FROM users WHERE email = ? "
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	row := stmt.QueryRowContext(ctx, req.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.CreatedAt, &user.UpdatedAt)
	if row != nil {
		resErr := &ResError{
			Error: "invalid login credentials",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resErr)
		return
	}

	err = utils.VerifyPassword(user.Password, req.Password)
	if err != nil {
		resErr := &ResError{
			Error: "invalid login credentials",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resErr)
		return
	}

	token, err := middlewares.GenerateToken(user.Email, time.Minute*15)

	if err != nil {
		resErr := &ResError{
			Error: "invalid login credentials",
		}
		json.NewEncoder(w).Encode(resErr)
		return
	}

	newUser := &LoginResponse{
		User:        NewUserResponse(user),
		AccessToken: token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)

}

func msgFortag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	}
	return fe.Error()
}
