// myInfoManageSystem project main.go
package main

import (
	//"fmt"
	"log"
	"runServer"
	"weblog"
)

func main() {

	logFile, err := webLog.CreateLog()
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象,并返回
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)

	//fmt.Println("Hello World!")
	runServer.Run(debugLog)
}
