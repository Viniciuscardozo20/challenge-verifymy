package mongodb

import (
	"challenge-verifymy/customerr"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository provides methods to interact with a specific MongoDB collection.
type Repository struct {
	coll *mongo.Collection
}

// Save inserts a single document into the repository.
func (r *Repository) Save(ctx context.Context, data any) error {
	_, err := r.coll.InsertOne(ctx, data)
	if err != nil {
		return errors.Join(err, customerr.ErrFailedToInsertDocument)
	}

	return nil
}

// Read retrieves a single document from the repository based on the given id.
func (r *Repository) Read(ctx context.Context, id string) (any, error) {
	filter, err := buildFilterByID(id)
	if err != nil {
		return nil, err
	}

	var result bson.M

	found := r.coll.FindOne(ctx, filter)
	if found.Err() != nil {
		if errors.Is(found.Err(), mongo.ErrNoDocuments) {
			return nil, customerr.ErrNoResult
		}

		return nil, errors.Join(found.Err(), customerr.ErrFailedToFindDocument)
	}

	if err := found.Decode(&result); err != nil {
		return nil, errors.Join(err, customerr.ErrFailedToUnmarshalDocument)
	}

	return result, nil
}

// ReadAll retrieves a list of documents from the repository based on the given id.
func (r *Repository) ReadAll(ctx context.Context) (any, error) {
	findOptions := options.Find()

	var results any

	cursor, err := r.coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, customerr.ErrNoResult
		}

		return nil, errors.Join(err, customerr.ErrFailedToFindDocument)
	}

	if err = cursor.All(ctx, results); err != nil {
		return nil, errors.Join(err, customerr.ErrFailedToUnmarshalDocument)
	}

	return results, nil
}

// Update changes from a single document into the repository.
func (r *Repository) Update(ctx context.Context, id string, data any) (any, error) {
	filter, err := buildFilterByID(id)
	if err != nil {
		return nil, err
	}

	_, err = r.coll.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return nil, err
	}

	result, err := r.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete remove a single document into the repository.
func (r *Repository) Delete(ctx context.Context, id string) error {
	filter, err := buildFilterByID(id)
	if err != nil {
		return err
	}

	_, err = r.coll.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customerr.ErrNoResult
		}

		return errors.Join(err, customerr.ErrFailedToFindDocument)
	}

	return err
}

func buildFilterByID(id string) (bson.M, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return bson.M{"_id": objectID}, nil
}
