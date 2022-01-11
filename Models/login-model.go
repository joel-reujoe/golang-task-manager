package taskmanager



const (
    Admin = iota
    Reviewer  = iota
	Other = iota
)


type User struct {
	Id uint32 `json:"id"`
	FirstName string 
	LastName string
	Email string `json:"email"`
	Password string `json:"password"`	
	UserType int `json:"usertype"`
}
