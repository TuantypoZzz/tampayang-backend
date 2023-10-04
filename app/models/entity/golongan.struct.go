package entity

import (
	"regexp"
	"strings"

	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
)

type Golongan struct {
	GolonganName string `json:"golongan_name"`
	Created_date string `json:"created_date"`
}

func (golongan Golongan) ValidationGolongan() map[string]interface{} {
	var errMessage map[string]interface{}
	// Validasi nama yang dikirim
	golongan.GolonganName = (strings.TrimSpace(golongan.GolonganName))

	if globalFunction.IsEmpty(golongan.GolonganName) {
		errMessage = globalFunction.GetMessage("gol002", nil)
	} else if !IsValidGolonganFormat(golongan.GolonganName) {
		errMessage = globalFunction.GetMessage("gol003", nil)
	} else {
		errMessage = nil
	}

	return errMessage
}

// Fungsi untuk memvalidasi format golongan.
func IsValidGolonganFormat(golongan string) bool {
	// Buat ekspresi reguler untuk memeriksa format yang diinginkan
	regex := regexp.MustCompile(`^[IVXLCDM]+/[a-z]$`)
	return regex.MatchString(golongan)
}
