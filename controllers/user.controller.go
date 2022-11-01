package controllers

import (
	"fmt"
	"net/http"

	"github.com/MountainGator/warbler/models"
	"github.com/MountainGator/warbler/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type Controller struct {
	UserService   services.UserService
	WarbleService services.WarbleService
	store         *sessions.CookieStore
}

func NewController(userservice services.UserService, warbleservice services.WarbleService, store *sessions.CookieStore) Controller {
	return Controller{
		UserService:   userservice,
		WarbleService: warbleservice,
		store:         store,
	}
}

func (cont *Controller) Auth(c *gin.Context) {
	fmt.Println("running auth")
	session, _ := cont.store.Get(c.Request, "session")
	user, ok := session.Values["user"]

	if !ok || user == nil {
		fmt.Println("error getting session value")
		c.JSON(http.StatusForbidden, gin.H{"Error": "Not logged in"})
		c.Abort()
		return
	}
	c.Next()
}

func (cont *Controller) GetCreds(c *gin.Context) {
	session, _ := cont.store.Get(c.Request, "session")
	_, ok := session.Values["user"]

	if !ok {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": "Not logged in"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": ok})
}

func (cont *Controller) CreateUser(c *gin.Context) {
	var user models.User
	var err error
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "can't bind JSON"})
		return
	}

	if err = cont.UserService.CreateUser(&user, c); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": "user created"})
}
func (cont *Controller) UserLogin(c *gin.Context) {
	var login_data models.Login
	var err error

	if err = c.ShouldBindJSON(&login_data); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "couldn't bind json"})
		return
	}
	if err = cont.UserService.UserLogin(&login_data, c); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "incorrect username or password"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": "logged in"})
}
func (cont *Controller) GetUserDetails(c *gin.Context) {
	var (
		user *models.User
		err  error
	)
	username := c.Param("name")
	user, err = cont.UserService.GetUserDetails(&username)

	if err != nil {
		fmt.Println("error finding user (controller)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	user.Pwd = ""
	c.JSON(http.StatusOK, gin.H{"data": user})

}
func (cont *Controller) Logout(c *gin.Context) {
	if err := cont.UserService.Logout(c); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error logging out": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "logged out"})
}
func (cont *Controller) UpdateUser(c *gin.Context) {
	var (
		user *models.User
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		fmt.Println("error binding json", err)
		c.JSON(http.StatusConflict, gin.H{"error": "error in request json"})
	}

	if err = cont.UserService.UpdateUser(user); err != nil {
		fmt.Println("error updating user", err)
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "updated user"})
}
func (cont *Controller) DeleteUser(c *gin.Context) {

	username := c.Param("name")

	if err := cont.UserService.DeleteUser(&username); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "deleted user"})
}
func (cont *Controller) CreateWarble(c *gin.Context) {
	var (
		new_warble *models.Warble
		err        error
	)

	if err = c.ShouldBindJSON(new_warble); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err = cont.WarbleService.CreateWarble(new_warble); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Success": "warble created"})
}
func (cont *Controller) EditWarble(c *gin.Context) {
	var (
		new_warble *models.Warble
		err        error
	)

	if err = c.ShouldBindJSON(new_warble); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	if err = cont.WarbleService.EditWarble(new_warble); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Success": "warble modified"})

}
func (cont *Controller) FindAll(c *gin.Context) {

	warble_list, err := cont.WarbleService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Success": warble_list})
}
func (cont *Controller) FindUserWarbles(c *gin.Context) {
	user_name := c.Param("name")

	warble_list, err := cont.WarbleService.FindUserWarbles(&user_name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Success": warble_list})
}

func (cont *Controller) DeleteWarble(c *gin.Context) {
	user_id := c.Param("id")
	if err := cont.WarbleService.DeleteWarble(&user_id); err != nil {
		c.JSON(http.StatusConflict, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Success": "Deleted Warble"})
}
