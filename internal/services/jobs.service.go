package services

import (
	"fmt"
	"redikru/internal/models"
	"redikru/internal/repositories"

	"github.com/google/uuid"
)

type JobsService interface {
	Insert(*models.Job) error
	Select(string, string) ([]models.JobCompany, error)
}

type jobsService struct {
	repo repositories.JobsRepo
}

func NewJobsServ(repo repositories.JobsRepo) JobsService {
	return &jobsService{repo: repo}
}

func (s jobsService) Insert(data *models.Job) error {
	data.ID = uuid.New()

	err := s.repo.Insert(data)
	return err
}

func (s jobsService) Select(keyword, companyName string) ([]models.JobCompany, error) {
	jobs, err := s.repo.Select(keyword, companyName)
	if err != nil {
		return nil, err
	}

	if len(jobs) == 0 {
		return nil, fmt.Errorf("tidak ada data jobs")
	}

	return jobs, err
}
