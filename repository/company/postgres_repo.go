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

// Update this query, because UI/UX already update theyr own data
func (repo *postgreSQLRepository) FindDashboard(companyID string) (domain company.DashboardDomain, err error) {
	// result := repo.db.Table("users").
	// 	Select("users.id AS user_id, companies.id AS company_id, users.full_name AS name_admin, companies.name AS name_company, (SELECT COUNT(id) FROM specializations WHERE company_id = companies.id) AS amount_specialization, (SELECT COUNT(level_access) FROM users WHERE level_access = 'employee') AS amount_employee").
	// 	Joins("INNER JOIN companies ON users.company_id = companies.id").
	// 	Where("companies.id = ?", companyID).
	// 	Order("users.level_access ASC").
	// 	Find(&domain)

	var count int64
	result := repo.db.Table("companies").
		Select("count(specializations.id)").
		Joins("INNER JOIN specializations ON companies.id = specializations.company_id").
		Where("companies.id = ?", companyID).
		Count(&count)
		// Find(&domain)

	fmt.Println("count = ", count)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return domain, exception.ErrNotFound
		}

		return domain, exception.ErrInternalServer
	}

	return domain, nil
}

func (repo *postgreSQLRepository) FindCompanyByID(companyID string) (domain *company.Domain, err error) {
	var company Company

	err = repo.db.Where("id = ?", companyID).First(&company).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrCompanyNotFound
		}

		return nil, exception.ErrInternalServer
	}

	domain = company.ToDomain()

	return domain, nil
}

func (repo *postgreSQLRepository) UpdateProfile(company company.Domain) (id string, err error) {
	newCompany := FromDomain(company)

	err = repo.db.Save(&newCompany).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newCompany.ID

	return id, nil
}

func (repo *postgreSQLRepository) CountSpecializationByCompanyID(companyID string) (count int64, err error) {
	result := repo.db.Table("specializations").
		Select("count(company_id)").
		Where("company_id = ?", companyID).
		Count(&count)

	if result.Error != nil {
		return 0, exception.ErrInternalServer
	}

	return count, nil
}

func (repo *postgreSQLRepository) CountEmployeeByCompanyID(companyID string) (count int64, err error) {
	result := repo.db.Table("users").
		Select("count(company_id)").
		Where("company_id = ?", companyID).
		Count(&count)

	if result.Error != nil {
		return 0, exception.ErrInternalServer
	}

	return count, nil
}

func (repo *postgreSQLRepository) CountCourseByCompanyID(companyID string) (count int64, err error) {
	result := repo.db.Table("courses").
		Select("count(company_id)").
		Where("company_id = ?", companyID).
		Count(&count)

	if result.Error != nil {
		return 0, exception.ErrInternalServer
	}

	return count, nil
}

func (repo *postgreSQLRepository) FindSpecializationEmployee(companyID string) (domain []company.DashboardCompanySpecialization, err error) {
	result := repo.db.Table("specializations").
		Limit(4).
		Select("specializations.id, specializations.company_id, specializations.name, specializations.invitation, specializations.created_at, specializations.updated_at, count(users.id) as amount_employee").
		Joins("LEFT JOIN users ON specializations.id = users.specialization_id").
		Where("specializations.company_id = ?", companyID).
		Group("specializations.id").
		Order("amount_employee desc").
		Find(&domain)

	if result.Error != nil {
		return nil, exception.ErrInternalServer
	}

	return domain, nil
}

func (repo *postgreSQLRepository) FindCourseEmployee(companyID string) (domain []company.DashboardCompanyCourse, err error) {
	result := repo.db.Table("courses").
		Limit(4).
		Select("courses.id, courses.company_id, courses.title, courses.created_at, courses.updated_at, count(user_courses.user_id) as amount_employee").
		Joins("LEFT JOIN user_courses ON courses.id = user_courses.course_id").
		Where("courses.company_id = ?", companyID).
		Group("courses.id").
		Order("amount_employee desc").
		Find(&domain)

	if result.Error != nil {
		return nil, exception.ErrInternalServer
	}

	return domain, nil
}

func (repo *postgreSQLRepository) FindCompanyByCompanyID(companyID string) (domain company.DashboardCompanyAdmin, err error) {
	result := repo.db.Table("companies").
		Select("users.id as user_id, companies.id as company_id, users.full_name as name_admin, companies.name as name_company").
		Joins("INNER JOIN users ON companies.id = users.company_id").
		Where("companies.id = ?", companyID).
		Find(&domain)

	if result.Error != nil {
		return domain, exception.ErrInternalServer
	}

	return domain, nil
}
