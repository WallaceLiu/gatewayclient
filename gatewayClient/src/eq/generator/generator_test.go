package generator

import (
	"common"
	"fmt"
	"testing"
)

func Test_NewSN(t *testing.T) {
	fmt.Println(NewSNNumber(common.GetRandCode()))
}

func Test_NewVINNumber(t *testing.T) {
	fmt.Println(NewVINNumber(common.GetRandCode()))
}
