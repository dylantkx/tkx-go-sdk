package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetupTestEnv()
	code := m.Run()
	CleanUpAfterTests()
	os.Exit(code)
}
