package db

import (
	"config"
	"log"
	"testing"
)

func Test_MyDBOpenClose(t *testing.T) {
	db := MyDBOpen(config.DB_TYPE, config.SK_DB_URL)

	defer MyDBClose(db)

	log.Println(db)
}
