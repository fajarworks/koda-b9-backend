package main

import (
	"github.com/fajarworks/koda-b9-backend/router"
	"github.com/gin-gonic/gin"
)

// type User struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func main() {
	r := gin.Default()
	router.MainRouter(r)

	r.Run("localhost:6000")
}
