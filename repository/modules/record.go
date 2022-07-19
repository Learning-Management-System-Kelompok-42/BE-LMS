package module

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userModules"
	"gorm.io/gorm"
)

type Module struct {
	ID          string `gorm:"primaryKey;size:200"`
	CourseID    string `gorm:"size:200"`
	YoutubeURL  string
	SlideURL    string
	Title       string
	Orders      int
	UserModules []userModules.UserModule `gorm:"foreignKey:ModuleID"`
	Quizs       []quiz.Quiz              `gorm:"foreignKey:ModuleID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (modules *Module) ToDomain() module.Domain {
	return module.Domain{
		ID:         modules.ID,
		CourseID:   modules.CourseID,
		YoutubeURL: modules.YoutubeURL,
		SlideURL:   modules.SlideURL,
		Title:      modules.Title,
		Orders:     modules.Orders,
		CreatedAt:  modules.CreatedAt,
		UpdatedAt:  modules.UpdatedAt,
	}
}

func (modules *Module) ToDomainList() course.ModulesResp {
	return course.ModulesResp{
		ModuleID:   modules.ID,
		CourseID:   modules.CourseID,
		YoutubeURL: modules.YoutubeURL,
		SlideURL:   modules.SlideURL,
		Title:      modules.Title,
		Orders:     modules.Orders,
		Quizzes:    quiz.ToDomainInBatch(modules.Quizs),
	}
}

func FromDomain(module module.Domain) Module {
	return Module{
		ID:         module.ID,
		CourseID:   module.CourseID,
		YoutubeURL: module.YoutubeURL,
		SlideURL:   module.SlideURL,
		Title:      module.Title,
		Orders:     module.Orders,
		CreatedAt:  module.CreatedAt,
		UpdatedAt:  module.UpdatedAt,
		DeletedAt:  gorm.DeletedAt{},
	}
}

func ToDomainBatchList(modules []Module) []module.Domain {
	var domains []module.Domain
	for _, module := range modules {
		domains = append(domains, module.ToDomain())
	}
	return domains
}

func ToDomainInBatch(modules []Module) []course.ModulesResp {
	var domains []course.ModulesResp
	for _, module := range modules {
		domains = append(domains, module.ToDomainList())
	}
	return domains
}
