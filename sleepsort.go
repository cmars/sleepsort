package sleepsort

import (
	"sync"
	"time"
)

// Base sleep duration used to sort 
const SLEEPSORT_DURATION = time.Millisecond

// Sort a slice in-place.
func SortSlice(in []int) {
	inChan := make(chan int)
	go func(){
		for _, n := range in {
			inChan <- n
		}
		close(inChan)
	}()
	outChan := SortChan(inChan)
	count := 0
	for n := range outChan {
		in[count] = n
		count++
	}
}

// Sort a channel of ints.
func SortChan(in chan int) (out chan int) {
	out = make(chan int)
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	ready := false
	var running sync.WaitGroup
	go func(){
		for n := range in {
			running.Add(1)
			send := n
			go func(){
				cond.L.Lock()
				for !ready {
					cond.Wait()
				}
				cond.L.Unlock()
				time.Sleep(time.Duration(send) * SLEEPSORT_DURATION)
				out <- send
				running.Done()
			}()
		}
		go func() {
			running.Wait()
			close(out)
		}()
		cond.L.Lock()
		ready = true
		cond.Broadcast()
		cond.L.Unlock()
	}()
	return out
}
