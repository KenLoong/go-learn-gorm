package main

import (
	"learn-gorm/util"
)

func init() {
	err := util.Connect("warson_test", "root", "dc")
	if err != nil {
		return
	}
}

func main() {
	util.Create2()
}
