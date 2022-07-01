package middlewares

import (
	"net/http"
	"strings"

	"github.com/Threx-code/go-api/exceptions"
)

const (
	headerKey  = "authorization"
	typeKey    = "token"
	payloadKey = "authorization_payload"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(headerKey)
		if len(authorizationHeader) == 0 {
			exceptions.NewValidationError(w, http.StatusUnauthorized, "authorization header is missing")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			exceptions.NewValidationError(w, http.StatusUnauthorized, "invalid authorization header format")
			return
		}

		if typeKey != strings.ToLower(fields[0]) {
			exceptions.NewValidationError(w, http.StatusUnauthorized, "unsupported authorization type, use "+typeKey)
			return
		}

		/* validate */
		authToken := fields[1]
		_, err := TokenVerification(authToken)
		if err != nil {
			exceptions.NewValidationError(w, http.StatusUnauthorized, err.Error())
		}

	}
}
