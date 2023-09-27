package entity

import (
	"regexp"
	"strings"
	"time"

	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
)

type Employee struct {
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

type EmployeeWithId struct {
	Id           int    `json:"id"`
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

type PaginationEmply struct {
	Data      []EmployeeWithId `json:"data"`
	TotalRow  int              `json:"total_row"`
	TotalPage int              `json:"total_page"`
	NoPerPage int              `json:"no_per_page"`
}

func (employee Employee) ValidationEmployee() map[string]interface{} {
	var errorMessage map[string]interface{}

	// validasi nama dikirim datanya
	employee.Name = (strings.TrimSpace(employee.Name))
	employee.Nip = (strings.TrimSpace(employee.Nip))
	employee.Bidang = (strings.TrimSpace(employee.Bidang))
	employee.Seksi = (strings.TrimSpace(employee.Seksi))
	employee.UnitKerja = (strings.TrimSpace(employee.UnitKerja))

	if globalFunction.IsEmpty(employee.Name) {
		errorMessage = globalFunction.GetMessage("emply002", nil)
	} else if globalFunction.IsEmpty(employee.Nip) {
		errorMessage = globalFunction.GetMessage("emply003", nil)
	} else if !isValidNIP(employee.Nip) {
		errorMessage = globalFunction.GetMessage("emply011", nil)
	} else if globalFunction.IsEmpty(employee.Bidang) { // validasi bidang dikirim datanya
		errorMessage = globalFunction.GetMessage("emply004", nil)
	} else if globalFunction.IsEmpty(employee.Seksi) { // validasi seksi dikirim datanya
		errorMessage = globalFunction.GetMessage("emply005", nil)
	} else if globalFunction.IsEmpty(employee.UnitKerja) { // validasi unit kerja dikirim datanya
		errorMessage = globalFunction.GetMessage("emply006", nil)
	} else if globalFunction.IsEmpty(employee.Gender) { // validasi gender dikirim datanya
		return globalFunction.GetMessage("emply007", nil)
	} else if !isValidGender(employee.Gender) { // validasi cek gender yang ada
		errorMessage = globalFunction.GetMessage("emply009", nil)
	} else if !isValidBirthDate(employee.BirthDate) { // validasi format tanggal
		errorMessage = globalFunction.GetMessage("emply010", nil)
	} else if !globalFunction.IsValidEmail(strings.TrimSpace(employee.Email)) { // validasi format email
		errorMessage = globalFunction.GetMessage("auth002", nil)
	} else {
		errorMessage = nil
	}

	return errorMessage

}

// Fungsi untuk memvalidasi jenis kelamin, 1 = laki-laki - 2 =perempuan
func isValidGender(gender int) bool {
	if gender == 1 || gender == 2 {
		return true
	}
	return false
}

// Fungsi untuk memeriksa apakah birth_date valid sesuai dengan format yang diharapkan
func isValidBirthDate(birthDate string) bool {
	inputDate, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return false
	}

	today := time.Now()

	return inputDate.Before(today) || inputDate.Equal(today)
}

func isValidNIP(nip string) bool {
	// Gunakan ekspresi reguler untuk memeriksa format NIP
	re := regexp.MustCompile(`^\d{8} \d{6} \d{1} \d{3}$`)
	return re.MatchString(nip)
}
