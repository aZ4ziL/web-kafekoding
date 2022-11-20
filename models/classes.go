package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Mentors     []*User        `gorm:"many2many:user_class_mentors" json:"mentors"`
	Members     []*User        `gorm:"many2many:user_class_members" json:"members"`
	Title       string         `gorm:"size:50;unique;index" json:"title"`
	Logo        string         `gorm:"size:255;null" json:"logo"`
	Description string         `gorm:"size:255" json:"description"`
	Content     string         `gorm:"type:longtext" json:"content"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type classModel struct {
	db *gorm.DB
}

func NewClassModel() *classModel {
	return &classModel{GetDB()}
}

func (c classModel) CreateNewClass(class *Class) error {
	return c.db.Create(class).Error
}

func (c classModel) GetAllClass() []Class {
	var classes []Class
	db.Model(&Class{}).Preload("Mentors").Preload("Members").Find(&classes)
	return classes
}

func (c classModel) GetClassByID(id uint) (Class, error) {
	var class Class
	err := db.Model(&Class{}).Where("id = ?", id).Preload("Mentors").Preload("Members").First(&class).Error
	return class, err
}
