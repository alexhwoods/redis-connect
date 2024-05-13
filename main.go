package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	client := ExampleClient()

	result, err := client.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
