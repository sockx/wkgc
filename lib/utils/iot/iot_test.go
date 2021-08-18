package iot

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	var dir = "D:\\code\\wkgc"
	fmt.Print(ScanDir(dir))
}

func TestIsGit(t *testing.T) {
	var dir = "D:\\code\\wkgc"
	fmt.Print(IsGit(dir))
}
