package main

import "fmt"

// this is a doubly linked list (each page has a next and prev)
type storyPage struct {
	text     string
	nextPage *storyPage
	//prevPage *storyPage
}

// receiver syntax for putting methods on structs
// go does not do tail call elimination, which means the stack
// will keep growing, rather than turn the last statement into a loop
// automatically, the way some compilers will do
func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage

	}
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// function = has a return value
// procedure = executes commands with no return value
// method = functions attached to a struct/object

func main() {
	page1 := storyPage{"You are standing in an open field west of a white house.", nil}
	page1.addToEnd("You climb into the attic, it is pitch black, you can't see a thing")
	page1.addToEnd("You are eaten by a Grue.")

	page1.addAfter("testing add after")
	page1.playStory() //may be refered to as a method call

}

// func addPage(currentPage *storyPage, newPage *storyPage) {
// 	followingPage := currentPage.nextPage
// 	currentPage.nextPage = newPage
// 	newPage.nextPage = followingPage
// 	newPage.prevPage = currentPage
// 	if newPage.nextPage != nil {
// 		newPage.nextPage.prevPage = newPage
// 	}
// }

// func deletePage() {

// }
