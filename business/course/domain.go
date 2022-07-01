package course

import "time"

type Domain struct {
	ID          string
	CompanyID   string
	Title       string
	Thumbnail   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCourse(id, companyID, title, thumbnail, description string) Domain {
	return Domain{
		ID:          id,
		CompanyID:   companyID,
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
