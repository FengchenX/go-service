package etcddb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
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

func GetCli() *Client {
	if cli == nil {
		initCli()
	}
	return cli
}

var cli *Client

func initCli() {
	client, e := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if e != nil {
		fmt.Println("connect failed, err:", e)
		return
	}
	fmt.Println("connect succ")
	cli = NewCli(client)
}
