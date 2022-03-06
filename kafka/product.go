package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

// SyncProducer 同步消息模式
func SyncProducer(message string) bool {
	// 配置
	config := sarama.NewConfig()
	// 属性设置
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// 创建生成者
	p, err := sarama.NewSyncProducer([]string{Address}, config)
	// 判断错误
	if err != nil {
		log.Printf("NewSyncProducer err, message=%s \n", err)
		return false
	}
	// 最后关闭生产者
	defer p.Close()
	// 主题名称
	topic := "quickstart-events"
	// 创建消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	// 发送消息
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", message, err)
		return false
	} else {
		fmt.Fprintf(os.Stdout, "数据："+message+" kafka写入成功，partition=%d, offset=%d \n", part, offset)
		return true
	}

}
