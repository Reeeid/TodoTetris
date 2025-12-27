package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

type contextKey string

var UserKey contextKey

func AuthJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized: No token found", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		// 署名検証
		unsignedToken := parts[0] + "." + parts[1]

		h := hmac.New(sha256.New, []byte(os.Getenv(("SECRET_KEY"))))
		h.Write([]byte(unsignedToken))
		expectedSig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

		if expectedSig != parts[2] {
			http.Error(w, "Unauthorized: Is not matched signed Token", http.StatusUnauthorized)
			return
		}

		// ペイロードでここにusernameと有効期限が入ってる
		payloadSegment := parts[1]

		payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var claims map[string]interface{}
		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				http.Error(w, "Unauthorized: Token expired", http.StatusUnauthorized)
				return
			}
		}
		if username, ok := claims["username"]; ok {
			ctx := context.WithValue(r.Context(), UserKey, username)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
}
