
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func Echo(ws *websocket.Conn)  {
	var err error
	for {
		var reply string

		// websocket 接受消息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}

		fmt.Println("reveived from client: " + reply)

		msg := "received:" + reply

		fmt.Println("send to client:" + msg)

		// 发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("send failed:", err)
			break
		}
	}
}

func web(w http.ResponseWriter, r *http.Request)  {

	// 打印请求方法
	fmt.Println("method", r.Method)

	if r.Method == "GET" { // 如果请求方法为get显示login.html,相应给前端

		t, _ := template.ParseFiles("websocket.html")
		t.Execute(w, nil)

	} else {

		// 打印输出post的参数username和password
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
	}
}

func main()  {

	// 接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))

	// html 页面
	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":1234", nil); err != nil{
		log.Fatal("ListenAndServer:", err)
	}
}