package mongodb

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	GetDsn() string
	GetDbName() string
}

type config struct {
	user   string
	pwd    string
	host   string
	port   int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config

	cfg.dbName = os.Getenv("MONGO_DATABASE")
	cfg.user = os.Getenv("MONGO_USER")
	cfg.pwd = os.Getenv("MONGO_PWD")
	cfg.host = os.Getenv("MONGO_HOST")

	var err error
	cfg.port, err = strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatalln("Error getting database port", err)
	}

	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.user, cfg.pwd, cfg.host, cfg.port, cfg.dbName)

	return &cfg
}

func (c *config) GetDsn() string {
	return c.dsn
}

func (c *config) GetDbName() string {
	return c.dbName
}
