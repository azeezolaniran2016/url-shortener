package mongo

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Config struct {
	Address      string `envconfig:"MONGO_ADDRESS" default:"mongodb://localhost:27017"`
	DatabaseName string `envconfig:"MONGO_DATABASE_NAME" default:"url_shortener"`
}

type Repository struct {
	*Config
	*mongo.Database
}

func New() (db *Repository, err error) {
	cfg, err := parseConfig("")
	if err != nil {
		return
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, cfg.Address)
	if err != nil {
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	db = &Repository{
		cfg,
		client.Database(cfg.DatabaseName),
	}

	return
}

// parseConfig parses the config from the enviroment
func parseConfig(prefix string) (cfg *Config, err error) {
	cfg = &Config{}
	err = envconfig.Process(prefix, cfg)
	return
}
