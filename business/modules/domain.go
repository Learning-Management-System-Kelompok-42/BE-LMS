package module

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
)

type Domain struct {
	ID         string
	CourseID   string
	YoutubeURL string
	SlideURL   string
	Title      string
	Orders     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DetailCourseModules struct {
	ID        string
	Title     string
	Orders    int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CourseModules struct {
	ID         string
	CourseID   string
	YoutubeURL string
	SlideURL   string
	Quiz       []quiz.Domain
}

func NewModule(id, courseID, youtubeURL, slideURL, title string, orders int) Domain {
	return Domain{
		ID:         id,
		CourseID:   courseID,
		YoutubeURL: youtubeURL,
		SlideURL:   slideURL,
		Title:      title,
		Orders:     orders,
	}
}

func (old *Domain) ModifyModule(title, youtubeURL, slideURL string, orders int) Domain {
	return Domain{
		ID:         old.ID,
		CourseID:   old.CourseID,
		YoutubeURL: youtubeURL,
		SlideURL:   slideURL,
		Title:      title,
		Orders:     orders,
		CreatedAt:  old.CreatedAt,
	}
}
