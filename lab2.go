package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	var firstValue int = 10
	var secValue int64 = 123
	fmt.Printf("\nfirstValue тип %T и кол-во байт %d\n", firstValue, unsafe.Sizeof(firstValue))
	sum := int64(firstValue) + secValue
	fmt.Println("Summa:", sum)

	var (
		floatFirst float32 = 12.33
		floatSec   float64 = 14.55
	)
	sub := float64(floatFirst) + floatSec
	fmt.Println("sum float: ", sub)
	fmt.Printf("\nsum: %.3f\n", sub)

	totalString, secString := "qwerty", "qwerty"
	fmt.Println("Сравниваем строку: ", strings.Compare(totalString, secString))

	str := "Hello, Валек"
	fmt.Println(len(str))
	// Русские символы занимают по два байта
	fmt.Println(utf8.RuneCountInString(str))

	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("heingt <= 100")
}
