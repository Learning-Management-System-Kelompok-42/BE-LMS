package response

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
)

type GetDetailCourseResp struct {
	ID            string                `json:"id"`
	Title         string                `json:"title"`
	Thumbnail     string                `json:"thumbnail"`
	Description   string                `json:"description"`
	CountModules  int64                 `json:"count_modules"`
	CountEmployee int64                 `json:"count_employee"`
	Modules       []moduleResp          `json:"modules"`
	RatingReviews []ratingReviewsRespon `json:"rating_reviews"`
}

type moduleResp struct {
	ID        string
	Title     string
	Orders    int
	CreatedAt string
	UpdatedAt string
}

type ratingReviewsRespon struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Rating  float32 `json:"rating"`
	Reviews string  `json:"reviews"`
}

func NewGetDetailCourseResp(domain course.DetailCourse) GetDetailCourseResp {
	// looping user and rating reviews
	var ratingReviews []ratingReviewsRespon
	var modules []moduleResp

	for _, module := range domain.Modules {
		modules = append(modules, moduleResp{
			ID:     module.ID,
			Title:  module.Title,
			Orders: module.Orders,
			// CreatedAt: module.CreatedAt.Local().Format("2006-01-02 15:04:05"),
			CreatedAt: module.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: module.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	for _, ratingReview := range domain.RatingReviews {
		ratingReviews = append(ratingReviews, ratingReviewsRespon{
			ID:      ratingReview.ID,
			Name:    ratingReview.Name,
			Email:   ratingReview.Email,
			Rating:  ratingReview.Rating,
			Reviews: ratingReview.Reviews,
		})
	}

	return GetDetailCourseResp{
		ID:            domain.ID,
		Title:         domain.Title,
		Thumbnail:     domain.Thumbnail,
		Description:   domain.Description,
		CountModules:  domain.CountModule,
		CountEmployee: domain.CountEmployee,
		Modules:       modules,
		RatingReviews: ratingReviews,
	}

}
