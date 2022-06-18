package module

import (
	"time"

	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/material"
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
	Materials   []material.Material      `gorm:"foreignKey:ModuleID"`
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
