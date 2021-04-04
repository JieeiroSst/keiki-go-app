package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	mutex sync.Mutex
	instance *MongoConn
)

type Config struct {
	DSN string
}

type MongoConn struct {
	db *mongo.Database
}

func GetMongoConnInstance(cf Config) *MongoConn {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn:=cf.DSN
			clientOptions:=options.Client().ApplyURI(dsn)
			client,err:=mongo.Connect(context.TODO(),clientOptions)
			if err!=nil{
				log.Println(err)
			}
			err=client.Ping(context.TODO(),nil)
			if err!=nil{
				log.Println(err)
			}
			db:=client.Database("db")
			instance = &MongoConn{
				db: db,
			}
		}
	}
	return instance
}

func NewMongoRepo(cf Config) *MongoConn {
	return &MongoConn{
		db :GetMongoConnInstance(cf).db,
	}
}