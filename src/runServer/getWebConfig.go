package runServer

import (
	"encoding/xml"
	"fmt"
	"os"

	//"io"
	"io/ioutil"
	"log"
	"weblog"
)

type WebConfig struct {
	XMLName xml.Name `xml:"config"`
	Port    []Port   `xml:"port"`
}

type Port struct {
	XMLName  xml.Name `xml:"port"`
	PortText string   `xml:",innerxml"`
}

func (webConfig *WebConfig) ReadWebConfigXML(wlog *log.Logger) error {
	//1. 读取配置文件
	//content, err := ioutil.ReadFile("../../../Owner-Information-Management-System/config/webConfig.xml")
	content, err := ioutil.ReadFile("../../config/webConfig.xml")
	if err != nil {
		webLog.Wlog(wlog, "[error]", "读取 webConfig.xml 配置文件失败："+fmt.Sprintf("%s", err))
		return err
	} else {
		webLog.Wlog(wlog, "[info]", "读取 webConfig.xml 配置文件成功")
	}

	//2. 解析XML内容，记录到 WebConfig 结构体的指针中
	//var result WebConfig
	err = xml.Unmarshal(content, webConfig)

	if err != nil {
		//log.Fatal(err)
		webLog.Wlog(wlog, "[error]", "解析 webConfig.xml 配置文件失败："+fmt.Sprintf("%s", err))
		return err
	} else {
		webLog.Wlog(wlog, "[info]", "解析 webConfig.xml 配置文件成功")
	}
	return nil
}

//获取端口
func (webConfig *WebConfig) GetWebPort() string {
	return webConfig.Port[0].PortText
}

//设置端口
func (webConfig *WebConfig) SetWebPort(port string) {
	webConfig.Port[0].PortText = port

}

//将调整后XML内容写入到配置文件中
func (webConfig *WebConfig) WriteWebConfigXML(wlog *log.Logger) error {

	//1. 判断 config 文件夹是否存在，如果不存在则创建一个 。
	filePath := "../../config/"
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		webLog.Wlog(wlog, "[error]", "配置文件夹【config】创建失败"+fmt.Sprintf("%s", err))
	} else {
		webLog.Wlog(wlog, "[info]", "配置文件夹【config】创建成功")
	}

	//2. 判断 webConfig.xml 文件是否存在，如果不存在则创建。
	fileName := "webConfig.xml"
	/*
		_, err = os.Create(filePath + fileName)
		if err != nil {
			fmt.Println(err)
			webLog.Wlog(wlog, "[error]", "配置文件【webConfig.xml】创建失败")
		} else {
			webLog.Wlog(wlog, "[info]", "配置文件【webConfig.xml】创建成功")
		}
	*/

	//3. 保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(webConfig, "	", "	")
	//xmlOutPut, outPutErr := xml.Marshal(webConfig, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(xml.Header)
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		//写入文件
		ioutil.WriteFile(filePath+fileName, xmlOutPutData, os.ModeAppend)
		webLog.Wlog(wlog, "[info]", "修改文件【webConfig.xml】成功")
	} else {
		//fmt.Println(outPutErr)
		webLog.Wlog(wlog, "[info]", "修改文件【webConfig.xml】失败："+fmt.Sprintf("%s", outPutErr))
		return outPutErr
	}
	return nil
}
