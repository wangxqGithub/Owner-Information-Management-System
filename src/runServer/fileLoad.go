package runServer

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"readFileLoad"
	"weblog"
)

type FileData struct {
	Title string
	Items []os.FileInfo
}

func fileLoad(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../view/fileload.html", "../../view/template/header.html")
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【fileload】加载失败"+fmt.Sprintf("%s", err))
	}
	var data FileData
	data.Title = "文件下载服务"
	//fmt.Println(data.Title)
	data.Items, err = rfl.ListAll()
	fmt.Println(data.Items)
	if err != nil {
		fmt.Println("read dir error")
	}
	err = t.Execute(w, data)
	if err != nil {
		webLog.Wlog(wdubuglog, "[error]", "页面【fileload】打开失败"+fmt.Sprintf("%s", err))
	}
}
