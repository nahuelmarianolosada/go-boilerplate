package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nahuelmarianolosada/go-boilerplate/pkg/helpers"
	"github.com/nahuelmarianolosada/go-boilerplate/pkg/models"
	"github.com/nahuelmarianolosada/go-boilerplate/pkg/persistence"
	"gopkg.in/go-playground/validator.v9"
)

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Create a New User
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	if errValid := v.Struct(newUser); errValid != nil {
		http.Error(w, errValid.Error(), http.StatusBadRequest)
		return
	}

	insertedID, err := persistence.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.RespondJSON(w, map[string]int64{"id": insertedID})
}
