package main

import (
	"context"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

func main() {
	name := flag.String("name", "", "Enter a name")
	flag.Parse()

	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cli)

	// create a sessions to elect a Leader
	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	e := concurrency.NewElection(s, "/leader-election")
	ctx := context.Background()
	// Elect a leader (or wait that the leader resign)
	if err := e.Campaign(ctx, "e"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("leader election for ", *name)

	fmt.Println("Do some work in", *name)
	time.Sleep(5 * time.Second)
	if err := e.Resign(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("resign ", *name)

}
