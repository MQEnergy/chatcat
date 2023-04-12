package cws

import (
	"chatcat/backend/service"
	gowebsocket "github.com/MQEnergy/go-websocket"
	"net/http"
)

func Hub(hub *gowebsocket.Hub, a *service.App) {
	// 日志注入
	gowebsocket.Logger = a.Log

	// ws连接
	http.HandleFunc("/chat", func(writer http.ResponseWriter, request *http.Request) {
		client, err := gowebsocket.WsServer(hub, writer, request, gowebsocket.Json)
		if err != nil {
			return
		}
		a.ClientId = client.ClientId
	})

	// 推送
	http.HandleFunc("/push", func(writer http.ResponseWriter, request *http.Request) {
		clientId := request.FormValue("client_id")
		data := request.FormValue("data")
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if clientId == "" || data == "" {
			writer.Write([]byte(`{"code": -1,"msg":"参数错误"}`))
			return
		}
		hub.ClientBroadcast <- &gowebsocket.BroadcastChan{
			Name: clientId,
			Msg:  []byte(data),
		}
		a.LogInfof("客户端消息发送成功 client_id: %s msg: %s", clientId, data)
		writer.Write([]byte(`{"code":0,"msg":"客户端消息发送成功"}`))
		return
	})
	a.LogInfo("websocket starting success port: 9991")

	go func() {
		if err := http.ListenAndServe(":9991", nil); err != nil {
			a.LogInfof("ListenAndServe err:", err.Error())
		}
	}()
}
