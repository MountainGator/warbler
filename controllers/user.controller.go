package controllers

import (
	"fmt"
	"net/http"

	"github.com/MountainGator/warbler/models"
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
	fmt.Println("running auth")
	session, _ := uc.store.Get(c.Request, "session")
	user, ok := session.Values["user"]

	if !ok || user == nil {
		fmt.Println("error getting session value")
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
	c.JSON(http.StatusAccepted, gin.H{"user": ok})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	var err error
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "can't bind JSON"})
		return
	}

	if err = uc.UserService.CreateUser(&user, c); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": "user created"})
}
func (uc *UserController) UserLogin(c *gin.Context) {
	var login_data models.Login
	var err error

	if err = c.ShouldBindJSON(&login_data); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "couldn't bind json"})
		return
	}
	if err = uc.UserService.UserLogin(&login_data, c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "incorrect username or password"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": "logged in"})
}
func (uc *UserController) GetUserDetails(c *gin.Context) {
	var (
		user *models.User
		err  error
	)
	username := c.Param("name")
	user, err = uc.UserService.GetUserDetails(&username)

	if err != nil {
		fmt.Println("error finding user (controller)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	user.Pwd = ""
	c.JSON(http.StatusOK, gin.H{"data": user})

}
func (uc *UserController) Logout(c *gin.Context) {
	if err := uc.UserService.Logout(c); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error logging out": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "logged out"})
}
func (uc *UserController) UpdateUser(c *gin.Context) {
	var (
		user *models.User
		err  error
	)

	err = c.ShouldBindJSON(&user)

	if err = uc.UserService.UpdateUser(user); err != nil {
		fmt.Println("error updating user", err)
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "updated user"})
}
func (uc *UserController) DeleteUser(c *gin.Context) {

	username := c.Param("name")

	if err := uc.UserService.DeleteUser(&username); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "deleted user"})
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
