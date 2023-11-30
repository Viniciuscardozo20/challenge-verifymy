package mongodb

import "errors"

var (
	ErrFailedToConnect           = errors.New("failed to connect to database")
	ErrFailedToDisconnect        = errors.New("failed to disconnect from database")
	ErrFailedToFindDocument      = errors.New("failed to find document")
	ErrFailedToInsertDocument    = errors.New("failed to insert document")
	ErrFailedToPing              = errors.New("failed to ping database")
	ErrFailedToUnmarshalDocument = errors.New("failed to unmarshal document")
	ErrNoResult                  = errors.New("no result found")
)