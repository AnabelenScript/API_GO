package http_handlers

import (
	"api_prueba/users/application"
	"encoding/json"
	"net/http"
)

//Manejo de datos

type UserHandler struct {
	Service *application.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Entrada invalida", http.StatusBadRequest)
		return
	}
	err := h.Service.CreateUser(input.Name, input.Email)
	if err != nil {
		http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuario creado :)"))
}
