package handler

import (
	"net/http"
)

func RoutePerson(mux *http.ServeMux, dbHandler dbHandler) {
	hdler := newPerson(dbHandler)

	mux.HandleFunc("/people/create", hdler.create)
	mux.HandleFunc("/people/get-all", hdler.getAll)

}
