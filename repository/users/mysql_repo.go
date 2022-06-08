package users

import (
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Insert(user domain.User) (id string, err error) {
	err = repo.db.Omit("specialization_id").Create(&user).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = user.ID

	return id, nil
}

func (repo *userRepository) Update(user domain.User) (err error) {
	err = repo.db.Save(&user).Error

	if err != nil {
		return exception.ErrInternalServer
	}

	return nil
}

func (repo *userRepository) GetByID(id string) (user *domain.User, err error) {
	err = repo.db.Where("Id =? ", id).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrDataNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return user, nil
}

func (repo *userRepository) Login(email, password string) (user domain.User, err error) {
	err = repo.db.Where("email = ? AND password = ?", email, password).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, exception.ErrDataNotFound
		}
		return user, exception.ErrInternalServer
	}

	return user, nil
}

func (repo *userRepository) CheckEmail(email string) error {
	var user domain.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrEmailExists
}
