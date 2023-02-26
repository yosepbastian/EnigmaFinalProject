package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataSourceName string
}

func (c *Config) InitDb() {
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	c.DataSourceName = dsn

}

func NewConfig() Config {
	config := Config{}
	config.InitDb()
	return config

}
