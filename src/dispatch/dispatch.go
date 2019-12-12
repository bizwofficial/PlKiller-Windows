package dispatch

import (
	"killlib"
	"sync"
)

func KillFromList(cptlist []string) []string {
	var counter sync.WaitGroup
	var state []string
	for _,each:=range cptlist{
		counter.Add(1)
		go killlib.Kill(each,false,"default",&counter,&state)
	}
	counter.Wait()
	return state
}