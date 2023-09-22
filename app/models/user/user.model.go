package user_model

import (
	"context"

	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertNewUserDatabase(query string, data InsertNewUserStruct) (int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	result, err := db.ExecContext(ctx, query, data.Name, data.Age, data.Created_date)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertId, nil

}

func SelectAllFromUser() ([]GetAllUserHandlerStruct, error) {
	sqlQuery := "SELECT user.id, user.name, user.age, user.created_date FROM user AS user"

	var result []GetAllUserHandlerStruct
	if err := database.QuerySelectWithoutCondition(sqlQuery, &result); err != nil {
		return nil, err
	}
	return result, nil
}
