package registry

import (
	"os"

	"github.com/iron-io/iron_go3/mq"
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
	NDEnv          = "ND_ENV"

	EnvDev = "dev"
)

// GetEnv return environment
func GetEnv() string {
	env := os.Getenv(NDEnv)
	if env == "" {
		env = EnvDev
	}

	return env
}

// CreateMongoDBClient create MongoDB client
func CreateMongoDBClient() (*mgo.Session, error) {
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

// CreateIronQueue creates and returns Iron.io queue
func CreateIronQueue(name string) (*mq.Queue, error) {
	subscribers := []mq.QueueSubscriber{}
	subscription := mq.PushInfo{
		Retries:      3,
		RetriesDelay: 60,
		ErrorQueue:   "error_queue",
		Subscribers:  subscribers,
	}
	queue_type := "push"
	queueInfo := mq.QueueInfo{Type: queue_type, MessageExpiration: 60, MessageTimeout: 56, Push: &subscription}
	_, err := mq.CreateQueue(name, queueInfo)
	if err != nil {
		return nil, err
	}

	queue := mq.New(name)
	return &queue, nil
}
