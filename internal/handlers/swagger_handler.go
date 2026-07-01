package handlers

import (
	"net/http"

	"projetoweb2/internal/docs"
)

func SwaggerDocHandler(w http.ResponseWriter, r *http.Request) {
	data, err := docs.SwaggerJSON()
	if err != nil {
		http.Error(w, "erro ao carregar documentação swagger", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
