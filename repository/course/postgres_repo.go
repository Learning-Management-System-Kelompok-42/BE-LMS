package course

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) course.CourseRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(course course.Domain) (id string, err error) {
	newCourse := FromDomain(course)
	err = repo.db.Create(&newCourse).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newCourse.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindCourseByIDDashboard(id string) (course course.Domain, err error) {
	var newCourse Course
	// query with nested preload for course module quizzes and multiple choice questions
	// err = repo.db.Where("id = ?", id).Preload("Modules").Find(&newCourse).Error
	err = repo.db.Where("id = ?", id).Find(&newCourse).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return course, exception.ErrNotFound
		}
		return course, exception.ErrNotFound
	}

	course = newCourse.ToDomain()

	return course, nil
}

func (repo *postgreSQLRepository) UpdateCourse(course course.Domain) (id string, err error) {
	newCourse := FromDomain(course)

	err = repo.db.Where("id = ?", course.ID).Save(&newCourse).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newCourse.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindAllCourseDashboard(companyID string) (course []course.Domain, err error) {
	var courses []Course
	result := repo.db.Where("company_id = ?", companyID).Find(&courses)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return course, exception.ErrCourseNotFound
		}
		return course, exception.ErrInternalServer
	}

	course = ToBatchList(courses)

	return course, nil
}

func (repo *postgreSQLRepository) FindAllCourseByUserID(userID string) (course []course.Domain, err error) {
	/**
	Next change structure DB on table user_courses and courses
	Add column rating on table courses
	calculate automatically rating when user give rating on course
	*/
	// subQuery := repo.db.Table("user_courses").Select("avg(rating)").Where("user_id = ? ", userID)

	result := repo.db.Table("courses").
		Select("courses.id, courses.title, courses.thumbnail, courses.description, courses.created_at, courses.updated_at").
		Joins("INNER JOIN user_courses ON courses.id = user_courses.course_id").
		Find(&course)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return course, nil
}

func (repo *postgreSQLRepository) FindAllCourseBySpecializationID(specializationID string) (courses []course.Domain, err error) {
	var course []Course

	result := repo.db.Table("specialization_courses").
		Select("courses.id, courses.title, courses.thumbnail, courses.description, courses.created_at, courses.updated_at").
		Joins("INNER JOIN courses ON specialization_courses.course_id = courses.id").
		Where("specialization_courses.specialization_id = ?", specializationID).
		Find(&course)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrCourseNotFound
		}
		return nil, exception.ErrInternalServer
	}

	courses = ToBatchList(course)

	return courses, nil
}

func (repo *postgreSQLRepository) CountModulesByCourseID(courseID string) (count int64, err error) {
	var countModule int64
	result := repo.db.Table("modules").Where("course_id = ?", courseID).Count(&countModule)

	if result.Error != nil {
		return count, exception.ErrInternalServer
	}

	return countModule, nil
}

func (repo *postgreSQLRepository) CountEmployeeByCourseID(courseID string) (count int64, err error) {
	var countEmployee int64
	result := repo.db.Table("user_courses").Where("course_id = ?", courseID).Count(&countEmployee)

	if result.Error != nil {
		return count, exception.ErrInternalServer
	}

	return countEmployee, nil
}

func (repo *postgreSQLRepository) FindCourseByID(id string) (course course.Domain, err error) {
	var courses Course
	err = repo.db.Where("id = ?", id).First(&courses).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return course, exception.ErrCourseNotFound
		}
		return course, exception.ErrInternalServer
	}

	course = courses.ToDomain()

	return course, nil
}

// func (repo *postgreSQLRepository) FindAllModuleByCourseID(courseID string) (modules []course.Module, err error) {
// 	result := repo.db.Table("modules").
// 		Select("modules.id, modules.course_id").
// 		Joins("INNER JOIN courses ON modules.course_id = courses.id").
// 		Where("modules.course_id = ?", courseID).
// 		Find(&modules)

// 	if result.Error != nil {
// 		if result.RowsAffected == 0 {
// 			return nil, exception.ErrModuleNotFound
// 		}
// 		return nil, exception.ErrInternalServer
// 	}

// 	return modules, nil
// }

func (repo *postgreSQLRepository) CountModulesCompletedByUserID(courseID, userID string) (count int64, err error) {
	// var id struct {
	// 	ID string
	// }
	result := repo.db.Table("user_modules").
		// Select("user_modules.id").
		Joins("INNER JOIN modules ON user_modules.module_id = modules.id").
		Where("user_modules.user_id = ? AND modules.course_id = ? AND status = true", userID, courseID).
		Count(&count)

	// count = int64(len(id.ID))

	if result.Error != nil {
		return count, exception.ErrInternalServer
	}

	return count, nil
}

func (repo *postgreSQLRepository) FindAllPointModuleByModuleID(courseID, userID string) (pointModules []course.Module, err error) {
	result := repo.db.Table("user_modules").
		Select("user_modules.id as id, modules.id as course_id, user_modules.point as point, user_modules.status as status").
		Joins("INNER JOIN modules ON user_modules.module_id = modules.id").
		Where("user_modules.user_id = ? AND modules.course_id = ?", userID, courseID).
		Find(&pointModules)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, exception.ErrModuleNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return pointModules, nil
}

func (repo *postgreSQLRepository) FindAllModuleByCourseID(courseID string) (modules []module.DetailCourseModules, err error) {
	result := repo.db.Table("modules").
		Select("modules.id, modules.course_id, modules.title, modules.youtube_url, modules.slide_url, modules.orders, user_modules.status as status, modules.created_at, modules.updated_at").
		Joins("INNER JOIN user_modules ON modules.id = user_modules.module_id").
		Where("user_modules.course_id = ?", courseID).
		Order("modules.orders ASC").
		Find(&modules)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return modules, exception.ErrModuleNotFound
		}
		return modules, exception.ErrInternalServer
	}

	return modules, nil
}
