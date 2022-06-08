package users

import (
	"time"
)

type User struct {
	ID               string
	CompanyID        string
	Role             string
	SpecializationID string
	FullName         string
	Email            string
	Password         string
	PhoneNumber      string
	Address          string
	LevelAccess      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

func NewUser(id, fullName, email, password, phoneNumber, address, role, companyID, levelAccess string) User {
	return User{
		ID:          id,
		CompanyID:   companyID,
		Role:        role,
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		LevelAccess: levelAccess,
	}
}

func (old *User) ModifyUser(fullName, phoneNumber, address string) User {
	return User{
		ID:          old.ID,
		CompanyID:   old.CompanyID,
		Role:        old.Role,
		FullName:    fullName,
		Email:       old.Email,
		Password:    old.Password,
		PhoneNumber: phoneNumber,
		Address:     address,
		LevelAccess: old.LevelAccess,
	}
}
