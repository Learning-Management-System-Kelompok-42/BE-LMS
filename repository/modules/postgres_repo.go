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
