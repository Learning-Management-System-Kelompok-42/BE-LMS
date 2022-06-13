package company

import (
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) domain.CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (repo *companyRepository) Insert(company domain.Domain) (id string, err error) {
	err = repo.db.Create(&company).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = company.ID

	return id, nil
}

func (repo *companyRepository) CheckEmail(email string) error {
	var company domain.Domain
	err := repo.db.Where("email = ?", email).First(&company).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrEmailExists
}
