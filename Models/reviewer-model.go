package taskmanager

import "gorm.io/gorm"



type ReviewersTask struct {
	gorm.Model
	ReviewerId uint 
	UserId uint
}