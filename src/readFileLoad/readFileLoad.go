package rfl

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"path/filepath"
)

func ListAll() ([]string, error) {
	fileName := make([]string, 10)
	dir_list, err := ioutil.ReadDir("../../fileload/download")
	if err != nil {
		fmt.Println("read dir error")
		return nil, err
	}
	for _, v := range dir_list {
		fileName = append(fileName, v.Name())
		//fmt.Println(i, "=", v.Name())
	}
	return fileName, nil
}
