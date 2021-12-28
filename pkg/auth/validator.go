package auth

import (
	"net/http"
	"strings"
)

// ValidateUser checks for a token and validates it
// After that execute the f function received
func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		tknStr := splitToken[1]

		tokenMaker, err := NewJWTMaker(SecretTokenKey)
		if err != nil {
			return
		}

		payload, err := tokenMaker.VerifyToken(tknStr)
		if err != nil || payload == nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		 	return
		}

		f(w, r)
	}
}


	

