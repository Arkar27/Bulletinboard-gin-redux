package request

type FindAllUserRequest struct {
}

type LoginRequest struct {
	Email    string `validate:"required,max=200,min=1" binding:"required" json:"email"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}

type UserRequest struct {
	Name            string `validate:"required,max=200,min=1" json:"name"`
	Email           string `validate:"required,max=200,min=1" json:"email"`
	Password        string `validate:"required,max=200,min=1" json:"password"`
	Profile         string `json:"profile"`
	Type            string `default:"1" json:"type"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	Dob             string `json:"dob"`
	Create_user_id  int    `json:"create_user_id"`
	Updated_user_id int    `json:"updated_user_id"`
}

type PostRequest struct {
	Title           string `validate:"required,max=200,min=1" json:"title"`
	Description     string `validate:"required,min=1" json:"description"`
	Status          int    `json:"status"`
	Create_user_id  int    `json:"create_user_id"`
	Updated_user_id int    `json:"updated_user_id"`
}
