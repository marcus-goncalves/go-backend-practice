package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
}

type conn struct {
	session  *mongo.Client
	database *mongo.Database
}

func NewConnection(cfg Config) (Connection, error) {
	log.Println("Connecting to: ", cfg.GetDbName())
	dbSession, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.GetDsn()))
	if err != nil {
		log.Fatalln("Couldn´t connect to mongoDB")
	}

	return &conn{
		session:  dbSession,
		database: dbSession.Database(cfg.GetDbName()),
	}, nil

}

func (c *conn) Close() {
	c.session.Disconnect(context.TODO())
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
