package middleware

import (
	"fmt"
	"net/http"
	"order_service/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
		fmt.Println(os.Getenv("USERNAME"))
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			response := common.BuildErrorResponse("Invalid Credentials", "Credentials are not of type BASE64", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		isValid := (username == os.Getenv("USERNAME")) && (password == os.Getenv("PASSWORD"))
		if !isValid {
			response := common.BuildErrorResponse("Invalid Credentials", "Username or Password is not correct", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}

}
