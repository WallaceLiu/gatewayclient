package common

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_Recognition(t *testing.T) {
	bArr := []byte{64, 64} // 新科 40 40

	fmt.Print(bArr)
}

func TestCatTsinghuabinary(t *testing.T) {

	file, err := os.Open("../resolve/testfile/tsinghuabinary.dat")
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	fmt.Println(file)

	payload := Cat(file)

	fmt.Println(len(payload))
}
