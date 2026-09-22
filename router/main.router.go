package router

import "github.com/gin-gonic/gin"

func MainRouter(router *gin.Engine) {
	UserRouter(router)
}
