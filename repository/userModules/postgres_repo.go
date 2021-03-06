package userModules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) userModules.UserModulesRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) UpdateProgress(progress userModules.Domain) (id string, err error) {
	newProgress := FromDomain(progress)

	err = repo.db.Save(&newProgress).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = progress.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindProgressByModuleID(moduleID string) (progress *userModules.Domain, err error) {
	var modules UserModule

	err = repo.db.Where("module_id = ?", moduleID).First(&modules).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrModuleNotFound
		}

		return nil, exception.ErrInternalServer
	}

	return modules.ToDomain(), nil
}

func (repo *postgreSQLRepository) CheckProgressExist(userID, courseID, moduleID string) (err error) {
	var progress UserModule

	err = repo.db.Where("user_id = ? AND course_id = ? AND module_id = ?", userID, courseID, moduleID).
		First(&progress).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}
	return exception.ErrProgressAlreadyExist
}

func (repo *postgreSQLRepository) InsertProgress(progress userModules.Domain) (id string, err error) {
	newProgress := FromDomain(progress)
	err = repo.db.Create(&newProgress).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}
	return id, nil
}
