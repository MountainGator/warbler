package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/MountainGator/warbler/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WarbleServiceImpl struct {
	warblecollection *mongo.Collection
	ctx              context.Context
}

func NewWarbleService(warblecollection *mongo.Collection, ctx context.Context) WarbleService {
	return &WarbleServiceImpl{
		warblecollection: warblecollection,
		ctx:              ctx,
	}
}

func (w *WarbleServiceImpl) CreateWarble(warble *models.Warble) error {
	_, err := w.warblecollection.InsertOne(w.ctx, warble)
	if err != nil {
		fmt.Println("warble insertion err", err)
		return err
	}

	return nil
}
func (w *WarbleServiceImpl) EditWarble(warble *models.Warble) error {
	filter := bson.D{primitive.E{Key: "_id", Value: warble.Id}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "content", Value: warble.Content},
			},
		},
	}

	result, _ := w.warblecollection.UpdateOne(w.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("couldn't find user")
	}

	return nil
}
func (w *WarbleServiceImpl) FindAll() ([]*models.Warble, error) {
	var (
		results     []bson.D
		warble_list []*models.Warble
		cursor      *mongo.Cursor
		err         error
	)

	cursor, err = w.warblecollection.Find(w.ctx, bson.D{{}})
	if err != nil {
		fmt.Println("find warble impl error")
		return nil, err
	}

	if err = cursor.All(w.ctx, &results); err != nil {
		return nil, err
	}

	if err = cursor.Close(w.ctx); err != nil {
		return nil, err
	}

	for _, warble := range results {
		var each *models.Warble
		bytes, _ := bson.Marshal(warble)
		bson.Unmarshal(bytes, &each)
		warble_list = append(warble_list, each)
	}

	return warble_list, nil
}
func (w *WarbleServiceImpl) FindUserWarbles(user_id *string) ([]*models.Warble, error) {
	var warble_list []*models.Warble
	return warble_list, nil
}
func (w *WarbleServiceImpl) DeleteWarble(warble_id *string) error {
	filter := bson.D{primitive.E{Key: "id", Value: warble_id}}
	result, _ := w.warblecollection.DeleteOne(w.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("error. could not delete warble")
	}
	return nil
}
