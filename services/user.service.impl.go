package services

import (
	"context"
	"errors"
	"fmt"

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
		n00b, berr := bcrypt.GenerateFromPassword([]byte(user.Pwd), 14)
		if berr != nil {
			fmt.Println("error hashing pwd", berr)
			return berr
		}
		user.Pwd = string(n00b)
		_, er := u.usercollection.InsertOne(u.ctx, user)
		if er != nil {
			fmt.Println("insertion err", er)
			return er
		}
		session, ses_err := u.store.Get(c.Request, "session")
		if ses_err != nil {
			fmt.Println("session err", ses_err)
			return ses_err
		}
		session.Values["user"] = user.Username
		session.Save(c.Request, c.Writer)
	} else {
		return errors.New("user already exists")
	}

	return nil
}
func (u *UserServiceImpl) UserLogin(data *models.Login, c *gin.Context) error {
	var (
		user *models.User
		err  error
	)
	query := bson.D{bson.E{Key: "username", Value: data.Username}}
	if err = u.usercollection.FindOne(u.ctx, query).Decode(&user); err != nil {
		fmt.Println("find user error", err)
		return err
	}
	fmt.Println("found user data", user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(data.Pwd))

	if err != nil {
		fmt.Println("hash pwd error", err)

		return err
	}

	session, ses_err := u.store.Get(c.Request, "session")
	if ses_err != nil {
		fmt.Println("store session error")
		return ses_err
	}

	session.Values["user"] = user.Username
	session.Save(c.Request, c.Writer)
	return nil
}
func (u *UserServiceImpl) GetUserDetails(name *string) (*models.User, error) {
	var user *models.User
	filter := bson.D{primitive.E{Key: "username", Value: name}}
	if err := u.usercollection.FindOne(u.ctx, filter).Decode(&user); err != nil {
		fmt.Println("error finding user", err)
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

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{primitive.E{Key: "_id", Value: user.Id}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "username", Value: user.Username},
			},
		},
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "email", Value: user.Email},
			},
		},
	}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("couldn't find user")
	}
	return nil
}
func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("error. could not delete user")
	}
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
