package employeemodel_model

type InsertNewEmployeeStruct struct {
	Name         string `json:"name"`
	Nip          string `json:"nip"`
	Bidang       string `json:"bidang"`
	Seksi        string `json:"seksi"`
	UnitKerja    string `json:"unit_kerja"`
	Gender       int    `json:"gender"`
	BirthPlace   string `json:"birth_place"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Created_date string `json:"created_date"`
}
