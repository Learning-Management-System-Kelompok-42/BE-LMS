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

func (repo *postgreSQLRepository) UpdateSpecializationName(userUpdate users.Domain) (id string, err error) {
	updateUser := FromDomain(userUpdate)
	err = repo.db.Save(&updateUser).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = userUpdate.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindByID(id string) (user users.Domain, err error) {
	// result := FromDomain(users.Domain{ID: id})
	var users User

	err = repo.db.Where("id = ?", id).First(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, exception.ErrEmployeeNotFound
		}
		return user, exception.ErrInternalServer
	}

	user = users.ToDomain()

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
		return nil, exception.ErrEmployeeNotFound
	} else if result.Error != nil {
		return nil, exception.ErrInternalServer
	}

	users = ToDomainList(user)

	return users, nil
}

func (repo *postgreSQLRepository) FindDetailUserDashboard(userID string) (user users.UserDetailDashboard, err error) {
	userQuery := repo.db.Table("users").
		Select("users.id, users.full_name AS name, users.email, users.phone_number AS phone_number, users.address, specializations.name AS role, users.created_at, users.updated_at").
		Joins("INNER JOIN specializations ON users.specialization_id = specializations.id").
		Where("users.id = ?", userID).
		Scan(&user)

	if userQuery.Error != nil {
		if userQuery.RowsAffected == 0 {
			return user, exception.ErrNotFound
		}
		return user, exception.ErrInternalServer
	}

	return user, nil
}

func (repo *postgreSQLRepository) FindDetailCourseDashboardUsers(userID string) (courses []users.CourseDetailDashboardUser, err error) {
	result := repo.db.Table("courses").
		Select("courses.id AS id, courses.title as name, courses.thumbnail, courses.description, AVG(enrollments.rating) as rating, courses.created_at, courses.updated_at").
		Joins("INNER JOIN enrollments ON courses.id = enrollments.course_id").
		Where("enrollments.user_id = ?", userID).
		Group("courses.id").
		Find(&courses)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrEmployeeNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return courses, nil
}

func (repo *postgreSQLRepository) FindAllUserBySpecializationID(specializationID string) (users []users.Domain, err error) {
	var user []User
	result := repo.db.Table("users").
		Select("users.id, users.company_id, users.specialization_id, users.role, users.full_name AS full_name, users.email, users.phone_number AS phone_number, users.address, users.created_at, users.updated_at").
		Joins("INNER JOIN specializations ON users.specialization_id = specializations.id").
		Where("specializations.id = ?", specializationID).
		Find(&user)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrEmployeeNotFound
		}
		return nil, exception.ErrInternalServer
	}

	users = ToDomainList(user)

	return users, nil
}

func (repo *postgreSQLRepository) FindAllUserByCourseID(courseID string) (users []users.Domain, err error) {
	var user []User
	result := repo.db.Table("users").
		Select("users.id, users.company_id, users.specialization_id, users.role, users.full_name AS full_name, users.email, users.phone_number AS phone_number, users.address, users.created_at, users.updated_at").
		Joins("INNER JOIN enrollments ON users.id = enrollments.user_id").
		Where("enrollments.course_id = ?", courseID).
		Find(&user)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrEmployeeNotFound
		}
		return nil, exception.ErrInternalServer
	}

	users = ToDomainList(user)

	return users, nil
}

func (repo *postgreSQLRepository) UpdateProfile(userUpdate users.Domain) (id string, err error) {
	oldUser := FromDomain(userUpdate)

	err = repo.db.Save(&oldUser).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = userUpdate.ID

	return id, nil
}

func (repo *postgreSQLRepository) UpdatePassword(userUpdate users.Domain) (id string, err error) {
	oldUser := FromDomain(userUpdate)

	err = repo.db.Save(&oldUser).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = userUpdate.ID

	return id, nil
}
