package main

import "fmt"


type Person struct {

	name string
	surname string
	age int
}

func (p Person) hello(say func(what string)) {

	hello:="Hello!"
	say(hello)
}

func main() {

	p:=Person{"Marko", "Haberl", 31}
	fmt.Println(p)
	p.hello(printToScreen)
}

func printToScreen(what string) {
	fmt.Println(what)
}
