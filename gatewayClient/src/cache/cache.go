package main

import (
	"cache/db"
	"config"
)

func main() {
	skDb := db.MyDBOpen(config.DB_TYPE, config.SK_DB_URL)
	vwDb := db.MyDBOpen(config.DB_TYPE, config.VW_DB_URL)

	client := db.MyRedisOpen(config.REDIS_HOST, config.REDIS_PORT)

	defer func() {
		db.MyDBClose(skDb)
		db.MyDBClose(vwDb)
		db.MyRedisQuit(client)
	}()

	db.MyRedisFlushAll(client)

	db.MySkRedis(skDb, client)
	db.MyVwRedis(vwDb, client)
}
