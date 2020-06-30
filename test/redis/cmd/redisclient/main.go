package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test Redis connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("ping to redis failed: %s", err)
	}
	log.Println(pong)

	// Set key-value
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err = client.Set(ctx, "name", "Elliot", 0).Err(); err != nil {
		log.Fatal(err)
	}

	// Read value from key
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(val)

	type Author struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Fill JSON with data
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		log.Fatal(err)
	}

	// Save composite value
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	if err = client.Set(ctx, "id1234", json, 0).Err(); err != nil {
		log.Fatal(err)
	}

	// Read composite value
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	val, err = client.Get(ctx, "id1234").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(val)
}
