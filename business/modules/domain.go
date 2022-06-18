package module

import "time"

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
	}
}
