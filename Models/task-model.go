package taskmanager

import "gorm.io/gorm"


type AccessDetails struct {
    AccessUuid string
    UserId   uint64
	UserType int
}


type Task struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	ApprovedBy uint32 `json:"approvedby"`
	Approved bool `json:"approved"`
	UserId uint32 `json:"userid"`
}