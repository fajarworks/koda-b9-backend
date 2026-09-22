package router

import (
	"github.com/fajarworks/koda-b9-backend/handler"
	"github.com/fajarworks/koda-b9-backend/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	userRouter := r.Group("/auth")
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	userRouter.POST("/login", userHandler.Login)
	userRouter.POST("/register", userHandler.Register)

}
