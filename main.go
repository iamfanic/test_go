package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// 连接到NATS服务器
	nc, err := nats.Connect("nats://150.158.33.69:14222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 订阅消息
	//sub, err := nc.Subscribe("subject", func(msg *nats.Msg) {
	//	log.Printf("Received message!")
	//	a := msg.Data
	//	var res *packets.PublishPacket
	//	json.Unmarshal(a, &res)
	//	fmt.Println(string(res.Payload))
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer sub.Unsubscribe()

	// 发布消息
	//err = nc.Publish("subject", []byte("Hello NATS!"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 等待消息处理
	//nc.Flush()
	//if err := nc.LastError(); err != nil {
	//	log.Fatal(err)
	//}

	// 持续订阅消息，直到程序退出
	//nc.Subscribe("subject", func(msg *nats.Msg) {
	//	log.Printf("Received message")
	//})

	nc.Subscribe("aaa", func(msg *nats.Msg) {
		fmt.Println("got it")
		msg.Respond([]byte("ni hao1"))
	})

	fmt.Println("cnm")
	// 保持主程序运行一段时间
	time.Sleep(30 * time.Second)
}
