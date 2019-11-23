This is an example from Rob Pike's talk on Concurrency at Google Developers Conference.

The slides for the talk. 
	https://talks.golang.org/2012/concurrency.slide#1

Setting up the tone for the example. 

This example precisely serves two purposes - 
	1. Illustration of how concurrency in Go works using Channels and implementation of different Concurrency patterns.
	2. When can concurrency help and its performance benefits.

Run Commands - 
	Every file is a different implementation of the simulated search engine.

	go run google-x.go

What happens in every implementation? 

	fakesearch is a function which simulates the search in real world. It is a closure in golang what I call function with state. It takes in the kind of search we wish to perform and returns Search type which itself is a function. 

	Search expects a query and it returns a Result which is a struct with a string field. It simulates real world search by simply sleeping for a random time and then returning the result struct.

	Google func is what changes in different and improved implementations of the simulated search engine.

google-1.go
	- simply fetches results for different types of queries serially and then concats all the results.
	- Expected performance - Addition of time taken to fetch results for all types.

google-2.go
	- launches concurrent go routines to fetch results.
	- Expected performance - Maximum of time taken to fetch results of all types.

google-3.go
	- launches concurrent go routines with replications to fetch results.
	- Expected performance - Maximum of time taken to fetch results of different types wherein results of different types are calculated as minimum of all replicas.











