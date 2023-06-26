package account

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httpstransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httpstransport.NewServer(endpoints.Create, decodeUserReq, encodeReponse))
	r.Methods("GET").Path("/user{id}").Handler(httpstransport.NewServer(endpoints.Get, decodeEmailReq, encodeReponse))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
