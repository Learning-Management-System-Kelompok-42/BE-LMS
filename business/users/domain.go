package users

import (
	"time"
)

type Domain struct {
	ID               string
	CompanyID        string
	SpecializationID string
	FullName         string
	Email            string
	Password         string
	PhoneNumber      string
	Address          string
	Role             string
	Avatar           string
	LevelAccess      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DashboardEmployee struct {
	DetailEmployee         DetailEmployeeDashboard
	AmountCourse           int64
	AmountCourseCompleted  int64
	AmountCourseIncomplete int64
	TopCourseProgress      []TopCourseProgress
	TopCourseOften7Days    []LastCourseOpen
}

type TopCourseProgress struct {
	CourseID  string
	Thumbnail string
	Title     string
	Progress  int64
}

type LastCourseOpen struct {
	CourseID string
	Title    string
}

type DetailEmployeeDashboard struct {
	UserID             string
	FullName           string
	SpecializationName string
}

type UserDetailDashboard struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CourseDetailDashboardUser struct {
	ID          string
	Name        string
	Thumbnail   string
	Description string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ToResponseDetailUserDashboard struct {
	User    UserDetailDashboard
	Courses []CourseDetailDashboardUser
}

func NewUser(id, companyID, specializationID, fullName, email, password, phoneNumber, address, role, levelAccess string) Domain {
	return Domain{
		ID:               id,
		CompanyID:        companyID,
		SpecializationID: specializationID,
		FullName:         fullName,
		Email:            email,
		Password:         password,
		PhoneNumber:      phoneNumber,
		Address:          address,
		Role:             role,
		LevelAccess:      levelAccess,
	}
}

func (old *Domain) ModifyUser(fullName, email, phoneNumber, address string, avatar string) Domain {
	return Domain{
		ID:               old.ID,
		CompanyID:        old.CompanyID,
		SpecializationID: old.SpecializationID,
		Role:             old.Role,
		FullName:         fullName,
		Email:            email,
		Password:         old.Password,
		PhoneNumber:      phoneNumber,
		Address:          address,
		LevelAccess:      old.LevelAccess,
		Avatar:           avatar,
	}
}

func (old *Domain) ModifySpecializationName(specializationID string) Domain {
	return Domain{
		ID:               old.ID,
		CompanyID:        old.CompanyID,
		SpecializationID: specializationID,
		Role:             old.Role,
		FullName:         old.FullName,
		Email:            old.Email,
		Password:         old.Password,
		PhoneNumber:      old.PhoneNumber,
		Address:          old.Address,
		LevelAccess:      old.LevelAccess,
		CreatedAt:        old.CreatedAt,
		UpdatedAt:        old.UpdatedAt,
	}
}

func (old *Domain) ModifyPassword(newPassword string) Domain {
	return Domain{
		ID:               old.ID,
		CompanyID:        old.CompanyID,
		SpecializationID: old.SpecializationID,
		Role:             old.Role,
		FullName:         old.FullName,
		Email:            old.Email,
		Password:         newPassword,
		PhoneNumber:      old.PhoneNumber,
		Address:          old.Address,
		LevelAccess:      old.LevelAccess,
		CreatedAt:        old.CreatedAt,
		UpdatedAt:        old.UpdatedAt,
	}
}
