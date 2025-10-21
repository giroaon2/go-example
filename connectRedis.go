package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a new Redis client
	// For a local Redis instance, the address is typically "localhost:6379"
	// Adjust the password and DB as needed for your setup
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Default DB
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Printf("Connected to Redis: %s\n", pong)

	// Example: Set and Get a key-value pair
	err = client.Set(ctx, "mykey", "myvalue", 0).Err()
	if err != nil {
		log.Fatalf("Error setting key: %v", err)
	}

	value, err := client.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatalf("Error getting key: %v", err)
	}
	fmt.Printf("Value of mykey: %s\n", value)

	// Close the client connection when done
	defer client.Close()
}
