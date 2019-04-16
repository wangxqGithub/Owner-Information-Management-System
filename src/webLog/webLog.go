package webLog

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	如果没有webLog文件夹，则创建，
	按照日期创建一个 *.log 文件
*/
func CreateLog() (*os.File, error) {

	// 创建日志文件夹
	filePath := "../../webLog/"
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// 根据时间创建日志文件
	//fileName := filePath + time.Now().Format("2006-01-02 15") + ".log"
	fileName := filePath + time.Now().Format("20060102") + ".log"
	//如果是 false 说明文件不不存在
	if PathExists(fileName) == false {
		return os.Create(fileName)
	} else {
		fileName = filePath + time.Now().Format("20060102150405") + ".log"
		return os.Create(fileName)
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Wlog(debugLog *log.Logger, prefix string, logInfo string) {

	//debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	debugLog.SetPrefix(prefix)
	debugLog.Println(logInfo)
	//配置log的Flag参数
	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	//debugLog.Println("A different prefix")

}
