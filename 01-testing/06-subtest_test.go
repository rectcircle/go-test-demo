package testingdemo

import (
	"fmt"
	"testing"
	"time"
)

// 运行指定测试
// go test -run '' ./01-testing       # 运行该包的所有测试。
// go test -run Foo ./01-testing      # 运行该包匹配 Foo 的顶级测试如 "TestFoo"。
// go test -run Foo/A= ./01-testing   # 运行该包匹配 Foo 的顶级测试，以及匹配 "A=" 的子测试。
// go test -run /A=1 ./01-testing     # 运行该包所有顶级测试，以及匹配 "A=1" 的子测试。
// go test -fuzz FuzzFoo ./01-testing # Fuzz 匹配 "FuzzFoo" 的目标。

// go test -run=FuzzFoo/9ddb952d9814  # -run 参数还可用于运行种子语料库中的特定值，以进行调试。例如：

func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) {})
	t.Run("A=2", func(t *testing.T) {})
	t.Run("B=1", func(t *testing.T) {})
	// <tear-down code>
}

func TestGroupedParallel(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{
			Name: "A=3",
		},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
		})
	}
}

func TestTeardownParallel(t *testing.T) {
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			t.Parallel()
			time.Sleep(1)
			fmt.Println("Test1")
		})
		t.Run("Test2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(1)
			fmt.Println("Test2")
		})
		t.Run("Test3", func(t *testing.T) {
			t.Parallel()
			time.Sleep(1)
			fmt.Println("Test3")
		})
	})
	// <tear-down code>
}
