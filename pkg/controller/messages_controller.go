package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/persistence"
	"gopkg.in/go-playground/validator.v9"
)

// SendMessage send a message from one user to another
func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var newMessageQuery models.Message
	
	if err := json.NewDecoder(r.Body).Decode(&newMessageQuery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	v := validator.New()
	if errValid := v.Struct(newMessageQuery); errValid != nil {
		http.Error(w, errValid.Error(), http.StatusBadRequest)
		return
	}
	
	newMsg, errMsg := persistence.CreateMessage(newMessageQuery)
	if errMsg != nil {
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
		return
	}
	helpers.RespondJSON(w, map[string]interface{}{"id":newMsg.ID, "timestamp":newMsg.LastUpdated})
}

// GetMessages get the messages from the logged user to a recipient
func (h Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages

	recipient := r.URL.Query().Get("recipient")
	start := r.URL.Query().Get("start")
	limit := r.URL.Query().Get("limit")

	if recipient == "" || start == "" {
		http.Error(w, "recipient or start param missing", http.StatusBadRequest)
		return
	}

	recipientInt,_ := strconv.Atoi(recipient)
	startInt,_ := strconv.Atoi(start)
	limitInt,_ := strconv.Atoi(limit)

	if limit == "" {
		limit = "5"
	}
	messages, errGetAll := persistence.GetAllMessages(recipientInt, startInt, limitInt)
	if errGetAll != nil {
		http.Error(w, errGetAll.Error(), http.StatusInternalServerError)
	}

	helpers.RespondJSON(w, messages)
}
