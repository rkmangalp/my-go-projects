package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	// Define the URL to scrape
	url := "https://reliasoftware.com/blog/golang-project-ideas"

	// Create a new Colly collector
	c := colly.NewCollector(
		// Enable debug mode to see more information about the scraping process
		colly.Debugger(&debug.LogDebugger{}),
	)

	// On request, print a message indicating the URL being visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// On response, extract and print the title of the page
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("Title:", title)
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
