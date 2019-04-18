package webLog

import (
	"fmt"
	"log"
	"os"
	"time"
)

//日志文件的命名格式
const logName string = "20060102" // 2006-01-02 15:04:05.000

//返回日志文件的文件夹和文件全名
func getLogPath() (string, string) {

	// 创建日志文件夹，没有则创建，有则创建失败。
	filePath := "../../webLog/"
	//fileName := filePath + time.Now().Format("2006-01-02 15:04:05.000") + ".log"
	fileName := filePath + time.Now().Format(logName) + ".log"
	return filePath, fileName
}

/*
	如果没有webLog文件夹，则创建，
	按照日志文件格式创建一个 *.log 文件
*/
func CreateLogFile() (*os.File, error) {
	//获取需要创建的日志路径与名称
	filePath, fileName := getLogPath()
	//创建一个文件夹
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	// 根据时间创建日志文件
	return os.Create(fileName)
}

// 返回 true 表示文件不存在，返回 false 表示文件存在
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

//写日志
func Wlog(debugLog *log.Logger, prefix string, logInfo string) {
	//1.	获取日志文件名，如果当前应该记录的日志文件不存在则创建一个新的，
	//	设置日志函数的记录文件指针
	//2.	如果日志文件存在则，直接将日志写入
	_, fileName := getLogPath()
	if PathExists(fileName) == false {
		fmt.Println("创建了一个新的日志文件" + fileName)
		logFile, err := CreateLogFile()
		defer logFile.Close()
		if err != nil {
			//log.Fatalln("open file error !")
			fmt.Println(err)
		}
		debugLog.SetOutput(logFile)
	}
	//logFile, err := CreateLogFile()
	//配置一个日志格式的前缀
	debugLog.SetPrefix(prefix)
	debugLog.Println(logInfo)
	//配置log的Flag参数
	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	//debugLog.Println("A different prefix")

}
