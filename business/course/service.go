package course

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	specModule "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	specQuiz "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CourseRepository interface {
	// Insert insert a new course
	Insert(course Domain) (id string, err error)

	// FindByID find a course by id
	FindCourseByIDDashboard(id string) (course Domain, err error)

	// Update update a course
	UpdateCourse(course Domain) (id string, err error)

	// FindCourseByID find a course by id
	FindCourseByID(id string) (course Domain, err error)

	// FindAllCourseDashboard get all course on dashboard admin
	FindAllCourseDashboard(companyID string) (course []Domain, err error)

	// FindAllCourseByUserID get all course by user id
	FindAllCourseByUserID(userID string) (course []Domain, err error)

	// FindAllCourseBySpecializationID get all course by specialization id
	FindAllCourseBySpecializationID(specializationID string) (courses []Domain, err error)

	// FindAllPointModuleByModuleID get all point module by module id
	FindAllPointModuleByModuleID(courseID, userID string) (pointModules []Module, err error)

	// CountModulesByCourseID get count modules by course id
	CountModulesByCourseID(courseID string) (count int64, err error)

	// CountModulesCompletedByUserID get count modules completed by course id
	CountModulesCompletedByUserID(courseID, userID string) (count int64, err error)

	// CountEmployeeByCourseID get count employee by course id
	CountEmployeeByCourseID(courseID string) (count int64, err error)

	// FindAllModuleByCourseID get all module by course id
	FindAllModuleByCourseID(courseID string) (modules []module.Domain, err error)
}

type CourseService interface {
	// Insert insert a new course
	CreateCourse(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error)

	// GetByID get a course by id
	GetDetailCourseByIDDashboard(courseID string) (courses DetailCourseDashboard, err error)

	// Update update a course``
	UpdateCourse(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error)

	// GetAllCourseDashboard get all course on dashboard admin
	GetAllCourseDashboard(companyID string) (course []Domain, err error)

	// GetDetailCourseByID get a course by id on dashboard employee
	GetDetailCourseByID(courseID string) (courses DetailCourse, err error)

	// GetAllCourse get all course by user id
	GetAllCourse(specializationID, userID string) (resp []ProgressCourse, err error)
}

type courseService struct {
	courseRepo     CourseRepository
	userRepo       users.UserRepository
	enrollmentRepo enrollments.EnrollmentRepository
	serviceModule  module.ModuleService
	serviceQuiz    quiz.QuizService
	validate       *validator.Validate
}

func NewCourseService(
	courseRepo CourseRepository,
	userRepo users.UserRepository,
	enrollmentRepo enrollments.EnrollmentRepository,
	serviceModule module.ModuleService,
	serviceQuiz quiz.QuizService,
) CourseService {
	return &courseService{
		courseRepo:     courseRepo,
		userRepo:       userRepo,
		enrollmentRepo: enrollmentRepo,
		serviceModule:  serviceModule,
		serviceQuiz:    serviceQuiz,
		validate:       validator.New(),
	}
}

func (s *courseService) CreateCourse(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error) {
	err = s.validate.Struct(&upsertCourseSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	courseID := uuid.New().String()
	newCourse := NewCourse(
		courseID,
		upsertCourseSpec.CompanyID,
		upsertCourseSpec.Title,
		upsertCourseSpec.Thumbnail,
		upsertCourseSpec.Description,
	)
	course, err := s.courseRepo.Insert(newCourse)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	for _, module := range upsertCourseSpec.Modules {
		module.CourseID = course
		newModules := specModule.UpsertModuleSpec{
			CourseID:   module.CourseID,
			Title:      module.Title,
			YoutubeURL: module.YoutubeURL,
			SlideURL:   module.SlideURL,
			Orders:     module.Orders,
		}

		modulesID, err := s.serviceModule.Create(newModules)
		if err != nil {
			if err == exception.ErrInvalidRequest {
				return "", exception.ErrInvalidRequest
			}

			return "", exception.ErrInternalServer
		}

		for _, quiz := range module.Quizzes {
			newQuiz := specQuiz.UpsertQuizSpec{
				ModuleID:       modulesID,
				Question:       quiz.Question,
				Answer:         quiz.Answer,
				MultipleChoice: quiz.MultipleChoice,
			}

			_, err := s.serviceQuiz.Create(newQuiz)
			if err != nil {
				if err == exception.ErrInvalidRequest {
					return "", exception.ErrInvalidRequest
				}
				return "", exception.ErrInternalServer
			}
		}

	}

	return course, nil
}

func (s *courseService) GetDetailCourseByIDDashboard(courseID string) (courses DetailCourseDashboard, err error) {
	course, err := s.courseRepo.FindCourseByIDDashboard(courseID)
	if err != nil {
		return courses, exception.ErrInternalServer
	}

	countModules, err := s.courseRepo.CountModulesByCourseID(courseID)
	if err != nil {
		return courses, exception.ErrInternalServer
	}

	countEmploye, err := s.courseRepo.CountEmployeeByCourseID(courseID)
	if err != nil {
		return courses, exception.ErrInternalServer
	}

	enrollments, err := s.enrollmentRepo.FindAllEnrollmentsByCourseID(courseID)
	if err != nil {
		return courses, exception.ErrInternalServer
	}

	user, err := s.userRepo.FindAllUserByCourseID(courseID)
	if err != nil {
		return courses, exception.ErrInternalServer
	}

	courses = DetailCourseDashboard{
		ID:            course.ID,
		CourseName:    course.Title,
		CountModules:  countModules,
		CountEmployee: countEmploye,
		Users:         user,
		RatingReviews: enrollments,
	}

	return courses, nil
}

func (s *courseService) UpdateCourse(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error) {
	err = s.validate.Struct(&upsertCourseSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldCourse, err := s.courseRepo.FindCourseByID(upsertCourseSpec.ID)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	upsertCourse := oldCourse.ModifyCourse(
		upsertCourseSpec.Title,
		upsertCourseSpec.Thumbnail,
		upsertCourseSpec.Description,
	)

	newCourse, err := s.courseRepo.UpdateCourse(upsertCourse)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	for _, module := range upsertCourseSpec.Modules {
		module.CourseID = newCourse
		newModules := specModule.UpsertModuleSpec{
			ID:         module.ModuleID,
			CourseID:   module.CourseID,
			Title:      module.Title,
			YoutubeURL: module.YoutubeURL,
			SlideURL:   module.SlideURL,
			Orders:     module.Orders,
		}

		modulesID, err := s.serviceModule.Update(newModules)
		if err != nil {
			if err == exception.ErrInvalidRequest {
				return "", exception.ErrInvalidRequest
			}

			return "", exception.ErrInternalServer
		}

		for _, quiz := range module.Quizzes {
			newQuiz := specQuiz.UpsertQuizSpec{
				ID:             quiz.QuizID,
				ModuleID:       modulesID,
				Question:       quiz.Question,
				Answer:         quiz.Answer,
				MultipleChoice: quiz.MultipleChoice,
			}

			_, err := s.serviceQuiz.Update(newQuiz)
			if err != nil {
				if err == exception.ErrInvalidRequest {
					return "", exception.ErrInvalidRequest
				}
			}
		}
	}

	return newCourse, nil
}

func (s *courseService) GetAllCourseDashboard(companyID string) (course []Domain, err error) {
	course, err = s.courseRepo.FindAllCourseDashboard(companyID)
	if err != nil {
		if err == exception.ErrNotFound {
			return course, exception.ErrNotFound
		}
		return course, exception.ErrInternalServer
	}

	return course, nil
}

func (s *courseService) GetAllCourse(specializationID, userID string) (resp []ProgressCourse, err error) {
	courses, err := s.courseRepo.FindAllCourseBySpecializationID(specializationID)
	if err != nil {
		if err == exception.ErrNotFound {
			return nil, exception.ErrNotFound
		}
		return nil, exception.ErrInternalServer
	}

	// insert courses to progress resp
	for _, v := range courses {
		// Calculate percentage progress course by user
		totalModule, _ := s.courseRepo.CountModulesByCourseID(v.ID)
		moduleCompleted, _ := s.courseRepo.CountModulesCompletedByUserID(v.ID, userID)
		percentage := float64(moduleCompleted) / float64(totalModule) * 100
		fmt.Println("percentage =", percentage)
		// percentage total module completed from course
		var percentageModule int64
		percentageModule = 0
		if totalModule != 0 {
			percentageModule = (moduleCompleted * 100) / totalModule
		}

		// message := fmt.Sprintf("totalModule = %d, percentageModule = %d, modulesCompleted = %d", totalModule, percentageModule, moduleCompleted)
		// fmt.Println(message)

		// Calculate score from module
		point, err := s.courseRepo.FindAllPointModuleByModuleID(v.ID, userID)
		if err != nil {
			return nil, exception.ErrInternalServer
		}

		fmt.Println("point = ", point)

		// Calculate total point from module
		var totalPoint int64
		totalPoint = 0
		for _, v := range point {
			totalPoint += v.Point
		}

		if totalPoint != 0 {
			totalPoint = totalPoint / totalModule
		}

		// message1 := fmt.Sprintf("totalPoint = %d untuk course = %s dengan jumlah module = %d", totalPoint, v.ID, totalModule)
		// fmt.Println(message1)

		progress := ProgressCourse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Thumbnail:   v.Thumbnail,
			Proggress:   percentageModule,
			Score:       totalPoint,
		}
		resp = append(resp, progress)
	}

	return resp, nil
}

func (s *courseService) GetDetailCourseByID(courseID string) (courses DetailCourse, err error) {
	course, err := s.courseRepo.FindCourseByID(courseID)
	if err != nil {
		if err == exception.ErrCourseNotFound {
			return courses, exception.ErrCourseNotFound
		}
		return courses, exception.ErrInternalServer
	}

	modules, err := s.courseRepo.FindAllModuleByCourseID(courseID)
	if err != nil {
		if err == exception.ErrModuleNotFound {
			return courses, exception.ErrModuleNotFound
		}
		return courses, exception.ErrInternalServer
	}

	employees, err := s.enrollmentRepo.FindAllEnrollmentsByCourseID(courseID)
	if err != nil {
		if err == exception.ErrEnrollmentNotFound {
			return courses, exception.ErrEnrollmentNotFound
		}
		return courses, exception.ErrInternalServer
	}

	amountModule := int64(len(modules))
	amountEmployee := int64(len(employees))

	courses = DetailCourse{
		ID:            course.ID,
		Title:         course.Title,
		Thumbnail:     course.Thumbnail,
		Description:   course.Description,
		CountModule:   amountModule,
		CountEmployee: amountEmployee,
		Modules:       modules,
		RatingReviews: employees,
	}

	return courses, nil
}
