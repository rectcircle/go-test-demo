package testingdemo

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("+++ TestMain +++")
	os.Exit(m.Run())
}
