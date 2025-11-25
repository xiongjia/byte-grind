package util

import (
	"fmt"
	"log/slog"
	"net/url"
	"path"
	"testing"
	"time"

	"github.com/samber/lo"
)

func TestDo(t *testing.T) {
	numbers := []int{1, 2, 2, 3, 1, 4}
	chunks := lo.Chunk(numbers, 5)
	for i, chunk := range chunks {
		t.Logf("chunk %d: %v", i, chunk)
		for j, v := range chunk {
			t.Logf("chunks[%d][%d] = %d", i, j, v)
		}
	}

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

	fmt.Append()
	time.Sleep(1000 * 10)
}

func TestConcurrentQuery(t *testing.T) {
	c := NewConcurrentQuery()

	result := c.Query([]string{"a", "b", "c"})
	for k, v := range result {
		fmt.Printf("k = %s, v = %s\n", k, v)
		t.Logf("k = %s, v = %s\n", k, v)
	}

	slog.Debug("TestConcurrentQuery completed")

	baseURL := "https://api.example.com/v1"
	endpoints := []string{"a/b", "c", "abc"}
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("URL 解析失败: %v\n", err)
		return
	}

	fullPath := path.Join(append([]string{u.Path}, endpoints...)...)
	u.Path = fullPath
	fmt.Printf("%s\n", u.String())

}
