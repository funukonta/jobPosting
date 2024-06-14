package handlers

import (
	"net/http"
	"redikru/internal/models"
	"redikru/internal/services"
	"redikru/pkg"
)

type JobsHandler interface {
	Insert(w http.ResponseWriter, r *http.Request) error
	Select(w http.ResponseWriter, r *http.Request) error
}

type jobsHandler struct {
	serv services.JobsService
}

func NewJobsHandler(serv services.JobsService) JobsHandler {
	return &jobsHandler{serv: serv}
}

func (h jobsHandler) Insert(w http.ResponseWriter, r *http.Request) error {
	job := &models.Job{}
	err := pkg.GetJsonBody(r, job)
	if err != nil {
		return err
	}

	err = h.serv.Insert(job)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil posting job"}).Send(w)
	return nil
}

func (h jobsHandler) Select(w http.ResponseWriter, r *http.Request) error {
	keyword := r.URL.Query().Get("keyword")
	companyName := r.URL.Query().Get("companyName")

	jobs, err := h.serv.Select(keyword, companyName)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil ambil jobs", Data: jobs}).Send(w)
	return nil
}
