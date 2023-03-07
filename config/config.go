package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Config struct {
	DataSourceName string
	TokenConfig
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
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

	c.TokenConfig = TokenConfig{
		ApplicationName:     "stockbite",
		JwtSignatureKey:     "DWici392-sl93wcFD@",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: 20 * time.Minute,
	}

}

func NewConfig() Config {
	config := Config{}
	config.InitDb()
	return config

}
