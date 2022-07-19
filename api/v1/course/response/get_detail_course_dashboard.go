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
	ID               string    `json:"id"`
	CompanyID        string    `json:"company_id"`
	SpecializationID string    `json:"specialization_id"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ratingReviewsResp struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Rating  float32 `json:"rating"`
	Reviews string  `json:"reviews"`
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
