package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server Server
	Mongo  Mongo
}

type Server struct {
	Port string
}

type Mongo struct {
	URI      string
	Database string
	Username string
	Password string
}

func InitConfig(folder, file string) (*Config, error) {
	cfg := new(Config)

	// viper.AddConfigPath(folder)
	// viper.SetConfigName(file)
	// if err := viper.ReadInConfig(); err != nil {
	// 	return nil, err
	// }
	// if err := viper.Unmarshal(cfg); err != nil {
	// 	return nil, err
	// }

	cfg, err := getEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func getEnv(cfg *Config) (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cfg.Mongo.URI = os.Getenv("MONGO_INITDB_ROOT_URI")
	cfg.Mongo.Database = os.Getenv("MONGO_INITDB_ROOT_DATABASE")
	cfg.Mongo.Username = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	cfg.Mongo.Password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	return cfg, nil
}
