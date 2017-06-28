package db

import (
	"config"
	"testing"
)

func TestWriteMsgToFile(t *testing.T) {
	db := MyDBOpen(config.DB_TYPE, config.MSG_DB_URL)

	WriteMsgToFile(db)
}

//func TestReadFile(t *testing.T) {
//	b, err := ioutil.ReadFile(MSG_FILE_NAME)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(b)
//}
