package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// UserContextKey はコンテキストにユーザー情報を格納するためのキーです
type UserContextKey string

const userContextKey UserContextKey = "user"

// AuthMiddleware は認証ミドルウェアです
func AuthMiddleware(next http.Handler, jwtSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}

			userID := claims["user_id"].(string)
			// 他のユーザー情報も必要に応じて追加可能
			user := map[string]interface{}{
				"user_id": userID,
				// 他の情報を追加する場合はここに記述
			}

			// コンテキストにユーザー情報を追加
			ctx := context.WithValue(r.Context(), userContextKey, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}

// UserFromContext はコンテキストからユーザー情報を取得します
func UserFromContext(ctx context.Context) (map[string]interface{}, bool) {
	user, ok := ctx.Value(userContextKey).(map[string]interface{})
	return user, ok
}
