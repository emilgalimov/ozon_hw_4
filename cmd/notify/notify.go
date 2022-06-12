package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/config"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/db"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/notifyService"
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
	client, err := sarama.NewConsumerGroup(cfg.Kafka.Brokers, "notify", saramaConfig)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	ctx := context.Background()

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Notifications.Db.User,
		cfg.Notifications.Db.Password,
		cfg.Notifications.Db.Host,
		cfg.Notifications.Db.Port,
		cfg.Notifications.Db.DbName,
	)

	conn, _ := pgxpool.Connect(ctx, connectString)

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	repo := db.NewNotifyRepo(conn)

	notify := notifyService.NewNotifyService(repo)

	saramaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Kafka.Brokers, saramaConfig)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	consumer := saga.NewNotifier(
		notify,
		syncProducer,
		3,
		cfg.Kafka.Delivered,
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{cfg.Kafka.Delivered}, consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()

}
