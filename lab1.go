package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("HELLO")
	fmt.Printf("\nHello %s\n", "Valek")
	var (
		personName = "Bob"
		personAge  = 36
		personUid  int
	)
	fmt.Printf("Name: %s\nAge: %d\nUid:%d\n", personName, personAge, personUid)
	personUid = 121231241343
	fmt.Printf("UID New: %d\n", personUid)

	var a, b = 30, "Vova"
	fmt.Println(a, b)

	var (
		name string
		age  int
	)
	fmt.Print("Ваше имя? ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Ваш возраст? ")
	fmt.Fscan(os.Stdin, &age)
	//fmt.Fscan(os.Stdin, &name, &age)
	fmt.Printf("\nMy name is: %s\nMy age is:%d\n", name, age)
}
