package middleware

import "net/http"

func AuthJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*認証フローは後で実装*/
	}
}
