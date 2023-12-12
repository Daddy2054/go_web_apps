package main

import "fmt"

func main() {
	fmt.Println("Hello, world")

	var whatToSay string
	var i int

	whatToSay = "Goodbay, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is seto to", i)
	whatWasSaid, theOtherThing := saySomething()

	fmt.Println("The function returned", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "something", "else"
}
