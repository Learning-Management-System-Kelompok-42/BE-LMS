package course

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course/spec"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	specModule "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	specQuiz "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CourseRepository interface {
	// Insert insert a new course
	Insert(course Domain) (id string, err error)

	// FindByID find a course by id
	FindByID(id string) (course Domain, err error)

	// Update update a course
	Update(course Domain) (id string, err error)

	// FindAllCourseDashboard get all course on dashboard admin
	FindAllCourseDashboard(companyID string) (course []Domain, err error)
}

type CourseService interface {
	// Insert insert a new course
	Create(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error)

	// GetByID get a course by id
	GetByID(id string) (course Domain, err error)

	// Update update a course
	Update(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error)

	// GetAllCourseDashboard get all course on dashboard admin
	GetAllCourseDashboard(companyID string) (course []Domain, err error)
}

type courseService struct {
	repo          CourseRepository
	serviceModule module.ModuleService
	serviceQuiz   quiz.QuizService
	validate      *validator.Validate
}

func NewCourseService(repo CourseRepository, serviceModule module.ModuleService, serviceQuiz quiz.QuizService) CourseService {
	return &courseService{
		repo:          repo,
		serviceModule: serviceModule,
		serviceQuiz:   serviceQuiz,
		validate:      validator.New(),
	}
}

func (s *courseService) Create(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error) {
	err = s.validate.Struct(&upsertCourseSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	courseID := uuid.New().String()
	newCourse := NewCourse(
		courseID,
		upsertCourseSpec.Title,
		upsertCourseSpec.Thumbnail,
		upsertCourseSpec.Description,
	)
	course, err := s.repo.Insert(newCourse)
	if err != nil {
		return "", exception.ErrInternalServer
	}
	// fmt.Println("course id dari insert = ", course)

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

		// fmt.Println("modules id dari insert = ", modulesID)

		for _, quiz := range module.Quizzes {
			newQuiz := specQuiz.UpsertQuizSpec{
				ModuleID:       modulesID,
				Question:       quiz.Question,
				Answer:         quiz.Answer,
				MultipleChoice: quiz.MultipleChoice,
			}

			quizID, err := s.serviceQuiz.Create(newQuiz)
			if err != nil {
				if err == exception.ErrInvalidRequest {
					return "", exception.ErrInvalidRequest
				}

				return "", exception.ErrInternalServer
			}

			fmt.Println("quiz id dari insert = ", quizID)
		}
	}

	return course, nil
}

func (s *courseService) GetByID(id string) (course Domain, err error) {
	course, err = s.repo.FindByID(id)
	if err != nil {
		if err == exception.ErrNotFound {
			return course, exception.ErrNotFound
		}
		return course, exception.ErrInternalServer
	}

	return course, nil
}

func (s *courseService) Update(upsertCourseSpec spec.UpsertCourseSpec) (id string, err error) {
	return id, nil
}

func (s *courseService) GetAllCourseDashboard(companyID string) (course []Domain, err error) {
	course, err = s.repo.FindAllCourseDashboard(companyID)
	if err != nil {
		if err == exception.ErrNotFound {
			return course, exception.ErrNotFound
		}
		return course, exception.ErrInternalServer
	}

	return course, nil
}
