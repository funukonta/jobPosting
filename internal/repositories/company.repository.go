package repositories

import (
	"database/sql"
	"redikru/internal/models"
)

type CompanyRepo interface {
	Insert(*models.Company) error
	SelectAll() ([]models.Company, error)
}

type companyRepo struct {
	*sql.DB
}

func NewCompanyRepo(db *sql.DB) CompanyRepo {
	return &companyRepo{db}
}

func (r companyRepo) Insert(data *models.Company) error {
	query := "INSERT into companies (id,name) values ($1,$2)"
	_, err := r.DB.Exec(query, data.ID, data.Name)
	return err
}

func (r companyRepo) SelectAll() ([]models.Company, error) {
	query := `select * from companies`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := []models.Company{}
	for rows.Next() {
		company := models.Company{}
		err = rows.Scan(&company.ID, &company.Name)
		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, err
}
