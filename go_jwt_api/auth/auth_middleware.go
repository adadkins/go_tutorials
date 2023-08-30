package auth

import (
	"fmt"
	"net/http"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			errMsg := "{\"error\": \"request does not contain an access token\"}"
			fmt.Println(errMsg)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
		err := ValidateToken(token)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.ServeHTTP(w, r)
	})
}
