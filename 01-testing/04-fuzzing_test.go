package testingdemo

import (
	"testing"
	"unicode/utf8"
)

// go test -fuzz=^FuzzReverse$ -fuzztime 2s -run ^$ ./01-testing
func FuzzReverse(f *testing.F) {
	// 1. 提供默认情况下的测试样例
	// 2. 告诉驱动器参数的类型
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) { // 2~n 个参数需要和上面 f.Add 类型一致
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
