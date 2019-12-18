package dispatch

import (
	"killlib"
	"sync"
)

func KillFromList(cptlist []string, if_process bool, process string) []string {
	var counter sync.WaitGroup
	var state []string
	for _, each := range cptlist {
		counter.Add(1)
		go killlib.Kill(each, if_process, process, &counter, &state)
	}
	counter.Wait()
	return state
}
