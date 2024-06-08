package services

import (
	"context"
	"errors"
	"go-crud/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	_,err := us.userCollection.InsertOne(us.ctx,user)

	return err
}

func (us *UserServiceImpl) GetUser(userName *string) (*models.User,error) {
	var user *models.User

	query := bson.D{bson.E{Key: "user_name", Value: userName}}
	err := us.userCollection.FindOne(us.ctx, query).Decode(&user)

	return user , err
}

func (us *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{
		Key: "$set", 
		Value: bson.D{bson.E{Key: "user_name", Value: user.Name }, 
						bson.E{Key: "user_age", Value: user.Age},
		}}}
	
	result,_ := us.userCollection.UpdateOne(us.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no document found to update")
	}
	return nil
}

