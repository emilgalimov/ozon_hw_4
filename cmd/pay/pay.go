package main

import (
	"context"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/config"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/db"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/payService"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/saga"
	"log"
	"sync"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	saramaConfig := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, "pay", saramaConfig)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	ctx := context.Background()

	repo := db.NewPayRepo()

	pay := payService.NewPayService(repo)

	saramaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Kafka.Brokers, saramaConfig)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	consumer := saga.NewOrderPayer(
		pay,
		syncProducer,
		cfg.Kafka.Delivered,
		cfg.Kafka.Rejected,
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			if err := client.Consume(ctx, []string{cfg.Kafka.WriteOff}, consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()

}
