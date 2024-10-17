package main

import (
	"fmt"

	"gosuda.org/snum/snum"
	"gosuda.org/snum/sort"
)

func main() {
	sn1 := snum.New("123456789.987654321")
	fmt.Println("snum :", sn1)

	st1 := sort.New(sn1)
	fmt.Println("sort :", st1)

	enc, _ := st1.MarshalJSON()
	fmt.Println("json :", string(enc))

	st1New := sort.New(0)
	st1New.UnmarshalJSON(enc)
	fmt.Println("new :", st1New)
}
