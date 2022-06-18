package course

import "time"

type Domain struct {
	ID          string
	Title       string
	Thumbnail   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCourse(id, title, thumbnail, description string) Domain {
	return Domain{
		ID:          id,
		Title:       title,
		Thumbnail:   thumbnail,
		Description: description,
	}
}

func (old *Domain) ModifyCourse(title, thumbnail, description string) Domain {
	return Domain{
		ID:          old.ID,
		Title:       title,
		Thumbnail:   thumbnail,
		Description: description,
	}
}
