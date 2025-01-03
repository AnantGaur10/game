package auth

import (
	"context"
	"net/http"
)

func JwtMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		claims, err := ValidateJWT(cookie.Value)

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"].(float64)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", uint(userID))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
