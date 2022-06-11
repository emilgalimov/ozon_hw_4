package main

import (
	"context"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/config"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/db"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/saga"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/storageService"
	"log"
	"sync"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	saramaConfig := sarama.NewConfig()
	client1, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, "storage", saramaConfig)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	client2, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, "storage", saramaConfig)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	ctx := context.Background()

	repo := db.NewStorageRepo()

	stor := storageService.NewStorage(repo)

	saramaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Kafka.Brokers, saramaConfig)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	consumer := saga.NewStorageReserver(stor, syncProducer, cfg.Kafka.WriteOff)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client1.Consume(ctx, []string{cfg.Kafka.ConfirmOrders}, consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	consumer2 := saga.NewStorageUnreserver(stor, syncProducer)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client2.Consume(ctx, []string{cfg.Kafka.Rejected}, consumer2); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}
