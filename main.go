package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = []User{
	{
		Name:     "fajar",
		Email:    "fajar@mail.com",
		Password: "12345678910",
	},
}

func main() {
	router := gin.Default()
	router.POST("/login", func(ctx *gin.Context) {
		var form User
		if err := ctx.ShouldBindWith(&form, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "terjadi kesalahan server",
			})
			return
		}

		for _, v := range Users {
			if form.Email != v.Email && form.Password != v.Password {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "email atau password salah",
				})

				return
			}

		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": fmt.Sprint("selamat datang ", form.Name),
		})

	})

	router.POST("/register", func(ctx *gin.Context) {
		var form User

		if err := ctx.ShouldBindWith(&form, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "terjadi kesalahan sever",
			})
			return
		}

		for _, v := range Users {
			if v.Email == form.Email {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "email sudah tersedia",
				})
				return
			}
		}
		if len(form.Password) < 8 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "panjang password harus lebih dari 8",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "register berhasil",
		})
		Users = append(Users, form)
	})

	router.Run("localhost:6000")
}
