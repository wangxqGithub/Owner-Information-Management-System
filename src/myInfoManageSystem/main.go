// myInfoManageSystem project main.go
package main

import (
	"fmt"
	"log"
	"runServer"
	"weblog"
)

func main() {

	logFile, err := webLog.CreateLogFile()
	defer logFile.Close()
	if err != nil {
		//log.Fatalln("open file error !")
		fmt.Println(err)
	}
	// 创建一个日志对象,并返回
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
	webLog.Wlog(debugLog, "[info]", "服务启动")

	runServer.RunStart(debugLog)

}
