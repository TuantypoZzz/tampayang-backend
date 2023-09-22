package example_model

type InsertExampleStruct struct {
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}

type GetExampleByIdStruct struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}

type GetAllExampleStruct struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}

type UpdateExampleStruct struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}
