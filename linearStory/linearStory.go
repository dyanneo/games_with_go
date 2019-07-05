package main

import "fmt"

// this is a doubly linked list (each page has a next and prev)
type storyPage struct {
	text     string
	nextPage *storyPage
	prevPage *storyPage
}

// receiver syntax for putting methods on structs
// go does not do tail call elimination, which means the stack
// will keep growing, rather than turn the last statement into a loop
// automatically, the way some compilers will do
func (page *storyPage) playStory() {
	if page != nil {
		fmt.Println(page.text)
		page = page.nextPage

	}
}

// function = has a return value
// procedure = executes commands with no return value
// method = functions attached to a struct/object

func main() {
	page1 := storyPage{"It was a dark and stormy night", nil, nil}
	page2 := storyPage{"You are alone and have to find the sacred helmet before the bad guys do", nil, &page1}
	page3 := storyPage{"You see a troll ahead", nil, &page2}
	page1.nextPage = &page2
	page2.nextPage = &page3

	page1.playStory() //may be refered to as a method call

}

func addPage(currentPage *storyPage, newPage *storyPage) {
	followingPage := currentPage.nextPage
	currentPage.nextPage = newPage
	newPage.nextPage = followingPage
	newPage.prevPage = currentPage
	if newPage.nextPage != nil {
		newPage.nextPage.prevPage = newPage
	}
}

func deletePage() {

}
