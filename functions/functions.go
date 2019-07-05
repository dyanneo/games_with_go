package main

import "fmt"

/*
functions have return values (go calls them functions even if no return value)
procedures don't have return values
methods are functions that are attached to ojbects like structs
*/
func sayHello(name string) {
	fmt.Println("heya", name)
}
func sayGoodbye(name string) {
	fmt.Println("seeya", name)
}
func beSocial(name string) {
	sayHello(name)
	fmt.Println("How's the weather")
	sayGoodbye(name)
}
func sayHelloALot() {
	fmt.Println("HELLO!")
	//sayHelloALot()
}
func addOne(x int) int {
	return x + 1
}
func main() {
	beSocial("dy")
	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		sayHelloALot()
	}
}
