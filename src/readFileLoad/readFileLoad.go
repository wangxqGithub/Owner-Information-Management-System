package rfl

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"path/filepath"
)

func ListAll() {
	dir_list, e := ioutil.ReadDir("../../fileload/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
	}
}
