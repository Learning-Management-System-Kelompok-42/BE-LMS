package module

import (
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) module.ModuleRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(domain module.Domain) (id string, err error) {
	newModule := FromDomain(domain)
	err = repo.db.Create(&newModule).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newModule.ID

	return id, nil
}

func (repo *postgreSQLRepository) Update(domain module.Domain) (id string, err error) {
	updateModule := FromDomain(domain)

	err = repo.db.Where("id = ?", domain.ID).Save(&updateModule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", exception.ErrNotFound
		}

		return "", exception.ErrInternalServer
	}

	id = updateModule.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindByID(id string) (domain module.Domain, err error) {
	returnModule := FromDomain(module.Domain{ID: id})

	err = repo.db.Where("id = ?", id).First(&returnModule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain, exception.ErrNotFound
		}

		return domain, exception.ErrInternalServer
	}

	domain = returnModule.ToDomain()

	return domain, nil
}
