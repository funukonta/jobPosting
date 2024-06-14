package handlers

import (
	"net/http"
	"redikru/internal/models"
	"redikru/internal/services"
	"redikru/pkg"
)

type CompanyHandler interface {
	Insert(w http.ResponseWriter, r *http.Request) error
	SelectAll(w http.ResponseWriter, r *http.Request) error
}

type companyHandler struct {
	serv services.CompanyService
}

func NewCompanyHandler(serv services.CompanyService) CompanyHandler {
	return &companyHandler{serv: serv}
}

// Insert Company godoc
// @Summary Insert Company
// @Description Insert Company
// @Param request body models.Company true "example"
// @Tags Companies
// @Router /companies [post]
func (h companyHandler) Insert(w http.ResponseWriter, r *http.Request) error {
	company := &models.Company{}
	err := pkg.GetJsonBody(r, company)
	if err != nil {
		return err
	}

	err = h.serv.Insert(company.Name)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil insert company"}).Send(w)
	return nil
}

// Show Company godoc
// @Summary Show Company
// @Description Show Company
// @Tags Companies
// @Router /companies [get]
func (h companyHandler) SelectAll(w http.ResponseWriter, r *http.Request) error {
	companies, err := h.serv.SelectAll()
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil mendapatkan companies", Data: companies}).Send(w)
	return nil
}
