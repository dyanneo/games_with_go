package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}
func main() {
	/* one way
	var p position
	p.x = 5
	p.y = 4
	*/
	// other likely better way
	p := position{4, 2}
	fmt.Println(p.x, p.y)

	b := badGuy{"Jabba the Hut", 100, p}
	//fmt.Println(b)
	whereIsBadGuy(b)
}
