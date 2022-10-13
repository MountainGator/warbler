package services

import (
	"context"
	"errors"

	"github.com/MountainGator/warbler/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	var (
		temp *models.User
		err  error
	)
	query := bson.D{bson.E{Key: "username", Value: user.Username}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&temp)

	if err != nil {
		_, er := u.usercollection.InsertOne(u.ctx, user)
		session, ses_err := u.store.Get(c.Request, "session")
		if ses_err != nil {
			return ses_err
		}
		session.Values["user"] = temp.Username
		session.Save(c.Request, c.Writer)
		if er != nil {
			return er
		}

	} else if temp.Username == user.Username {
		return errors.New("user already exists")
	}
	return nil
}
func (u *UserServiceImpl) UserLogin(name *string, pwd string, c *gin.Context) error {
	var (
		user *models.User
		err  error
	)
	query := bson.D{bson.E{Key: "username", Value: name}}
	if err = u.usercollection.FindOne(u.ctx, query).Decode(&user); err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd))

	if err != nil {
		return err
	}

	session, ses_err := u.store.Get(c.Request, "session")
	if ses_err != nil {
		return ses_err
	}

	session.Values["user"] = name
	session.Save(c.Request, c.Writer)
	return nil
}
func (u *UserServiceImpl) GetUserDetails(*string) (*models.User, error) {
	var user *models.User
	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	if err := u.usercollection.FindOne(u.ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UserServiceImpl) Logout(c *gin.Context) error {
	session, err := u.store.Get(c.Request, "session")
	if err != nil {
		return err
	}
	session.Values["user"] = nil
	session.Save(c.Request, c.Writer)
	return nil
}

func (u *UserServiceImpl) UpdateUser(*models.User) error {
	return nil
}
func (u *UserServiceImpl) DeleteUser(*string) error {
	return nil
}

func (w *WarbleServiceImpl) CreateWarble(warble *models.Warble) error {
	return nil
}
func (w *WarbleServiceImpl) EditWarble(warble *models.Warble) error {
	return nil
}
func (w *WarbleServiceImpl) FindAll() ([]*models.Warble, error) {
	var warble_list []*models.Warble
	return warble_list, nil
}
func (w *WarbleServiceImpl) FindUserWarbles(user_id *string) ([]*models.Warble, error) {
	var warble_list []*models.Warble
	return warble_list, nil
}
