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
	err := repo.db.Where("web = ?", web).First(&company).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrWebExists
}

func (repo *postgreSQLRepository) FindDashboard(companyID string) (domain company.DashboardDomain, err error) {
	result := repo.db.Raw(`SELECT U.id AS user_id, CO.id AS company_id, U.full_name AS name_admin, CO.name AS name_company, (SELECT COUNT(id) FROM specializations WHERE company_id = CO.id) AS amount_specialization, (SELECT COUNT(level_access) FROM users WHERE level_access = 'employee') AS amount_employee 
	FROM users U
	INNER JOIN companies CO ON U.company_id = CO.id
	WHERE CO.id = ?`, companyID).Find(&domain)

	fmt.Println("domain = ", domain)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return domain, exception.ErrNotFound
		}

		return domain, exception.ErrInternalServer
	}

	fmt.Println("Masuk repo")

	return domain, nil
}
