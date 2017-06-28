package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	MSG_TABLE     = "incoming_20160910"
	MSG_FILE_NAME = "obdbinary.dat"
)

func WriteMsgToFile(db *sql.DB) {

	var (
		processingData []byte
	)

	sql := `SELECT processingdata
		 FROM incoming_20160910 LIMIT 0,100000 ;`

	rows, err := db.Query(sql)

	fl, err := os.OpenFile(MSG_FILE_NAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fl.Close()
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		err := rows.Scan(&processingData)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(processingData[0:])

		_, werr := fl.Write(processingData[0:])
		if werr != nil {
		}
	}
	fmt.Println("End.")
}
