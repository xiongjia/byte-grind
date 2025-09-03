package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadgerData(t *testing.T) {
	err := BadgerData()
	assert.Error(t, err)
}
