package b

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestB1(t *testing.T) {
	t.Parallel()
	fmt.Println("+++", "B1 Goroutine Num", runtime.NumGoroutine(), "B1 Pid", os.Getpid())
	fmt.Println()
}

func TestB2(t *testing.T) {
	t.Parallel()
	fmt.Println("+++", "B2 Goroutine Num", runtime.NumGoroutine(), "B2 Pid", os.Getpid())
	fmt.Println()
}

func TestMain(m *testing.M) {
	fmt.Println("+++", "B TestMain Goroutine Num", runtime.NumGoroutine(), "B TestMain Pid", os.Getpid())
	os.Exit(m.Run())
}
