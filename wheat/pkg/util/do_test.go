package util

import (
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	SamberDo()
}

func TestAddWg(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("Worker %d: Working...\n", i)
			time.Sleep(200 * time.Millisecond) // pretend to do work
			fmt.Printf("Worker %d: Finished.\n", i)
		}()
	}

	time.Sleep(1000 * 10)
}
