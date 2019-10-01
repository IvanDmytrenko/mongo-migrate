// +build integration

package migrate

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const globalTestCollection = "test-global"

func init() {
	Register(func(db *mongo.Database) error {
		return db.C(globalTestCollection).Insert(bson.M{"a": "b"})
	}, func(db *mongo.Database) error {
		return db.C(globalTestCollection).Remove(bson.M{"a": "b"})
	})
}
