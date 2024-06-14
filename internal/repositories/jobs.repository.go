package repositories

import (
	"database/sql"
	"fmt"
	"redikru/internal/models"
	"strings"
)

type JobsRepo interface {
	Insert(*models.Job) error
	Select(string, string) ([]models.JobCompany, error)
}

type jobsRepo struct {
	*sql.DB
}

func NewJobsRepo(db *sql.DB) JobsRepo {
	return &jobsRepo{db}
}

func (r jobsRepo) Insert(data *models.Job) error {
	query := `INSERT into jobs (id,company_id,title,description) values ($1,$2,$3,$4)`
	_, err := r.DB.Exec(query, data.ID, data.CompanyID, data.Title, data.Description)
	if err != nil {
		return err
	}

	return err
}

func (r jobsRepo) Select(keyword, companyName string) ([]models.JobCompany, error) {
	query := `SELECT jobs.id, jobs.company_id, jobs.title, jobs.description, jobs.created_at, companies.name as company_name
              FROM jobs
              JOIN companies ON jobs.company_id = companies.id`

	if keyword != "" {
		query += ` AND (jobs.title ILIKE :keyword OR jobs.description ILIKE :keyword)`
		key := fmt.Sprintf("'%%%s%%'", keyword)
		query = strings.ReplaceAll(query, ":keyword", key)
	}
	if companyName != "" {
		query += ` AND companies.name ILIKE :compname`
		key := fmt.Sprintf("'%%%s%%'", companyName)
		query = strings.ReplaceAll(query, ":compname", key)
	}

	query += ` ORDER BY jobs.created_at DESC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobs := []models.JobCompany{}

	for rows.Next() {
		job := models.JobCompany{}
		err := rows.Scan(&job.ID, &job.CompanyID, &job.Title, &job.Description, &job.CreatedAt, &job.CompanyName)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, err
}
