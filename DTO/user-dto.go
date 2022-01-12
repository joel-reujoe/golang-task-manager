package taskmanager

type CreateUserDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsReviewer bool `json:"isReviewer"`
}


type GetUserByIdDto struct {
	UserId uint `uri:"userId"`
}

type GetUserByIdResponseDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsReviewer bool `json:"isReviewer"`
}

type GetTaskByUserIdDto struct {
	UserId uint `uri:"userId"`
} 

type AssignTaskToReviewerDto struct {
	UserId uint `json:"userId"`
	ReviewerId []uint `json:"reviewerId"`
}

type GetReviewerTaskByUserId struct{
	UserId uint `uri:"userId"`
}