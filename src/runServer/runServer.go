package runServer

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	port, err := GetWebConfig()
	if err != nil {
		port = "41599"
	}
	log.Println("读取 webConfig.xml 配置文件成功")
	http.HandleFunc("/index", index)
	log.Println("注册监听，启动服务")
	log.Println("请使用浏览器访问：http://localhost:" + port + "/index")
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
