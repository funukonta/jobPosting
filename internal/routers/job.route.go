package routers

import (
	"database/sql"
	"net/http"
	"redikru/internal/handlers"
	"redikru/internal/repositories"
	"redikru/internal/services"
)

func JobsRouter(m *http.ServeMux, db *sql.DB) {
	repo := repositories.NewJobsRepo(db)
	serv := services.NewJobsServ(repo)
	jobHandler := handlers.NewJobsHandler(serv)

	m.Handle("POST /jobs", Handler(jobHandler.Insert))
	m.Handle("GET /jobs", Handler(jobHandler.Select))
}
