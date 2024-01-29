package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Addr     string
	Port     int
	Username string
	Password string
}

func NewMongo(c *Config) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	credential := options.Credential{
		Username: c.Username,
		Password: c.Password,
	}
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", c.Addr, c.Port)).SetAuth(credential))
	if err != nil {
		panic(err)
	}
	return cli
}
