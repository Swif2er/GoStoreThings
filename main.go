package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/go-redis/redis/v8"
)

var CLI struct {
	Ping struct{} `cmd:"" help:"Ping Redis to check if the connection is working."`
	Set  struct {
		Key   string `short:"k" help:"A key to be stored."`
		Value string `short:"v" help:"A value to be stored."`
	} `cmd:"" help:"Set a key-value pair to be stored in DB."`
	Get struct {
		Key string `short:"k" help:"A key of the value to be retireved."`
	} `cmd:"" help:"Retrieve a value for a given key from DB."`
}

func main() {
	cliContext := kong.Parse(&CLI,
		kong.Name("GoStoreThings"),
		kong.Description("A tool to store and retrieve data from Redis."))

	dbClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	dbContext := dbClient.Context()

	switch cliContext.Command() {
	case "ping":
		_, err := dbClient.Ping(dbContext).Result()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Pong!")
		}
	case "set":
		err := dbClient.Set(dbContext, CLI.Set.Key, CLI.Set.Value, 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pair '%s:%s' is set.\n", CLI.Set.Key, CLI.Set.Value)
	case "get":
		value, err := dbClient.Get(dbContext, CLI.Get.Key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value for key '%s' is '%s'.\n", CLI.Get.Key, value)
	}
}
