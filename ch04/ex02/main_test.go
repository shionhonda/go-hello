package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS256(t *testing.T) {
	assert.Len(t, s256("hoge"), 2*32)
}

func TestS384(t *testing.T) {
	assert.Len(t, s384("fugafuga"), 2*48)
}

func TestS512(t *testing.T) {
	assert.Len(t, s512("piyopiyopiyo"), 2*64)
}
