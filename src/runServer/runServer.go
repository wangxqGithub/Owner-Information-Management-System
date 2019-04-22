package runServer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"weblog"
)

var wdubuglog *log.Logger //定义一个log指针，用于页面访问时记录日志
var port string           //记录 webconfig.xml 配置文件中的 端口号
var webConfig WebConfig   //记录 webconfig.xml 配置文件的全部内容

//服务的起点
func RunStart(wlog *log.Logger) {
	wdubuglog = wlog
	run()
}

//创建服务监听
func run() {
	//1. 获取端口号，如果配置文件读取失败则，默认端口 41599
	//webConfig = WebConfig{}
	err := webConfig.ReadWebConfigXML(wdubuglog)
	//获取监听端口，如果获取失败，则端口默认为 41599
	//port, err := GetWebConfig(wdubuglog)
	if err != nil {
		port = "41599"
		webLog.Wlog(wdubuglog, "[error]", "读取 webConfig.xml 配置文件失败"+fmt.Sprintf("%s", err))
	} else {
		port = webConfig.GetWebPort()
		webLog.Wlog(wdubuglog, "[info]", "读取 webConfig.xml 配置文件成功")
	}

	//2. 注册页面
	handleFuns()

	//3. 注册监听
	webLog.Wlog(wdubuglog, "[info]", "注册监听，启动服务")
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func handleFuns() {
	//注册测试页面
	http.HandleFunc("/test", test)
	webLog.Wlog(wdubuglog, "[info]", "测试页面注册成功，请使用浏览器访问：http://localhost:"+port+"/test")
	//注册设置端口页面
	http.HandleFunc("/setPort", setPort)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func setPort(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	//webConfig.SetWebPort("41598")
	//webConfig.WriteWebConfigXML(wdubuglog)
	t, err := template.ParseFiles("../../view/setPort.html", "../../view/template/header.html")
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【setPort】加载失败"+fmt.Sprintf("%s", err))
	}

	data := struct {
		Title string
	}{
		Title: "设置端口",
	}

	err = t.Execute(w, data)
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【setPort】打开失败"+fmt.Sprintf("%s", err))
	}
}
