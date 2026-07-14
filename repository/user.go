package repository

import (
	"context"
	"errors"

	"github.com/durgaprasad97005/bikeFinance/database"
	"github.com/durgaprasad97005/bikeFinance/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// struct for UserRepository
type UserRepository struct{
	collection *mongo.Collection
}

// Constructor like func to initialize UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection(database.UsersCollection), 
	}
}

// Get all Users
func (r *UserRepository) Find(ctx context.Context, filter bson.M) ([]models.User, error) {
	// Pending - implementation
	return make([]models.User, 0), nil 
}

// Get a User 
func (r *UserRepository) Get(ctx context.Context, filter bson.M) (*models.User, error) {
	// Find in the collection
	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	// If document doesn't exist
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	// If some other error
	if err != nil {
		return nil, err
	}

	// Successfully returned user
	return &user, nil
}

// Create a User and populate ID field in user instance
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	// Insert into database
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	// check whether we got InsertedId in result
	if result.InsertedID.(bson.ObjectID) == bson.NilObjectID {
		return errors.New("Failed user creation")
	}

	// Update ID field in the user struct
	user.ID = result.InsertedID.(bson.ObjectID)
	return nil
}

// Update a User
func (r *UserRepository) Update(ctx context.Context, filter bson.M, user *models.User) error {
	// Pending implementation
	return nil
}

// Delete a User
func (r *UserRepository) Delete(ctx context.Context, filter bson.M) error {
	// Pending implementation
	return nil
}