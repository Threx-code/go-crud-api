package middlewares

import (
	"time"

	"github.com/Threx-code/go-api/token"
)

var tokenType, err = token.NewPasetoMaker("12345456787890987654321234567890")

func GenerateToken(email string, duration time.Duration) (string, error) {
	//maker, err := token.NewPasetoMaker("12345456787890987654321234567890")
	if err != nil {
		return "", err
	}

	token, err := tokenType.CreateToken(email, duration)
	if err != nil {
		return "", err
	}

	return string(token), nil

}

func TokenVerification(payload string) (*token.Payload, error) {
	//maker, err := token.NewPasetoMaker("12345456787890987654321234567890")
	if err != nil {
		return nil, err
	}

	tokenPayload, err := tokenType.VerifyToken(payload)
	if err != nil {
		return nil, err
	}

	return tokenPayload, nil
}
