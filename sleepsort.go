/* This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://sam.zoy.org/wtfpl/COPYING for more details. */
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
	ready := make(chan struct{})
	var running sync.WaitGroup
	go func(){
		for n := range in {
			running.Add(1)
			send := n
			go func(){
				<-ready
				time.Sleep(time.Duration(send) * SLEEPSORT_DURATION)
				out <- send
				running.Done()
			}()
		}
		go func() {
			running.Wait()
			close(out)
		}()
		close(ready)
	}()
	return out
}
