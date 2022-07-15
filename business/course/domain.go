package course

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
)

type Domain struct {
	ID          string
	CompanyID   string
	Title       string
	Thumbnail   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID               string
	CompanyID        string
	SpecializationID string
	FullName         string
	Email            string
	PhoneNumber      string
	Address          string
	Role             string
	LevelAccess      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ProgressCourse struct {
	ID          string
	Title       string
	Thumbnail   string
	Description string
	Proggress   int64
	Score       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Module struct {
	ID       string
	CourseID string
	Point    int64
	Status   bool
}

// type Modules struct {
// 	ID         string
// 	CourseID   string
// 	YoutubeURL string
// 	SlideURL   string
// 	Title      string
// 	Orders     int
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }

type DetailCourseDashboard struct {
	ID            string
	CourseName    string
	CountModules  int64
	CountEmployee int64
	Users         []users.Domain
	RatingReviews []enrollments.RatingReviews
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DetailCourse struct {
	ID            string
	Title         string
	Thumbnail     string
	Description   string
	CountModule   int64
	CountEmployee int64
	Modules       []module.DetailCourseModules
	RatingReviews []enrollments.RatingReviews
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
		CompanyID:   old.CompanyID,
		Title:       title,
		Thumbnail:   thumbnail,
		Description: description,
	}
}
