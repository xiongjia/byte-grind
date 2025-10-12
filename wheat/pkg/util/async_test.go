package util

import (
	"fmt"
	"testing"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Producer: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func TestRange(t *testing.T) {
	dataChan := make(chan int)
	go producer(dataChan)
	for v := range dataChan {
		t.Logf("Consumer: %d\n", v)
	}
	t.Logf("test finished\n")

	dataChan2 := make(chan int)
	go func() {
		dataChan2 <- 1
		close(dataChan2)
	}()

	r, ok := <-dataChan2
	if ok {
		fmt.Printf("r1 =  %d\n", r)
	}
	r, ok = <-dataChan2
	if ok {
		fmt.Printf("r2 =  %d\n", r)
	}
}
