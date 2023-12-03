package mongodb

import (
	"challenge-verifymy/customerror"
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
func (r *Repository) Save(ctx context.Context, data, output any) error {
	result, err := r.coll.InsertOne(ctx, data)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return customerror.ErrUserAlreadyExists
		}
		return errors.Join(err, customerror.ErrFailedToInsertDocument)
	}

	err = r.Read(ctx, result.InsertedID.(primitive.ObjectID).Hex(), output)
	if err != nil {
		return err
	}

	return nil
}

// Read retrieves a single document from the repository based on the given id.
func (r *Repository) Read(ctx context.Context, id string, output any) error {
	filter, err := buildFilterByID(id)
	if err != nil {
		return err
	}

	found := r.coll.FindOne(ctx, filter)
	if found.Err() != nil {
		if errors.Is(found.Err(), mongo.ErrNoDocuments) {
			return customerror.ErrFailedToFindDocument
		}

		return errors.Join(found.Err(), customerror.ErrFailedToFindDocument)
	}

	if err := found.Decode(output); err != nil {
		return errors.Join(err, customerror.ErrFailedToUnmarshalDocument)
	}

	return nil
}

// ReadAll retrieves a list of documents from the repository based on the given id.
func (r *Repository) ReadAll(ctx context.Context, output any) error {
	findOptions := options.Find()

	cursor, err := r.coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customerror.ErrNoResult
		}

		return errors.Join(err, customerror.ErrFailedToFindDocument)
	}

	if err = cursor.All(ctx, output); err != nil {
		return errors.Join(err, customerror.ErrFailedToUnmarshalDocument)
	}

	return nil
}

// Update changes from a single document into the repository.
func (r *Repository) Update(ctx context.Context, id string, data, output any) error {
	filter, err := buildFilterByID(id)
	if err != nil {
		return err
	}

	_, err = r.coll.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return err
	}

	err = r.Read(ctx, id, output)
	if err != nil {
		return err
	}

	return nil
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
			return customerror.ErrNoResult
		}

		return errors.Join(err, customerror.ErrFailedToFindDocument)
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
