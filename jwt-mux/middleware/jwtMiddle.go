package middleware

import (
	"fmt"
	"jwt-mux/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTmiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Fprintf(w, "error cookie: %s", err.Error())
				return
			}
		}

		// mengambil token value
		tokenString := cookie.Value
		claims := &config.JWTClaim{}

		// parsing token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			fmt.Fprint(w, "err parsing token")
			return
		}

		if !token.Valid {
			fmt.Fprint(w, "err token invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
