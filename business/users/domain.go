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

func NewUser(id, companyID, specializationID, role, fullName, email, password, phoneNumber, address, levelAccess string) Domain {
	return Domain{
		ID:               id,
		CompanyID:        companyID,
		SpecializationID: specializationID,
		Role:             role,
		FullName:         fullName,
		Email:            email,
		Password:         password,
		PhoneNumber:      phoneNumber,
		Address:          address,
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
