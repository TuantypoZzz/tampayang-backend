package user_handler

type CreateUserHandlerStruct struct {
	Name         string `json:"name`
	Age          int    `json:"age"`
	Created_date string `json:"created_date"`
}

type lastIdResponse struct {
	Message string `json:"message"`
	LastId  int64  `json:"lastId"`
}
