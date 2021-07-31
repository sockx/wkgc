package iot

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	var dir = "D:\\code\\wkgc"
	fmt.Print(scanDir(dir))
}

func TestIsGit(t *testing.T) {
	var dir = "D:\\code\\wkgc"
	fmt.Print(isGit(dir))
}
