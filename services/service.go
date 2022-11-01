package services

import (
	"github.com/MountainGator/warbler/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(*models.User, *gin.Context) error
	UserLogin(*models.Login, *gin.Context) error
	GetUserDetails(*string) (*models.User, error)
	Logout(*gin.Context) error
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type WarbleService interface {
	CreateWarble(*models.Warble) error
	EditWarble(*models.Warble) error
	FindAll() ([]*models.Warble, error)
	FindUserWarbles(*string) ([]*models.Warble, error)
	DeleteWarble(*string) error
}
