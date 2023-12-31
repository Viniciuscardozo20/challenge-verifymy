package mongodb

import (
	"challenge-verifymy/core/ports"
	"challenge-verifymy/customerror"
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client represents a MongoDB database connection and provides methods to interact with the database.
type Client struct {
	client       *mongo.Client
	db           *mongo.Database
	disconnected chan struct{}
}

const (
	connectionTimeout = 30 * time.Second
	pingTimeout       = 5 * time.Second
)

// Disconnected returns a channel that signals when the database has disconnected.
func (c *Client) Disconnected() <-chan struct{} {
	return c.disconnected
}

// disconnect disconnects from the database.
func (c *Client) disconnect(ctx context.Context) {
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	if err := c.client.Disconnect(ctx); err != nil {
		log.Error(errors.Join(err, customerror.ErrFailedToDisconnect))
	}

	close(c.disconnected)
}

// GetRepository retrieves a MongoDB collection as a Repository given its name.
func (c *Client) GetRepository(ctx context.Context, name string) (ports.UserRepository, error) {
	repo := &Repository{coll: c.db.Collection(name)}

	index := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := repo.coll.Indexes().CreateOne(ctx, index)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// New initializes and returns a new MongoDB Client instance by connecting to the provided URI. It also ensures a successful connection by
// pinging the server.
func New(ctx context.Context, uri, database string) (*Client, error) {
	connectCtx, stop := context.WithTimeout(ctx, connectionTimeout)
	defer stop()

	mongoClient, err := mongo.Connect(connectCtx, options.Client().
		ApplyURI(uri),
	)
	if err != nil {
		return nil, errors.Join(err, customerror.ErrFailedToConnect)
	}

	pingCtx, pingStop := context.WithTimeout(ctx, pingTimeout)
	defer pingStop()

	if err = mongoClient.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, errors.Join(err, customerror.ErrFailedToPing)
	}

	client := &Client{
		client:       mongoClient,
		db:           mongoClient.Database(database),
		disconnected: make(chan struct{}),
	}

	go client.disconnect(ctx)

	return client, nil
}
