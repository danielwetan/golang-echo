package controllers

import (
	"net/http"

	"github.com/danielwetan/golang-mvc/src/responses"
)

// Home ...
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hello there!")
}

