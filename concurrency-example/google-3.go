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
	Web1   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Web2   = fakeSearch("web")
	Image2 = fakeSearch("image")
	Video2 = fakeSearch("video")
)

type Search func(query string) Result

// use of a closure here
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- first(query, Web1, Web2)
	}()
	go func() {
		c <- first(query, Video1, Video2)
	}()
	go func() {
		c <- first(query, Image1, Image2)
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
