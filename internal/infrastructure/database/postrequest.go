package database

import (
	"context"
	"time"

	"github.com/bhushan-aruto/justarquest/internal/entity"
)

func (db *MongoDBdatabase) CreateUser(justrequeest *entity.Jsutrequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("internship")
	_, err := collection.InsertOne(ctx, justrequeest)

	return err
}
