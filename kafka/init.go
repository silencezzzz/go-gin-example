package kafka

import (
	"fmt"
	"github.com/Silencezzzz/go-gin-example/pkg/setting"
	"log"
)

var (
	Address string
	Topic   string
	GroupId string
)

func init() {
	sec, err := setting.Cfg.GetSection("kafka")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Address = sec.Key("KAFKA_ADDRESS").String()
	Topic = sec.Key("KAFKA_TOPIC").String()
	GroupId = sec.Key("KAFKA_GROUP_ID").String()
	fmt.Println("init kafka config success")

}
