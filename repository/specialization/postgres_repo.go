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
