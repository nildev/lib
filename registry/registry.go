package registry

import (
	"os"

	"github.com/nildev/lib/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

var (
	mongoDBClientURL = "mongodb://localhost:27017/test"
	databaseName     = "default"
)

// These constants is public API, being taken by lib from environment
const (
	NDMongoDBURL   = "ND_MONGODB_URL"
	NDDatabaseName = "ND_DATABASE_NAME"
)

// GetMongoDBClient create mongoDB client based on client
func GetMongoDBClient() (*mgo.Session, error) {
	envValue := os.Getenv(NDMongoDBURL)
	if envValue != "" {
		mongoDBClientURL = envValue
	}

	session, err := mgo.Dial(mongoDBClientURL)
	//	session.SetSafe(&mgo.Safe{})
	return session, err
}

// GetDatabaseName returns DB name
func GetDatabaseName() string {
	envValue := os.Getenv(NDDatabaseName)
	if envValue != "" {
		databaseName = envValue
	}
	return databaseName
}
