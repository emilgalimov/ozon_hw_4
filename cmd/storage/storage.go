package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/config"
	"log"
	"sync"
)

type storageManager interface {
	reserve(orderID int, productID int) bool
}

func main() {
	cfg, err := config.NewConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	consumer, err := sarama.NewConsumer(cfg.Kafka.Brokers, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions(cfg.Kafka.ConfirmOrders)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	var wg sync.WaitGroup

	fmt.Println(partitionList)
	for partition := range partitionList {
		wg.Add(1)
		pc, err := consumer.ConsumePartition(cfg.Kafka.ConfirmOrders, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			defer wg.Done()

			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
