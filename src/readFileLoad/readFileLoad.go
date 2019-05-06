package rfl

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"path/filepath"
)

func ListAll() ([]string, error) {
	fileName := make([]string, 1)
	dir_list, err := ioutil.ReadDir("../../fileload/download")
	if err != nil {
		fmt.Println("read dir error")
		return nil, err
	}
	for i, v := range dir_list {
		fileName[i] = v.Name()
		//fmt.Println(i, "=", v.Name())
	}
	return fileName, nil
}
