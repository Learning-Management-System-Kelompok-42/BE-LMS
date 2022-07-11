package requestFeat

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) requestFeat.RequestFeatRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) InsertRequestFeat(domain requestFeat.Domain) (id string, err error) {
	newRequest := FromDomain(domain)

	fmt.Println("before insert = ", newRequest)

	err = repo.db.Create(&newRequest).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newRequest.ID

	fmt.Println("after insert = ", newRequest)

	return id, nil
}
