package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/buraksezer/olric/client"
	"github.com/buraksezer/olric/serializer"
)

func main() {
	cc := &client.Config{
		Addrs:      []string{"127.0.0.1:3320"},
		MaxConn:    10,
		Serializer: serializer.NewMsgpackSerializer(),
	}

	// Create a new client instance
	c, err := client.New(cc)
	if err != nil {
		log.Fatalf("Olric client returned error: %s", err)
	}
	defer c.Close()

	// Create a DMap instance
	dm := c.NewDMap("my-dmap")
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i)
		value := strconv.Itoa(i * 2)
		if err = dm.Put(key, value); err != nil {
			log.Fatalf("put failed for %s: %s", key, err)
		}
		fmt.Printf("[PUT] Key: %s, Value: %s\n", key, value)
	}

	fmt.Printf("\n")

	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i)
		value, err := dm.Get(key)
		if err != nil {
			log.Fatalf("get failed for %s: %s", err)
		}
		fmt.Printf("[GET] Key: %s, Value: %s\n", key, value)
	}
}