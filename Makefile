#!/bin/bash
include .env.default
include .env

notify:
	go run ./cmd/notify/notify.go

pay:
	go run ./cmd/pay/pay.go

storage:
	go run ./cmd/storage/storage.go

producer:
	go run ./cmd/producer/producer.go

migrate: migratenotify migratepay migratestorage

migratestorage:
	cd migrations/storage; ./goose postgres "host=${DB_HOST_1} port=${DB_PORT_1} user=${DB_USER_1} password=${DB_PASSWORD_1} dbname=${DB_NAME_1} sslmode=disable" up

migratepay:
	cd migrations/pay; ./goose postgres "host=${DB_HOST_2} port=${DB_PORT_2} user=${DB_USER_2} password=${DB_PASSWORD_2} dbname=${DB_NAME_2} sslmode=disable" up

migratenotify:
	cd migrations/notify; ./goose postgres "host=${DB_HOST_3} port=${DB_PORT_3} user=${DB_USER_3} password=${DB_PASSWORD_3} dbname=${DB_NAME_3} sslmode=disable" up

