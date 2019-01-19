package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"
)

const (
	collection = "url"
)

// Data ...
type Data struct {
	ID        string    `bson:"_id" json:"id"`
	URL       string    `bson:"url" json:"url"`
	ShortURL  string    `bson:"shortUrl" json:"shortUrl"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatetAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// UpsertURL inserts a new url
func (r *Repository) UpsertURL(ctx context.Context, data *Data) (id string, err error) {
	filter := bson.M{
		"url": data.URL,
	}

	update := bson.M{
		"$currentDate": bson.M{
			"updatedAt": true,
		},
		"shortUrl": data.ShortURL,
		"$setOnInsert": bson.M{
			"shortUrl": data.ShortURL,
			"url":      data.URL,
			"$currentDate": bson.M{
				"updatedAt": true,
				"createdAt": true,
			},
		},
	}

	upsert := true
	opt := &options.FindOneAndUpdateOptions{
		Upsert: &upsert,
		Projection: bson.M{
			"_id": 1,
		},
	}

	res := r.Collection(collection, nil).FindOneAndUpdate(ctx, filter, update, opt)

	meta := struct {
		ID string `bson:"_id"`
	}{}

	res.Decode(&meta)

	fmt.Printf("inserted dataID: %v ", meta.ID)

	return
}
