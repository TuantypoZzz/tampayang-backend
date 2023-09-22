package user_model

type InsertNewUserStruct struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Created_date string `json:"created_date"`
}

type GetAllUserHandlerStruct struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Age          string `json:"age"`
	Created_date string `json:"created_date"`
}
