package controllers

import (
	"github.com/MountainGator/warbler/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService   services.UserService
	WarbleService services.WarbleService
}

func NewUserController(userservice services.UserService, warbleservice services.WarbleService) UserController {
	return UserController{
		UserService:   userservice,
		WarbleService: warbleservice,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {

}
