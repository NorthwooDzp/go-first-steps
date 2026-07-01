package runtimex

import (
	"strconv"
	"sync"
	"time"
)

func GoRoutinesSimple() {
	// var wg sync.WaitGroup
	wg := new(sync.WaitGroup)

	runWithDelay := func(wg *sync.WaitGroup, index int, delay int) {
		time.Sleep(time.Duration(delay) * time.Second)
		println("Execution #" + strconv.Itoa(index) + " is executed with delay " + strconv.Itoa(delay) + " seconds")
		wg.Done()
	}

	wg.Add(1)
	go runWithDelay(wg, 1, 3) // run function in a new goroutine; wg tracks completion
	wg.Add(1)
	go runWithDelay(wg, 2, 4) // run function in a new goroutine; wg tracks completion
	wg.Add(1)
	go runWithDelay(wg, 3, 2) // run function in a new goroutine; wg tracks completion
	wg.Wait()                 // wait until all goroutines have finished
	println("Execution line below postponed func")
}

func GoRoutinesWithChannels() {
	addWithDelay := func(a, b int, ch chan int) {
		/*
			Channel argument can be directional:
			<-chan type - receive-only channel
			chan<- type - send-only channel
			chan type   - bidirectional channel (receive and send)
		*/
		time.Sleep(time.Duration(1) * time.Second)
		ch <- a + b // without sending a value to the channel, the goroutine would block and may deadlock
	}

	ch := make(chan int)

	go addWithDelay(16, 32, ch)
	res := <-ch // receive blocks until a value is sent on the channel
	println(res)

	go addWithDelay(5, 7, ch)
	go addWithDelay(5, 16, ch)
	go addWithDelay(9, 7, ch)
	go addWithDelay(14, 6, ch)

	res = <-ch // receive values from channel one by one, blocking until each send completes
	println(res)
	res = <-ch
	println(res)
	res = <-ch
	println(res)
	res = <-ch
	println(res)

	// res = <-ch // redundant blocking code; no more sends will happen and this would deadlock
	// println(res)

	ch2 := make(chan int)
	ch3 := make(chan int)

	go addWithDelay(15, 34, ch2)
	go addWithDelay(44, 36, ch3)

	res2 := <-ch2
	res3 := <-ch3

	println("channels 2 and 3 values =>", res2, "|", res3)

	ch4 := make(chan int)
	ch5 := make(chan int)

	go addWithDelay(54, 89, ch4)
	go addWithDelay(89, 88, ch5)

	select { // select chooses the first channel that is ready for receive
	case res4 := <-ch4:
		println("Result 4 is =>", res4)
	case res5 := <-ch5:
		println("Result 5 is =>", res5)
	}

	/*
		Deadlock can occur with an unbuffered channel if a send is attempted before there is a receiver.
		The send blocks until another goroutine receives the value, so both sides must synchronize.
	*/
	/*
		ch6 := make(chan bool)
		ch6 <- true
		<-ch6
		println("Done with ch6")
	*/

	buffCh := make(chan bool, 2) // buffered channel with capacity 2
	buffCh <- true
	<-buffCh
	println("Done with buffered channel")

}

type Counter struct {
	value int
	lock  sync.Locker
}

func (c *Counter) increase() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
	println("Current counter value =>", c.value)
}

func GoRoutinesWithLock() {
	counter := Counter{0, new(sync.Mutex)}
	wg := new(sync.WaitGroup)

	for range 200 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increase()
		}()
		// Simpler record for previous code

		wg.Go(func() {
			counter.increase()
		})

	}

	wg.Wait()

}
