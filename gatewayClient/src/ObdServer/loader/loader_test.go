package loader

import (
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	log.Println(config)
}
