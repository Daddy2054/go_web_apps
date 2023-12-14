package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		x, y float32
		want float32
		err  error
	}{
		{100, 2, 50, nil},
		{100, 0, 0, errors.New("cannot divide by 0")},
	}

	for _, test := range tests {
		got, err := divide(test.x, test.y)
		assert.Equal(t, test.want, got)
		assert.Equal(t, test.err, err)
	}
}
