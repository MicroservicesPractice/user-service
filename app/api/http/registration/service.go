package registration

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) CreateUser(c *gin.Context, user *User) error {
	db := c.MustGet("db").(*sql.DB)

	sqlStatement := `INSERT INTO users (email, password, nickname, phone_number)
	VALUES ($1, $2, $3, $4);`

	_, err := db.Exec(sqlStatement, user.Email, user.Password, user.Nickname, user.PhoneNumber)

	return err
}

var UserServiceInstance = UserService{}
