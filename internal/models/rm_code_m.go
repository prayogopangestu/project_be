package models

import "gorm.io/gorm"

type RmCodeM struct {
	gorm.Model
	RmCodeKode      string `gorm:"column:rm_code_kode;type:varchar(255);NOT NULL;unique" json:"rm_code_kode"`
	RmCodeNama      string `gorm:"column:rm_code_nama;type:varchar(255);NOT NULL" json:"rm_code_nama"`
	RmCodeDeskripsi string `gorm:"column:rm_code_deskripsi;NOT NULL" json:"rm_code_deskripsi"`
	CreateUserID    uint   `gorm:"column:create_user_id;NOT NULL" json:"create_user_id"`
	CreateUser      *User  `gorm:"foreignKey:CreateUserID;constraint:onDelete:RESTRICT,onUpdate:RESTRICT" json:"create_user"`
	UpdateUserID    *uint  `gorm:"column:update_user_id" json:"update_user_id"`
	UpdateUser      *User  `gorm:"foreignKey:UpdateUserID;constraint:onDelete:RESTRICT,onUpdate:RESTRICT" json:"update_user"`
	DeleteUserID    *uint  `gorm:"column:delete_user_id" json:"delete_user_id"`
	DeleteUser      *User  `gorm:"foreignKey:DeleteUserID;constraint:onDelete:RESTRICT,onUpdate:RESTRICT" json:"delete_user"`
}

func (m *RmCodeM) TableName() string {
	return "rm_code_m"
}
