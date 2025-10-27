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

func TestConcurrentQuery(t *testing.T) {
	c := NewConcurrentQuery()

	result := c.Query([]string{"a", "b", "c"})
	for k, v := range result {
		fmt.Printf("k = %s, v = %s\n", k, v)
		t.Logf("k = %s, v = %s\n", k, v)
	}
}
