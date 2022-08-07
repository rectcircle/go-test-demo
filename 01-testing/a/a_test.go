package a

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestA1(t *testing.T) {
	fmt.Println("+++", "A1 Goroutine Num", runtime.NumGoroutine(), "A1 Pid", os.Getpid())
	fmt.Println()
}

func TestA2(t *testing.T) {
	fmt.Println("+++", "A2 Goroutine Num", runtime.NumGoroutine(), "A2 Pid", os.Getpid())
	fmt.Println()
}

func TestMain(m *testing.M) {
	fmt.Println("+++", "A TestMain Goroutine Num", runtime.NumGoroutine(), "A TestMain Pid", os.Getpid())
	os.Exit(m.Run())
}
