package main

import (
	"avalanche_indexer/collector"
	"sync"
)

const (
	wgCount = 2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(wgCount)

	go collector.Collect()
	wg.Wait()
}
