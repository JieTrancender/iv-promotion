package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var prefix = "/config/nsq_to_dingding/"

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.3.212:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}
	defer cli.Close()

	getResp, _ := cli.Get(context.TODO(), prefix, clientv3.WithPrefix())
	watchStartVer := getResp.Header.Revision + 1
	fmt.Println(getResp)
	watchChan := cli.Watch(context.Background(), prefix, clientv3.WithRev(watchStartVer), clientv3.WithPrefix())
	for resp := range watchChan {
		for _, ev := range resp.Events {
			fmt.Printf("Type: %v,Key: %v, Value: %v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
