package runServer

import (
	"fmt"
	"html/template"
	"net/http"
	"weblog"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../view/index.html", "../../view/template/header.html")
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【index】加载失败"+fmt.Sprintf("%s", err))
	}

	data := struct {
		Title string
	}{
		Title: "文件传输服务",
	}

	err = t.Execute(w, data)
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【index】打开失败"+fmt.Sprintf("%s", err))
	}
}
