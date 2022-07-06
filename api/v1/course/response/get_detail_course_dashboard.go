package response

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
)

type GetDetailCourseDashbordResp struct {
	ID            string              `json:"id"`
	Title         string              `json:"title"`
	CountModules  int64               `json:"count_modules"`
	CountEmployee int64               `json:"count_employee"`
	User          []userResp          `json:"users"`
	RatingReviews []ratingReviewsResp `json:"rating_reviews"`
}

type userResp struct {
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

type ratingReviewsResp struct {
	ID      string
	Name    string
	Email   string
	Rating  float32
	Reviews string
}

func NewGetDetailCourseDashbordResp(domain course.DetailCourseDashboard) GetDetailCourseDashbordResp {
	// looping user and rating reviews
	var users []userResp
	var ratingReviews []ratingReviewsResp

	for _, user := range domain.Users {
		users = append(users, userResp{
			ID:               user.ID,
			CompanyID:        user.CompanyID,
			SpecializationID: user.SpecializationID,
			FullName:         user.FullName,
			Email:            user.Email,
			PhoneNumber:      user.PhoneNumber,
			Address:          user.Address,
			Role:             user.Role,
			LevelAccess:      user.LevelAccess,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
		})
	}

	for _, ratingReview := range domain.RatingReviews {
		ratingReviews = append(ratingReviews, ratingReviewsResp{
			ID:      ratingReview.ID,
			Name:    ratingReview.Name,
			Email:   ratingReview.Email,
			Rating:  ratingReview.Rating,
			Reviews: ratingReview.Reviews,
		})
	}

	return GetDetailCourseDashbordResp{
		ID:            domain.ID,
		Title:         domain.CourseName,
		CountModules:  domain.CountModules,
		CountEmployee: domain.CountEmployee,
		User:          users,
		RatingReviews: ratingReviews,
	}

}
