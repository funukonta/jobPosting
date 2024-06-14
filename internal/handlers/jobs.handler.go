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

// Insert Jobs godoc
// @Summary Insert Job
// @Description Insert job posting
// @Param request body models.Job true "example"
// @Tags Jobs
// @Router /jobs [post]
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

// Search Jobs godoc
// @Summary Search job, can add query parameter
// @Description Search job, can add query parameter
// @Param        keyword    query     string  false  "keyword search"
// @Param        companyName    query     string  false  "companyName search"
// @Tags Jobs
// @Router /jobs [get]
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
