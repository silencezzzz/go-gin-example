package kafka

import (
	"fmt"
	"sync"
	"time"
)

var (
	MsgCh = make(chan string, 10)
)

func WorkerInit(goWorksNums int) {
	//锁
	var wg sync.WaitGroup
	for i := 0; i < goWorksNums; i++ {
		num := i
		wg.Add(1)
		go func() {
			fmt.Printf("启动业务消费者:%d \n", num)
			defer wg.Done()
			for msg := range MsgCh {

				//实际业务处理
				fmt.Printf("go func: %d, time: %d , msg : %s\n", num, time.Now().Unix(), msg)
				//模拟处理时间
				time.Sleep(time.Second * 5)
			}
		}()
	}
	defer close(MsgCh)
	wg.Wait()
}
