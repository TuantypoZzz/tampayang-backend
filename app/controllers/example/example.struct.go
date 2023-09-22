package example_controller


type CreateExampleStruct struct {
	Name			string		`json:"name"`
	Created			string		`json:"created"`
	Rating			float64		`json:"rating"`
	Booleandesu		bool		`json:"booleandesu"`
	Created_date	string		`json:"created_date"`
}

// Custom response
type CreateExampleResponse struct {
	ExampleId	int64		`json:"exampleId"`
}