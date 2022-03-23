package handlers

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeHomeHandler(r *mux.Router, n *negroni.Negroni) {
	r.Handle("/v1", n.With(
		negroni.Wrap(getHome()),
	)).Methods("GET", "OPTIONS")
}

/*
Para testar:
curl ' http://localhost:4000/v1
*/
func getHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello World"}`))
	})
}
