package auth

import (
	"context"
	"fmt"
	"net/http"
)

func JwtMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			fmt.Println("Cookie is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := ValidateJWT(cookie.Value)

		if err != nil {
			fmt.Println("Cookie is invalid")
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["userID"].(float64)
		if !ok {
			fmt.Println("Cookie does not have user Id")
			http.Error(w, "Unauthorized: Invalid User ID", http.StatusUnauthorized)
			return // Stop further execution
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", uint(userID))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
