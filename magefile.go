//+build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// Test - mage run
func Test() error {
	return sh.RunV("go", "test", "-v", "-cover", ".", "-coverprofile=coverage.out")
}

// Coverage - checking code coverage
func Coverage() error {
	if _, err := os.Stat("./coverage.out"); err != nil {
		return fmt.Errorf("run mage test befor checking the code coverage")
	}
	return sh.RunV("go", "tool", "cover", "-html=coverage.out")
}
