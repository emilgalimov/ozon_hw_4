package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/config"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/saga"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	saramacfg := sarama.NewConfig()
	saramacfg.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Kafka.Brokers, saramacfg)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	messages := 0

	for {
		messages++
		if messages > 10 {
			return
		}

		time.Sleep(time.Second)
		orderNumber := rand.Intn(1000000)
		productsCount := rand.Intn(3) + 1

		var products []int

		for i := 0; i < productsCount; i++ {
			products = append(products, rand.Intn(1000000))
		}

		message, _ := json.Marshal(saga.OrderToConfirmMessage{
			OrderID:    orderNumber,
			ProductIds: products,
		})
		par, off, err := syncProducer.SendMessage(&sarama.ProducerMessage{
			Topic: cfg.Kafka.Delivered,
			Key:   sarama.StringEncoder(strconv.Itoa(orderNumber)),
			Value: sarama.ByteEncoder(message),
		})
		log.Printf("%v -> %v; %v", par, off, err)
	}
}
