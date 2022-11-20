package models

import (
	"time"

	"github.com/aZ4ziL/web-kafekoding/auth"
	"gorm.io/gorm"
)

// Implementing the user model
type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FullName        string    `gorm:"size:30" json:"full_name" form:"full_name" validate:"required"`
	StudentIDNumber uint      `gorm:"unique;index" json:"student_id_number" form:"student_id_number" validate:"required,number"`
	Avatar          string    `gorm:"size:255;null" json:"avatar"`
	PhoneNumber     string    `gorm:"size:15" json:"phone_number" form:"phone_number" validate:"required,number"`
	Campus          string    `gorm:"size:30" json:"campus" form:"campus" validate:"required"`
	Major           string    `gorm:"size:20" json:"major" form:"major" validate:"required"`
	Email           string    `gorm:"size:30;unique;index" json:"email" form:"email" validate:"required,email"`
	Password        string    `gorm:"size:150" json:"-" form:"password"`
	IsAdmin         bool      `gorm:"default:0" json:"is_admin"`
	LastLogin       time.Time `gorm:"null" json:"last_login"`
	DateJoined      time.Time `gorm:"autoCreateTime" json:"date_joined"`
	ClassMentors    []*Class  `gorm:"many2many:user_class_mentors" json:"class_mentors"`
	ClassMembers    []*Class  `gorm:"many2many:user_class_members" json:"class_members"`
}

// userModel
type userModel struct {
	db *gorm.DB
}

func NewUserModel() *userModel {
	return &userModel{GetDB()}
}

// CreateNewUser
// Creating new user if user is exist return error
func (u userModel) CreateNewUser(user *User) error {
	user.Password = auth.EncryptionPassword(user.Password)
	return u.db.Create(user).Error
}

// GetUserByStudentIDNumber
// Get user by given student id number
func (u userModel) GetUserByStudentIDNumber(idNumber uint) (User, error) {
	var user User
	err := db.Model(&User{}).Where("student_id_number = ?", idNumber).Preload("ClassMentors").Preload("ClassMembers").First(&user).Error
	return user, err
}

// GetUserByEmail
// return user by passing the Email
func (u userModel) GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("email = ?", email).Preload("ClassMentors").Preload("ClassMembers").First(&user).Error
	return user, err
}
