package main

import (
	"github.com/denverdino/aliyungo/slb"
)

func AliClient(Key string, Secret string) *slb.Client {
	client := slb.NewClient(Key, Secret)

	return client
}