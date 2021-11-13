package handlers

import (
	"KValDB/internal/model"
	"KValDB/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrNil = errors.New("Nil Key and Value")
)

type KeyValue struct {
	Key   string
	Value string
}

type KValDB struct {
}

func NewKVal() KValDB {
	return KValDB{}
}

func (kv *KValDB) Get(db string, key model.Key) (*KeyValue, error) {

	var v = model.KValDBType{}
	err := storage.UnMarshal("data/"+db, &v)
	if err != nil {
		return nil, err
	}

	if v[key] == "" {
		return nil, ErrNil
	}

	fmt.Printf("Reading key: %s - value: %s ", key, v[key])

	return &KeyValue{Key: string(key), Value: string(v[key])}, nil
}

func (kv *KValDB) Set(db string, key model.Key, value model.Value) error {
	var v = model.KValDBType{}
	err := storage.UnMarshal("data/"+db, &v)
	if err != nil {
		return err
	}

	v[key] = value

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	fmt.Printf("Setting key: %s - value: %s ", key, value)

	return storage.WriteFile("data/"+db, b)
}
