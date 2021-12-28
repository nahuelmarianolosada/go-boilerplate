package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-boilerplate/pkg/auth"
	"github.com/go-boilerplate/pkg/helpers"
	"github.com/go-boilerplate/pkg/models"
	"github.com/go-boilerplate/pkg/persistence"
	"gopkg.in/go-playground/validator.v9"
)

// Login authenticates a user and returns a token
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: User must login and a token must be generated

	// First of all, we need to search the user. 
	// After that, we generate a new token
	var newLogin models.Login

	if err := json.NewDecoder(r.Body).Decode(&newLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	if errValid := v.Struct(newLogin); errValid != nil {
		http.Error(w, errValid.Error(), http.StatusBadRequest)
		return
	}

	if usFinded, err := persistence.GetByUsername(newLogin.Username); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	} else {

		tokenMaker, err := auth.NewJWTMaker(auth.SecretTokenKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		 	return
		}

		ss, err := tokenMaker.CreateToken(newLogin.Username, 5 * time.Minute)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		 	return
		}
		
		fmt.Printf("Token for the user %s [%s]", usFinded.ID, ss)
		helpers.RespondJSON(w, map[string]string{"id": usFinded.ID, "token": ss})
	}
}
