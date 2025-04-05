package main

import (
	"github.com/z46-dev/go-logger"
	"github.com/z46-dev/gotest"
)

func main() {
	bruh := gotest.NewTestSuite("Test Suite 1").SetMeta("Testing stuff")
	bruh.AddGroups(gotest.NewTestGroup("TG 1").SetMeta("Bruh", false).AddTests(gotest.NewTest("Test 1", func(log *logger.Logger) (passFail bool) {
		return true
	})))

	bruh.Run()

	bruh = gotest.NewTestSuite("Test Suite 2").SetMeta("Testing stuff")
	bruh.AddGroups(gotest.NewTestGroup("TG 1").SetMeta("Bruh", false).AddTests(gotest.NewTest("Test 1", func(log *logger.Logger) (passFail bool) {
		return false
	})))

	bruh.Run()
}
