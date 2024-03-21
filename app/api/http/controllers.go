package httpApi

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"user-service/app/api/http/registration"
)

func Controllers(r *gin.Engine, db *sql.DB) {

	InitMiddlewares(r, db)

	r.POST("/registration/signUp", registration.SignUp)
}
