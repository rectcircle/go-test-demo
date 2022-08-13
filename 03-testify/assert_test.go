package testifydemo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	// 相等断言
	assert.Equal(t, 123, 124, "they should be equal")

	// 不等断言
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// nil 断言
	assert.Nil(t, nil)

	// 非 nil 断言
	got2 := "Something"
	if assert.NotNil(t, got2) {
		// 现在 got2 不是 nil
		// 可以安全地进行进一步的断言而不会导致任何错误
		assert.Equal(t, "Something", got2)
	}
	t.FailNow()
}
