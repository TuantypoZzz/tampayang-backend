package entity

type LoginRequest struct {
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

type UserLogin struct {
	User_id			int			`json:"user_id"`	
	User_name		string		`json:"user_name"`	
	User_email		string		`json:"user_email"`	
	User_password	string		`json:"user_password"`
	User_role		string		`json:"user_role"`
}

type ResultToken struct {
	Access_token 		string		`json:"access_token"`
}
