package controller

import (
	"net/http"

	"github.com/nahuelmarianolosada/go-boilerplate/pkg/helpers"
	"github.com/nahuelmarianolosada/go-boilerplate/pkg/models"
	"github.com/nahuelmarianolosada/go-boilerplate/pkg/persistence"
)

// Check returns the health of the service and DB
func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	// TODO: Check service health. Feel free to add any check you consider necessary
	if err := persistence.GetConnection().Ping(); err != nil {
		helpers.RespondJSON(w, models.Health{Health: "not ok - " + err.Error()})
		return 
	}
	helpers.RespondJSON(w, models.Health{Health: "ok"})
}
