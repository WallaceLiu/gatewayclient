package db

import (
	"config"
	"log"
	"testing"
)

func Test_MyRedisOpenClose(t *testing.T) {
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer MyRedisQuit(rd)

	log.Println(rd)
}

func Test_hmset(t *testing.T) {
	rd := MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer MyRedisQuit(rd)

	s, err := rd.HMSet("user:1", "name", "liuning", "age", 30)

	log.Println(s)
	log.Println(err)
}
