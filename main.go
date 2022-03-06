package main

import (
	"fmt"
	"github.com/Silencezzzz/go-gin-example/kafka"
	"github.com/Silencezzzz/go-gin-example/pkg/setting"
	"github.com/Silencezzzz/go-gin-example/routers"
	"net/http"
)

//
func main() {
	//订阅两个分区 仅仅取数据
	go kafka.ConsumerGroup("消费者1")
	go kafka.ConsumerGroup("消费者2")
	//启动指定数量的业务消费者 实际处理处理业务逻辑
	go kafka.WorkerInit(10)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
