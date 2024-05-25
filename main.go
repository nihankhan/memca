package main

import (
	"fmt"
	"log"

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

	// Function to get the memcache client for  specific key
	getClient := func(key string) (*memcache.Client, string) {
		server, err := ch.Get(key)
		if err != nil {
			log.Fatalf("Error getting server for key: ", err)
		}

		return memcache.New(server), server
	}

	// // Example: Set a key-value pair
	// key := "key250"
	// value := "bar"
	// client, server := getClient(key)
	// err := client.Set(&memcache.Item{
	// 	Key: key,
	// 	Value: []byte(value),
	// })

	// if err != nil {
	// 	log.Fatalf("Error setting key: ", err)
	// }

	fmt.Printf("Retrieves value of the '%s' from the server: %s\n", key, server)

	// Example: Get the value for the key
	client, _ = getClient(key)
	item, err := client.Get(key)
	if err != nil {
		log.Fatalf("Error getting key:", err)
	}

	// Print the retrieved value
	fmt.Printf("The value of '%s' is: %s\n", key, item.Value)
}