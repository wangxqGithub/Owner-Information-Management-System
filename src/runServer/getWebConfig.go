package runServer

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type WebConfig struct {
	XMLName xml.Name `xml:"config"`
	Port    []Port   `xml:"port"`
}

type Port struct {
	XMLName  xml.Name `xml:"port"`
	PortText string   `xml:",innerxml"`
}

func GetWebConfig() (string, error) {
	//content, err := ioutil.ReadFile("../../../Owner-Information-Management-System/config/webConfig.xml")
	content, err := ioutil.ReadFile("../../config/webConfig.xml")
	log.Println("读取 webConfig.xml 配置文件")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var result WebConfig
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//log.Println(result)
	log.Println("服务端口:", result.Port[0].PortText)
	return result.Port[0].PortText, nil
	/*
		for _, o := range result.Port {
			log.Println(o.PortText + "===" + o.PortText)
		}*/

}
