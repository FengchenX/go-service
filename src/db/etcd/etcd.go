package etcd

import (
	"conf"
	"context"
	"encoding/json"
	"fmt"
	//"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Client struct {
	*clientv3.Client
}

func NewCli(cli *clientv3.Client) *Client {
	return &Client{
		Client: cli,
	}
}

func (cli Client) Put(key string, value interface{}, opts ...clientv3.OpOption) error {
	buf, e := json.Marshal(&value)
	if e != nil {
		return e
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, e = cli.Client.Put(ctx, key, string(buf), opts...)
	cancel()
	if e != nil {
		return e
	}
	return nil
}

func (cli Client) Get(key string, value interface{}, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, e := cli.Client.Get(ctx, key, opts...)
	cancel()
	if e != nil {
		return e
	}
	if len(resp.Kvs) == 0 {
		return fmt.Errorf("no this key")
	}
	for _, ev := range resp.Kvs {
		e := json.Unmarshal(ev.Value, value)
		return e

	}
	return nil
}

func DefaultCli() *Client {
	client, e := clientv3.New(clientv3.Config{
		Endpoints:   conf.AgfunInst().Etcd,
		DialTimeout: 5 * time.Second,
	})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("connect succ")
	return NewCli(client)
}
