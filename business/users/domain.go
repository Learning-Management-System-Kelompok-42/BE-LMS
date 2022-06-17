package users

import (
	"time"
)

type Domain struct {
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

func (old *Domain) ModifyUser(fullName, phoneNumber, address string) Domain {
	return Domain{
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
