package controllers

import (
	"fmt"
	"net/http"

	"github.com/MountainGator/warbler/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type UserController struct {
	UserService   services.UserService
	WarbleService services.WarbleService
	store         *sessions.CookieStore
}

func NewUserController(userservice services.UserService, warbleservice services.WarbleService, store *sessions.CookieStore) UserController {
	return UserController{
		UserService:   userservice,
		WarbleService: warbleservice,
		store:         store,
	}
}

func (uc *UserController) Auth(c *gin.Context) {
	session, _ := uc.store.Get(c.Request, "session")
	_, ok := session.Values["user"]

	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"Error": "Not logged in"})
		c.Abort()
		return
	}
	c.Next()
}

func (uc *UserController) GetCreds(c *gin.Context) {
	session, _ := uc.store.Get(c.Request, "session")
	_, ok := session.Values["user"]

	if !ok {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": "Not logged in"})
		return
	}
	fmt.Println("session user", ok)
	c.JSON(http.StatusAccepted, gin.H{"user": ok})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"success": "user created"})
}
func (uc *UserController) UserLogin(c *gin.Context) {

}
func (uc *UserController) GetUserDetails(c *gin.Context) {

}
func (uc *UserController) Logout(c *gin.Context) {

}
func (uc *UserController) UpdateUser(c *gin.Context) {

}
func (uc *UserController) DeleteUser(c *gin.Context) {

}
func (uc *UserController) CreateWarble(c *gin.Context) {

}
func (uc *UserController) EditWarble(c *gin.Context) {

}
func (uc *UserController) FindAll(c *gin.Context) {

}
func (uc *UserController) FindUserWarbles(c *gin.Context) {

}
func (uc *UserController) DeleteWarble(c *gin.Context) {

}
