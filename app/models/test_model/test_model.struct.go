package test_model

// List of all stuct for Category Query

type GetAllUserHandlerStruct struct {
    Id              int64     	`json:"id"`
    Name            string  	`json:"name"`
    Rating          float64 	`json:"rating"`
    Booleandesu     int64     	`json:"booleandesu"`
	Created 	    string		`json:"created"`
    Created_date    string      `json:"created_date"` //date time handle logic in handler/model after struct is returned
}

type InsertCategoryStruct struct {
    Name            string      `json:"name"`
    Created         string      `json:"created"`
    Created_date    string      `json:"created_date"`
}
