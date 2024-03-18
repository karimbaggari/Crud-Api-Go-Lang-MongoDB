package service

import (
	"context"
	"scriptology/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"errors"
) 

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
        userCollection: userCollection,
        ctx: ctx,
    }
}


func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if (err != nil) {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetUser(firstName *string) (*models.User, error) { 
	var user *models.User
	query := bson.D{bson.E{Key:"user_first_name", Value: firstName}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})
	if err!= nil {
        return nil, err
    }
	for cursor.Next(u.ctx) {
		var user models.User
        if err := cursor.Decode(&user); err!= nil {
            return nil, err
        }
        users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	
	return users,nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_first_name", Value:user.FirstName}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "user_first_name", Value: user.FirstName}, bson.E{Key: "user_last_name", Value: user.LastName}, bson.E{Key: "user_age", Value: user.Age },  bson.E{ Key: "user_address", Value: user.Address}}}}
	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return fmt.Errorf("no matched document found for update")
	}
	return nil
}


func (u *UserServiceImpl) DeleteUser(firstName *string) error {
	filter := bson.D{bson.E{Key: "user_first_name", Value:firstName}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return fmt.Errorf("no matched document found for update")
	}
	return nil
}
