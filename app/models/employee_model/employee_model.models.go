package employeemodel_model

import (
	"context"

	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertNewEmployeeDatabase(query string, data InsertNewEmployeeStruct) (int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	result, err := db.ExecContext(ctx, query, data.Name,
		data.Nip,
		data.Bidang,
		data.Seksi,
		data.UnitKerja,
		data.Gender,
		data.BirthPlace,
		data.BirthDate,
		data.Phone,
		data.Email,
		data.Created_date)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertId, nil
}
