package record

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               string `gorm:"primaryKey;size:200;not null"`
	CompanyID        string `gorm:"size:200"`
	SpecializationID string `gorm:"size:200"`
	FullName         string
	Email            string `gorm:"size:250;uniqueIndex"`
	Password         string
	PhoneNumber      string
	Address          string
	Role             string
	LevelAccess      string
	UserCourses      []UserCourse  `gorm:"foreignKey:CourseID"`
	UserModules      []UserModule  `gorm:"foreignKey:CourseID"`
	Certificates     []Certificate `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Company struct {
	ID             string `gorm:"primaryKey,size:200"`
	Name           string
	Address        string
	Web            string
	Email          string `gorm:"size:250;uniqueIndex"`
	Sector         string
	Logo           string
	Users          []User          `gorm:"foreignKey:CompanyID"`
	RequestCourses []RequestCourse `gorm:"foreignKey:CompanyID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type RequestCourse struct {
	ID         string `gorm:"primaryKey,size:200"`
	CompanyID  string `gorm:"size:200"`
	NameCourse string
	Reason     string
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Specialization struct {
	ID                    string `gorm:"primaryKey;size:200"`
	Name                  string
	Invitation            string
	Users                 []User                 `gorm:"foreignKey:SpecializationID"`
	SpecializationCourses []SpecializationCourse `gorm:"foreignKey:SpecializationID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

type Faq struct {
	ID        string `gorm:"primaryKey;size:200"`
	Question  string
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserCourse struct {
	ID        string  `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID  string  `gorm:"size:200"`
	UserID    string  `gorm:"size:200"`
	Rating    float32 `gorm:"type:numeric(2,2)"`
	Reviews   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Certificate struct {
	ID        string `gorm:"primaryKey;size:200"`
	CourseID  string `gorm:"size:200"`
	UserID    string `gorm:"size:200"`
	Signature string
	Expired   time.Time
	CreatedAt time.Time //will be aliasing as Publish
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserModule struct {
	ID        string `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID  string `gorm:"size:200"`
	ModuleID  string `gorm:"size:200"`
	Point     int32
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Module struct {
	ID          string `gorm:"primaryKey;size:200"`
	CourseID    string `gorm:"size:200"`
	Title       string
	Orders      int32
	UserModules []UserModule `gorm:"foreignKey:ModuleID"`
	Materials   []Material   `gorm:"foreignKey:ModuleID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type SpecializationCourse struct {
	ID               string `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID         string `gorm:"size:200"`
	SpecializationID string `gorm:"size:200"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Material struct {
	ID        string `gorm:"primaryKey;size:200"`
	ModuleID  string `gorm:"size:200"`
	Title     string
	Url       string
	Type      string
	Orders    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Quiz struct {
	ID        string `gorm:"primaryKey;size:200"`
	ModuleID  string `gorm:"size:200"`
	Title     string
	Question  string
	Options   []Option `gorm:"foreignKey:QuizID"`
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Option struct {
	ID        string `gorm:"primaryKey;size:200"`
	QuizID    string `gorm:"size:200"`
	Option    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
