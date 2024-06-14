package routers

import (
	"database/sql"
	"net/http"
	"redikru/pkg"
)

func Routes(m *http.ServeMux, db *sql.DB) {
	JobsRouter(m, db)
	CompanyRouter(m, db)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		pkg.Response(400, &pkg.JsonBod{Message: err.Error()}).Send(w)
	}
}
