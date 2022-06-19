package caches

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"
	"strconv"
)

type storageMemcached struct {
	client *memcache.Client
}

func NewStorageMemcached(server ...string) *storageMemcached {
	return &storageMemcached{
		client: memcache.New(server...),
	}
}

func (s storageMemcached) Get(orderID int) (storeItems []storage.StoreItem, err error) {
	marshalledItems, err := s.client.Get(strconv.Itoa(orderID))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshalledItems.Value, &storeItems)
	if err != nil {
		return nil, err
	}

	return
}

func (s storageMemcached) Set(orderID int, items []storage.StoreItem) error {
	itemsBytes, _ := json.Marshal(items)

	err := s.client.Set(&memcache.Item{
		Key:   strconv.Itoa(orderID),
		Value: itemsBytes,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s storageMemcached) Delete(orderID int) error {
	return s.client.Delete(strconv.Itoa(orderID))
}
