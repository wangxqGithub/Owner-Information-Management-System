package runServer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"strings"
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
	//handleFuns()
	http.HandleFunc("/", handleFuns)
	//3. 注册监听
	webLog.Wlog(wdubuglog, "[info]", "注册监听，启动服务")
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func handleFuns(w http.ResponseWriter, r *http.Request) {
	//注册测试页面
	switch r.URL.Path {
	case "/":
		//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		index(w, r)
	case "/setPort":
		setPort(w, r)
	default:
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
}

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

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// 解析url传递的参数
	r.ParseForm()
	/*
		fmt.Println(r.Form)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			// join() 方法用于把数组中的所有元素放入一个字符串。
			// 元素是通过指定的分隔符进行分隔的
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/
	// 输出到客户端
	port := r.Form["port"]
	for _, v := range port {
		//fmt.Println(i)
		fmt.Fprintf(w, "port:%v\n", v)
	}

}

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
