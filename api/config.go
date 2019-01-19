package api

// Config is the api configuration struct
type Config struct {
	Port    string `envconfig:"PORT" default:"80"`
	APIMode bool   `encconfig:"API_MODE" default:"false"`
}

