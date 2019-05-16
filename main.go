package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	config, _ := Config{}.load()
	links := config.Links

	if len(links) == 0 {
		fmt.Println("no 'links' found")
		os.Exit(1)
	}
	fmt.Printf("links: %v", links)

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)

	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		str := "I think link: " + l + " is down!"
		fmt.Println(str)
		c <- l
		return
	}
	str := "link: " + l + " is up!"
	fmt.Println(str)
	c <- l
}
