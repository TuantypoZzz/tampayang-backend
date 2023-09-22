package entity

type Example struct {
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}

type ExampleWithId struct {
	Id           int	 `json:"id"`
	Name         string  `json:"name"`
	Created      string  `json:"created"`
	Rating       float64 `json:"rating"`
	Booleandesu  bool    `json:"booleandesu"`
	Created_date string  `json:"created_date"`
}

// Custom response
type ExampleId struct {
	ExampleId int64 `json:"exampleId"`
}