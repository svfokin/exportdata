package main

import (
	"exportdata/pkg/repository"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	} else {
		sdn := db.DriverName()
		fmt.Println(sdn)

	}

	abon1, err := repository.NewSelectDB(db)
	if err != nil {
		log.Fatalf("failed to SELECT: %s", err.Error())
	} else {
		fmt.Println(abon1[0])
	}

	err = repository.NewQueryDB(db)

	if err != nil {
		log.Fatalf("failed to QUERY: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
