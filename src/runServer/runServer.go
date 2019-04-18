package runServer

import (
	"fmt"
	"log"
	"net/http"
	"weblog"
)

//定义一个log指针，用于页面访问时记录日志
var wdubuglog *log.Logger

//服务的起点
func RunStart(wlog *log.Logger) {
	wdubuglog = wlog
	run()
}

//创建服务监听
func run() {
	//获取监听端口，如果获取失败，则端口默认为 41599
	port, err := GetWebConfig(wdubuglog)
	if err != nil {
		port = "41599"
	}
	webLog.Wlog(wdubuglog, "[info]", "读取 webConfig.xml 配置文件成功")
	http.HandleFunc("/index", index)
	webLog.Wlog(wdubuglog, "[info]", "注册监听，启动服务")
	webLog.Wlog(wdubuglog, "[info]", "请使用浏览器访问：http://localhost:"+port+"/index")
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	webLog.Wlog(wdubuglog, "[info]", "访问index页面")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
