package models

import "gorm.io/gorm"

//群消息
type GroupBasic struct {
	gorm.Model
	Name  string
	Owner uint
	Icon  string
	Type  int
	Desc  string
}

func (table *GroupBasic) TableName() string {
	return "groupbasic"
}
