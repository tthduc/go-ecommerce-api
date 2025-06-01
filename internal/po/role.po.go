package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleName string `gorm:"column:name;not null" json:"name"`
	RoleNode string `gorm:"column:role_node" json:"role_node"`
}

func (Role) TableName() string {
	return "go_db_role"
}
