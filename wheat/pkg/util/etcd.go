package util

import (
	"context"
	"log/slog"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func EtcdMgr() {
	dialTimeout := 5 * time.Second
	requestTimeout := 10 * time.Second
	endpoints := []string{"localhost:2379"}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		slog.Error("client", slog.Any("err", err))
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel() // Ensure the context is canceled

	key := "testk1"
	getResponse, err := client.Get(ctx, key)
	if err != nil {
		slog.Error("query", slog.Any("err", err))
	}
	if len(getResponse.Kvs) > 0 {
		k, v := string(getResponse.Kvs[0].Key), string(getResponse.Kvs[0].Value)
		slog.Debug("Retrieved key: %s, value: %s", k, v)
	} else {
		slog.Debug("Keynot found", slog.String("keyh", key))
	}
}
