package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsible(t *testing.T) {
	err := AnsibleExec()
	assert.Error(t, err)
}
