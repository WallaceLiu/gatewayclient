package db

import (
	"github.com/gosexy/redis"
	"log"
)

func MyRedisOpen(host string, port uint) *redis.Client {

	client := redis.New()

	err := client.Connect(host, port)

	if err != nil {
		log.Printf("Connect failed: %s\n", err.Error())
		return nil
	}

	log.Println("Connected to redis-server.")

	log.Printf("Sending PING...\n")
	s, err := client.Ping()

	if err != nil {
		log.Printf("Could not ping: %s\n", err.Error())
		return nil
	}

	log.Printf("Received %s!\n", s)
	return client
}

func MyRedisQuit(client *redis.Client) {
	client.Quit()
}

func MyRedisFlushAll(client *redis.Client) {
	client.FlushAll()
}
