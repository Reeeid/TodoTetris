package middleware

import "net/http"

type ctxKey string

type Handler func(http.ResponseWriter, *http.Request)

func AuthJWT(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		/*認証フローは後で実装)*/

	}
}
