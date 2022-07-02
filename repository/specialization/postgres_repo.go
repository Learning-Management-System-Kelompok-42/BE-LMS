package specialization

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) specialization.SpecializationRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(specialization specialization.Domain) (id string, err error) {
	newSpec := FromDomain(specialization)

	err = repo.db.Create(&newSpec).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = specialization.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindInvitation(invitation string) (specialization.Domain, error) {
	var spec Specialization

	err := repo.db.Where("invitation = ?", invitation).First(&spec).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return specialization.Domain{}, exception.ErrNotFound
		}
		return specialization.Domain{}, exception.ErrInternalServer
	}

	return spec.ToDomain(), nil
}

func (repo *postgreSQLRepository) FindDashboardSpecialization(companyID string) (specializations []specialization.Domain, err error) {
	var specs []Specialization
	err = repo.db.Where("company_id = ?", companyID).Find(&specs).Error

	if err != nil {
		if len(specializations) == 0 {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternalServer
	}

	specializations = ToDomainList(specs)

	return specializations, nil
}

func (repo *postgreSQLRepository) CountCourse(specID string) (result int64, err error) {
	err = repo.db.Table("specialization_courses").Where("specialization_id = ?", specID).Count(&result).Error
	if err != nil {
		return 0, exception.ErrInternalServer
	}

	return result, nil
}

func (repo *postgreSQLRepository) CountEmployee(companyID, specID string) (result int64, err error) {
	err = repo.db.Table("users").Where("company_id = ? AND specialization_id = ?", companyID, specID).Count(&result).Error
	if err != nil {
		return 0, exception.ErrInternalServer
	}

	return result, nil
}

func (repo *postgreSQLRepository) CheckLinkInviation(link string) (err error) {
	var spec Specialization

	err = repo.db.Where("invitation = ?", link).First(&spec).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return nil
}
