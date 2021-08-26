package ctl

import (
	"encoding/json"
	"testing"
)

func Test_GetAllDirinfo(t *testing.T) {
	res := GetAllDirifoList()
	s, _ := json.MarshalIndent(res, "", "    ")
	println(string(s))
}
