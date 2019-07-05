package main

import "fmt"

// this is a linked list
type storyPage struct {
	text     string
	nextPage *storyPage
	prevPage *storyPage
}

/*
func (page *storyPage) playStory() {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	page.nextPage.playStory()
}
*/
func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}
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
