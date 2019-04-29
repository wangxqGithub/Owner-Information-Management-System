package rfl

import (
	"fmt"
	"io/ioutil"
	"os"
	//"path/filepath"
)

func ListAll() ([]os.FileInfo, error) {
	//fileName := make([]string, 1)
	dir_list, err := ioutil.ReadDir("../../fileload/download")
	if err != nil {
		fmt.Println("read dir error")
		return nil, err
	}
	for i, v := range dir_list {
		//fileName = append(fileName, v.Name())
		fmt.Println(i, "=", v.Name())
	}
	return dir_list, nil
}
