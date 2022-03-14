package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("TestMain setup.")

	retCode := m.Run() // 执行测试，包括单元测试、性能测试和示例测试

	println("TestMain tear-down.")

	os.Exit(retCode)
}
