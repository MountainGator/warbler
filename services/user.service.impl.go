package services

import (
	"context"

	"github.com/MountainGator/warbler/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	store          *sessions.CookieStore
	ctx            context.Context
}

type WarbleServiceImpl struct {
	warblecollection *mongo.Collection
	ctx              context.Context
}

func NewUserService(usercollection *mongo.Collection, store *sessions.CookieStore, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		store:          store,
		ctx:            ctx,
	}
}

func NewWarbleService(warblecollection *mongo.Collection, ctx context.Context) WarbleService {
	return &WarbleServiceImpl{
		warblecollection: warblecollection,
		ctx:              ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User, c *gin.Context) error {

}
func (u *UserServiceImpl) UserLogin(name *string, pwd string, c *gin.Context) error {

}
func (u *UserServiceImpl) GetUserDetails(*string) (*models.User, error) {
	var user *models.User
	return user, nil
}
func (u *UserServiceImpl) Logout(*gin.Context) error {
	return nil
}
func (u *UserServiceImpl) UpdateUser(*models.User) error {
	return nil
}
func (u *UserServiceImpl) DeleteUser(*string) error {
	return nil
}

func (u *WarbleServiceImpl) CreateWarble(warble *models.Warble) error {
	return nil
}
func (u *WarbleServiceImpl) EditWarble(warble *models.Warble) error {
	return nil
}
func (u *WarbleServiceImpl) FindAll() ([]*models.Warble, error) {
	var warble_list []*models.Warble
	return warble_list, nil
}
func (u *WarbleServiceImpl) FindUserWarbles(user_id *string) ([]*models.Warble, error) {
	var warble_list []*models.Warble
	return warble_list, nil
}
