package controller

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	return r
}
