package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// 连接到NATS服务器
	nc, err := nats.Connect("nats://150.158.33.69:14222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.PublishRequest("aaa", "bbb", []byte("haha"))

	sub, _ := nc.SubscribeSync("bbb")

	msg, err := sub.NextMsg(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(msg.Data))

	//packet := packets.NewControlPacket(packets.Publish).(*packets.PublishPacket)
	//packet.TopicName = "i am nat-puber"
	//packet.Payload = []byte("haha")
	//
	//data, _ := json.Marshal(packet)
	//
	//// 发布消息
	//err = nc.Publish("subject", data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 等待消息处理
	//nc.Flush()
	//if err := nc.LastError(); err != nil {
	//	log.Fatal(err)
	//}

	// 保持主程序运行一段时间
	//time.Sleep(5 * time.Second)
}
