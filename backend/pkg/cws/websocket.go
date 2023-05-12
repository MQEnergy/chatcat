package cws

import (
	"chatcat/backend/service"
	"fmt"
	gowebsocket "github.com/MQEnergy/go-websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Hub(hub *gowebsocket.Hub, a *service.App) {
	// 日志注入
	//gowebsocket.Logger = a.Log
	gowebsocket.Logger = logrus.New()
	gowebsocket.Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
	// ws连接
	http.HandleFunc("/chat", func(writer http.ResponseWriter, request *http.Request) {
		_, err := gowebsocket.WsServer(hub, writer, request, gowebsocket.Json)
		if err != nil {
			return
		}
	})

	// 推送单客户端
	http.HandleFunc("/push", func(writer http.ResponseWriter, request *http.Request) {
		clientId := request.FormValue("client_id")
		data := request.FormValue("data")
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if clientId == "" || data == "" {
			writer.Write([]byte(`{"code": -1,"msg":"parameter error"}`))
			return
		}
		hub.ClientBroadcast <- &gowebsocket.BroadcastChan{
			Name: clientId,
			Msg:  []byte(data),
		}
		writer.Write([]byte(`{"code":0,"msg":"push success"}`))
		return
	})

	// 推送到群组
	http.HandleFunc("/push_to_group", func(writer http.ResponseWriter, request *http.Request) {
		groupId := request.FormValue("group_id")
		data := request.FormValue("data")
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if groupId == "" || data == "" {
			writer.Write([]byte(`{"code": -1,"msg":"parameter error"}`))
			return
		}
		hub.GroupBroadcast <- &gowebsocket.BroadcastChan{
			Name: groupId,
			Msg:  []byte(data),
		}
		writer.Write([]byte(`{"code":0,"msg":"push group success"}`))
		return
	})
	a.LogInfof("websocket starting success port: %d", a.Cfg.App.WsPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.Cfg.App.WsPort), nil); err != nil {
		a.LogInfof("ListenAndServe err: %s", err.Error())
	}
}
