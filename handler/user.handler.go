package handler

import (
	"fmt"
	"net/http"

	"github.com/fajarworks/koda-b9-backend/dto"
	"github.com/fajarworks/koda-b9-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

var Users = []dto.User{}

func (h *UserHandler) Login(ctx *gin.Context) {
	var form dto.User
	if err := ctx.ShouldBindWith(&form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "terjadi kesalahan server",
		})
		return
	}

	if err := h.service.WrongEmailAndPass(form); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := h.service.LengthPassword(form); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprint("selamat datang ", form.Name),
	})

}

func (h *UserHandler) Register(ctx *gin.Context) {
	var form dto.User

	if err := ctx.ShouldBindWith(&form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "terjadi kesalahan sever",
		})
		return
	}

	if err := h.service.EmailExist(form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := h.service.LengthPassword(form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "register berhasil",
	})
	Users = append(Users, form)
}
