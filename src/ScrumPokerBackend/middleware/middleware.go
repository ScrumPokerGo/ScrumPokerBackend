package middleware

import (
	"net/http"
)

func ResponseJsonMiddleware(h func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-type", `application/json`)
		return h(w, r)
	}
}

func MiddlewareSet(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	middleware_stack := ResponseJsonMiddleware(h)
	return MiddlewareError(middleware_stack)

}
