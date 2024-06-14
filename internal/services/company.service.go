package services

import (
	"fmt"
	"redikru/internal/models"
	"redikru/internal/repositories"

	"github.com/google/uuid"
)

type CompanyService interface {
	Insert(string) error
	SelectAll() ([]models.Company, error)
}

type companyService struct {
	repo repositories.CompanyRepo
}

func NewCompanyServ(repo repositories.CompanyRepo) CompanyService {
	return &companyService{repo: repo}
}

func (s companyService) Insert(name string) error {
	company := &models.Company{
		ID:   uuid.New(),
		Name: name,
	}

	err := s.repo.Insert(company)
	return err
}

func (s companyService) SelectAll() ([]models.Company, error) {
	companies, err := s.repo.SelectAll()
	if err != nil {
		return nil, err
	}

	if len(companies) == 0 {
		return nil, fmt.Errorf("tidak ada data companies")
	}

	return companies, err
}
