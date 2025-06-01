package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid" json:"id"`
	Name     string    `gorm:"column:name;not null" json:"name"`
	IsActive bool      `gorm:"column:is_active;not null" json:"is_active"`
	Roles    []Role    `gorm:"many2many:go_user_roles" json:"roles"`
}

func (User) TableName() string {
	return "go_db_user"
}
