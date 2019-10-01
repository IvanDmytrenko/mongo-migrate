// +build integration

package migrate

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const globalTestIndexName = "test_idx_2"

func init() {
	Register(func(db *mongo.Database) error {
		return db.Collection(globalTestCollection).EnsureIndex(mongo.Index{Name: globalTestIndexName, Key: []string{"a"}})
	}, func(db *mongo.Database) error {
		return db.Collection(globalTestCollection).DropIndexName(globalTestIndexName)
	})
}
