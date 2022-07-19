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
	var module Module

	err = repo.db.Where("id = ?", id).First(&module).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain, exception.ErrNotFound
		}

		return domain, exception.ErrInternalServer
	}

	domain = module.ToDomain()

	return domain, nil
}

func (repo *postgreSQLRepository) FindAllModuleByCourseID(courseID string) (modules []module.Domain, err error) {
	var modulesDB []Module

	err = repo.db.Where("course_id = ?", courseID).
		Order("orders ASC").
		Find(&modulesDB).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return modules, exception.ErrModuleNotFound
		}

		return modules, exception.ErrInternalServer
	}

	modules = ToDomainBatchList(modulesDB)

	return modules, nil
}
