package cache

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

var ValkeyClient valkey.Client

func SetCache(address string) {
	var err error
	ValkeyClient, err = valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{address},
	})
	if err != nil {
		panic("Failed to open valkey connection")
	}
}

func SetKey(key string, value string, seconds int64) error {
	ctx := context.Background()
	err := ValkeyClient.Do(ctx, ValkeyClient.B().Set().Key(key).Value(value).ExSeconds(seconds).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}

func GetKey(key string) (string, error) {
	ctx := context.Background()
	value, err := ValkeyClient.Do(ctx, ValkeyClient.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return "", err
	}
	return value, nil
}

func DelKey(key string) error {
	ctx := context.Background()
	err := ValkeyClient.Do(ctx, ValkeyClient.B().Del().Key(key).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}
