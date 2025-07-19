package models

import (
	"context"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

func GetUserLoginByEmail(userEmail string) entity.UserLogin {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var userLogin entity.UserLogin

	sqlQuery := "SELECT usr.user_id, usr.user_name, usr.user_email, usr.user_password, usr.user_role FROM user AS usr WHERE usr.user_email = ?"
	result, err := db.QueryContext(ctx, sqlQuery, userEmail)
	if err != nil {
		panic("models - GetUserLoginByEmail, db.QueryContext " + err.Error())
	}

	if result.Next() {
		err := result.Scan(
			&userLogin.User_id,
			&userLogin.User_name,
			&userLogin.User_email,
			&userLogin.User_password,
			&userLogin.User_role)

		if err != nil {
			panic("models - getUserLoginByEmail, result.Scan " + err.Error())
		}
	}
	return userLogin
}
