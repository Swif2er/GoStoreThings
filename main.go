package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/go-redis/redis/v8"
)

var CLI struct {
	Ping struct{} `cmd:"" help:"Ping Redis to check if the connection is working."`
	Set  struct {
		Key   string `short:"k" help:"A key to be stored. "`
		Value string `short:"v" help:"A value to be stored. "`
	} `cmd:"" help:"Set a key-value pair to be stored in DB. [-k <key-name> -v <value>]"`
	Get struct {
		Key string `short:"k" help:"A key of the value to be retireved."`
	} `cmd:"" help:"Retrieve a value for a given key from DB. [-k <key-name>]"`
}

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	cliContext := kong.Parse(&CLI,
		kong.Name("go-store-things"),
		kong.Description("CLI to store and retrieve data from Redis."))

	var (
		host     = getEnv("REDIS_HOST", "localhost")
		port     = string(getEnv("REDIS_PORT", "6379"))
		password = getEnv("REDIS_PASSWORD", "")
	)

	dbClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
