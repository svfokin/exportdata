package main

import (
	"exportdata/pkg/repository"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// catchSig cleans up our websocket conenction if we kill the program
// with a ctrl-c
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	// block on waiting for a signal
	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}
	return
}

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

	/*
		abon1, err := repository.NewSelectDB(db)
		if err != nil {
			log.Fatalf("failed to SELECT: %s", err.Error())
		} else {
			fmt.Println(abon1[0])
		}


		Abon2, err := repository.NewQueryDB(db)

		if err != nil {
			log.Fatalf("failed to QUERY: %s", err.Error())
		}
	*/

	// connect the os signal to our channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// use the ws:// Scheme to connect to the websocket
	u := "ws://192.168.1.194:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// dispatch our signal catcher
	go catchSig(interrupt, c)

	err = repository.SelectAndSend(db, c)
	if err != nil {
		log.Fatal("send:", err)
	}
	//repository.Process(c, Abon2)

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
