package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	msg string
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

// use of a closure here
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Video(query)
	}()
	go func() {
		c <- Image(query)
	}()

	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("Druva")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
