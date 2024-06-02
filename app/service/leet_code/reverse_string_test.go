package leetcode

import (
	"testing"
	"time"
)

func TestReverseString(t *testing.T) {
	println(time.Now().Nanosecond())
	reverseString([]byte("hello"))
	println(time.Now().Nanosecond())
	reverseStringBestMs([]byte("hello"))
	println(time.Now().Nanosecond())
	reverseStringBestMemory([]byte("hello"))
}
