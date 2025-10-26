package util

import (
	"testing"

	_ "github.com/alicebob/miniredis/v2"
)

func TestRedisMgr_WithMiniredis(t *testing.T) {
	RedisMgr()
}
