package main

import "fmt"

func main() {
	feed, err := FetchFeed()
	if err != nil {
		fmt.Printf("%#v", err)
	}
	fmt.Printf("%#v", feed)
}
