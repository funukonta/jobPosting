package routers

import (
	"database/sql"
	"net/http"
	"redikru/internal/handlers"
	"redikru/internal/repositories"
	"redikru/internal/services"
)

func CompanyRouter(m *http.ServeMux, db *sql.DB) {
	repo := repositories.NewCompanyRepo(db)
	serv := services.NewCompanyServ(repo)
	companyHandler := handlers.NewCompanyHandler(serv)

	m.Handle("POST /companies", Handler(companyHandler.Insert))
	m.Handle("GET /companies", Handler(companyHandler.SelectAll))
}
