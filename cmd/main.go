package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var endpoints = []string{"localhost:32771", "localhost:32772", "localhost:32773"}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	log.Println("etcd client created")

	go func() {
		time.Sleep(1 * time.Second)
		ctx := context.Background()

		// Set key value
		putRes, err := cli.Put(ctx, "log_level", "error")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("put: %v", putRes)

		time.Sleep(3 * time.Second)

		// Create a new lease
		l := clientv3.NewLease(cli)
		lgp, err := l.Grant(ctx, int64((5 * time.Second).Seconds()))
		if err != nil {
			log.Fatal(err)
		}

		// Set lease to key
		putRes, err = cli.Put(ctx, "log_level", "", clientv3.WithLease(lgp.ID), clientv3.WithIgnoreValue())
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("put: %v", putRes)

		time.Sleep(10 * time.Second)

		// Close client
		_ = cli.Close()
	}()

	// Watch key changes
	watchChan := cli.Watcher.Watch(ctx, "log_level")
	for resp := range watchChan {
		log.Printf("watch_response: %+v\n", resp)
		for _, e := range resp.Events {
			log.Printf("watch_event: %+v\n", e)
		}
	}

	log.Print("Done!")
}
