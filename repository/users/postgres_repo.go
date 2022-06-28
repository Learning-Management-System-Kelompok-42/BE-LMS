package users

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) users.UserRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(user users.Domain) (id string, err error) {
	newUser := FromDomain(user)
	err = repo.db.Create(&newUser).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = user.ID

	return id, nil
}

func (repo *postgreSQLRepository) Update(user users.Domain) (err error) {
	updateUser := FromDomain(user)
	err = repo.db.Save(&updateUser).Error

	if err != nil {
		return exception.ErrInternalServer
	}

	return nil
}

func (repo *postgreSQLRepository) FindByID(id string) (user users.Domain, err error) {
	result := FromDomain(users.Domain{ID: id})

	err = repo.db.First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, exception.ErrDataNotFound
		}
		return user, exception.ErrInternalServer
	}

	user = result.ToDomain()

	return user, nil
}

func (repo *postgreSQLRepository) CheckEmail(email string) error {
	var user User
	err := repo.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrEmailExists
}

func (repo *postgreSQLRepository) FindAllUsers(companyID string) (users []users.Domain, err error) {
	// Next we will add limit and offset
	var user []User
	result := repo.db.Where("company_id = ? AND level_access = ?", companyID, "employee").Find(&user)
	// check if result rows is empty
	if result.RowsAffected == 0 {
		return nil, exception.ErrDataNotFound
	} else if result.Error != nil {
		return nil, exception.ErrInternalServer
	}

	users = ToDomainList(user)

	return users, nil
}

func (repo *postgreSQLRepository) FindDetailUserDashboard(userID string) (user users.Domain, err error) {
	return user, nil
}
