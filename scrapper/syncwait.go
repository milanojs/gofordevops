package main

import (
	"sync"
	"time"
)

var sites []string = []string{
	"http://example.com/site1.html",
	"http://example.com/site2.html",
	"http://example.com/site3.html",
}
var activeThreads = 0
var doneCount = 0

const maxActiveThreads = 1

func scrapeSite(site string, condition *sync.Cond) {
	condition.L.Lock()
	if activeThreads >= maxActiveThreads {
		condition.Wait()
	}
	activeThreads++
	condition.L.Unlock()
	println("scraping " + site)
	// Scraping code goes here ...
	condition.L.Lock()

	activeThreads--
	doneCount++
	condition.L.Unlock()
	condition.Signal()
}

func main() {
	var l = sync.Mutex{}
	var c = sync.NewCond(&l)

	for _, site := range sites {
		println("starting scraper for " + site)
		go scrapeSite(site, c)
	}
	for doneCount < len(sites) {
		time.Sleep(2 * time.Second)
	}
	println("Done!")
}
