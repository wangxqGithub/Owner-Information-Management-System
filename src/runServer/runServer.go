package runServer

import (
	"fmt"
	"log"
	"net/http"
	"weblog"
)

func Run(wlog *log.Logger) {
	port, err := GetWebConfig(wlog)
	if err != nil {
		port = "41599"
	}
	webLog.Wlog(wlog, "[info]", "读取 webConfig.xml 配置文件成功")
	http.HandleFunc("/index", index)
	webLog.Wlog(wlog, "[info]", "注册监听，启动服务")
	webLog.Wlog(wlog, "[info]", "请使用浏览器访问：http://localhost:"+port+"/index")
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
