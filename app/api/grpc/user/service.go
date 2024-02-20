package user

import (
	"user-service/app/config/initializers"
	"user-service/app/types"
)

type UserService struct {
}

func (u *UserService) CreateUser(user *types.User) error {
	sqlStatement := `INSERT INTO users (email, password, nickname, phone_number)
	VALUES ($1, $2, $3, $4);`

	_, err := initializers.DB.Exec(sqlStatement, user.Email, user.Password, user.Nickname, user.PhoneNumber)

	return err
}

func (u *UserService) GetUserPassword(email string) (types.User, error) {
	result := types.User{}

	sqlStatement := `SELECT password, id
	FROM users
	WHERE users.email=$1;`

	err := initializers.DB.QueryRow(sqlStatement, email).Scan(&result.Password, &result.ID)

	return result, err
}

var UserServiceInstance = UserService{}
