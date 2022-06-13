package company

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) company.CompanyRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

func (repo *postgreSQLRepository) Insert(company company.Domain) (id string, err error) {
	newCompany := FromDomain(company)
	err = repo.db.Create(&newCompany).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newCompany.ID

	return id, nil
}

func (repo *postgreSQLRepository) CheckWeb(web string) error {
	var company Company
	fmt.Println("web = ", web)
	err := repo.db.Where("web = ?", web).First(&company).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrWebExists
}
