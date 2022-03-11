package handlers

import (
	"myapp/data"
	"net/http"
	"time"

	"github.com/namnguyen191/goravel"
)

type Handlers struct {
	App    *goravel.Goravel
	Models data.Models
}

func (h *Handlers) Home(rw http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(rw, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}
