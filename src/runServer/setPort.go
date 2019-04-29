package runServer

import (
	"fmt"
	"html/template"
	"net/http"
	"weblog"
)

func setPort(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	//webConfig.SetWebPort("41598")
	//webConfig.WriteWebConfigXML(wdubuglog)

	//验证请求类型
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
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
	} else {
		r.ParseForm()
		//fmt.Println("port:", r.Form["port"])
		webConfig.SetWebPort(r.Form["port"][0])
		webConfig.WriteWebConfigXML(wdubuglog)
		fmt.Fprintf(w, "端口设置成功 ： %q 请重启服务使用新端口访问！\n", r.Form["port"][0])
	}
}
