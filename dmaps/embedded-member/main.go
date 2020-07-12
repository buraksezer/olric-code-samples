package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/buraksezer/olric"
	"github.com/buraksezer/olric/config"
)

func main() {
	// Deployment scenario: embedded-member
	// This creates a single-node Olric cluster. It's good enough for experimenting.

	// config.New returns a new config.Config with sane defaults. Available values for env:
	// local, lan, wan
	c := config.New("local")

	// Callback function. It's called when this node is ready to accept connections.
	ctx, cancel := context.WithCancel(context.Background())
	c.Started = func() {
		defer cancel()
		log.Println("[INFO] Olric is ready to accept connections")
	}

	db, err := olric.New(c)
	if err != nil {
		log.Fatalf("Failed to create Olric instance: %v", err)
	}

	go func() {
		// Call Start at background. It's a blocker call.
		err = db.Start()
		if err != nil {
			log.Fatalf("olric.Start returned an error: %v", err)
		}
	}()

	<-ctx.Done()

	dm, err := db.NewDMap("bucket-of-arbitrary-items")
	if err != nil {
		log.Fatalf("olric.NewDMap returned an error: %v", err)
	}

	// Magic starts here!
	fmt.Println("##")
	fmt.Println("Operations on a DMap instance:")
	err = dm.Put("string-key", "buraksezer")
	if err != nil {
		log.Fatalf("Failed to call Put: %v", err)
	}
	stringValue, err := dm.Get("string-key")
	if err != nil {
		log.Fatalf("Failed to call Get: %v", err)
	}
	fmt.Printf("Value for string-key: %v, reflect.TypeOf: %s\n", stringValue, reflect.TypeOf(stringValue))

	err = dm.Put("uint64-key", uint64(1988))
	if err != nil {
		log.Fatalf("Failed to call Put: %v", err)
	}
	uint64Value, err := dm.Get("uint64-key")
	if err != nil {
		log.Fatalf("Failed to call Get: %v", err)
	}
	fmt.Printf("Value for uint64-key: %v, reflect.TypeOf: %s\n", uint64Value, reflect.TypeOf(uint64Value))
	fmt.Println("##")

	// Don't forget the call Shutdown when you want to leave the cluster.
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to shutdown Olric: %v", err)
	}
}