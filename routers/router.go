package routers

import (
	"github.com/gin-gonic/gin"
	"sunderer/service"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/searchUser", service.GetUserInfo)

	return r
}
