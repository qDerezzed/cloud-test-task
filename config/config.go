package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	// databasePassword := os.Getenv("DB_PASSWORD")
	// if databasePassword == "" {
	// 	log.Fatal("$DB_PASSWORD must be set")
	// }
	databasePassword := "password"

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	cfg := &Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: databasePassword,
	}

	return cfg, nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
