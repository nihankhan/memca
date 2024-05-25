package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stathat/consistent"
)

func main() {
	// Create a consistent hash ring
	ch := consistent.New()

	servers := []string{
		"127.0.0.1:11211",
		"127.0.0.1:11212",
		"127.0.0.1:11213",
	}

	for _, server := range servers {
		ch.Add(server)
	}

	// Function to get the memcache client for a specific key
	getClient := func(key string) (*memcache.Client, string) {
		server, err := ch.Get(key)
		if err != nil {
			log.Fatalf("Error getting server for key: %v", err)
		}

		return memcache.New(server), server
	}

	// Add 100 keys with corresponding values
	for i := 100; i <= 500; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		client, server := getClient(key)

		err := client.Set(&memcache.Item{
			Key:   key,
			Value: []byte(value),
		})
		if err != nil {
			log.Fatalf("Error setting key %s on server %s: %v", key, server, err)
		}

		fmt.Printf("Set key '%s' with value '%s' on server: %s\n", key, value, server)
	}
}
