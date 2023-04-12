package cws

import (
	"chatcat/backend/service"
	gowebsocket "github.com/MQEnergy/go-websocket"
	"log"
	"net/http"
)

func Hub(hub *gowebsocket.Hub, a *service.App) {
	// 日志注入
	gowebsocket.Logger = a.Log

	// ws连接
	http.HandleFunc("/chat", func(writer http.ResponseWriter, request *http.Request) {
		_, err := gowebsocket.WsServer(hub, writer, request, gowebsocket.Json)
		if err != nil {
			return
		}
	})

	// 推送到单个客户端
	http.HandleFunc("/push_to_client", func(writer http.ResponseWriter, request *http.Request) {
		clientId := request.FormValue("client_id")
		data := request.FormValue("data")
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if clientId == "" || data == "" {
			writer.Write([]byte("{\"msg\":\"参数错误\"}"))
			return
		}
		hub.ClientBroadcast <- &gowebsocket.BroadcastChan{
			Name: clientId,
			Msg:  []byte(data),
		}
		writer.Write([]byte(`{"msg":"客户端消息发送成功"}`))
		return
	})
	a.LogInfo("ws服务启动成功。端口号 :9991")
	go func() {
		if err := http.ListenAndServe(":9991", nil); err != nil {
			log.Println("ListenAndServe: ", err)
		}
	}()
}
