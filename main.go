package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)

	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklink(link, c)
		}(l)
	}
}
func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}
