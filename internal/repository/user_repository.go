package repository

import (
	"context"
	"errors"

	"sea-cinema-api/internal/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	UpdateUser(ctx context.Context, id string, user model.User) error
}

type userRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		database:   database,
		collection: database.Collection(collectionName),
	}
}

func (repository *userRepository) GetUserById(ctx context.Context, id string) (model.User, error) {
	var user model.User

	err := repository.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return model.User{}, err
	}

	return user, nil
}

func (repository *userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User

	err := repository.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return model.User{}, err
	}

	return user, nil
}

func (repository *userRepository) CreateUser(ctx context.Context, user model.User) error {
	_, err := repository.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) UpdateUser(ctx context.Context, id string, user model.User) error {
	_, err := repository.collection.ReplaceOne(ctx, bson.M{"_id": id}, user)
	if err != nil {
		return err
	}

	return nil
}
