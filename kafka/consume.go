package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func InitConsumer() {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{Address}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	fmt.Printf("当前分区数:%d 创建指定数目消费进程\n", len(partitionList))

	for partition := range partitionList { // 遍历所有的分区
		wg.Add(1)
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {

			for msg := range pc.Messages() {
				fmt.Printf("消费成功：当前分区：:%d Offset:%d Key:%v 值：:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	fmt.Println("init Consumer success")
	wg.Wait()
	_ = consumer.Close()
}
