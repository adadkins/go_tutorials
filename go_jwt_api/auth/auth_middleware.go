package auth

import (
	"fmt"
	"net/http"
)

func Auth(h http.Handler) http.Handler {
	fmt.Println("Are we in here?")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We are in Auth middleware")
		token := r.Header.Get("Authorization")
		err := ValidateToken(token)
		if err != nil {
			return
		}
		h.ServeHTTP(w, r)
	})
}

// if tokenString == "" {
// 	//context.JSON(401, gin.H{"error": "request does not contain an access token"})
// 	//context.Abort()
// 	return
// }
// err := ValidateToken(tokenString)
// if err != nil {
// 	//context.JSON(401, gin.H{"error": err.Error()})
// 	//context.Abort()
// 	return
// }
