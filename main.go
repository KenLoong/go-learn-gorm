package main

import (
	"learn-gorm/util"
)

func main() {
	err := util.Connect("warson_test", "root", "dc")
	if err != nil {
		return
	}
	//学习更进一步
	//master:1a
	//master:2b

}
