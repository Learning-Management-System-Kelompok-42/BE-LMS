package auth

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) auth.AuthRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Login(email string) (user users.User, err error) {
	err = repo.db.Select("id", "email", "password", "level_access", "company_id").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, exception.ErrNotFound
		}
		return user, exception.ErrInternalServer
	}

	return user, nil
}
