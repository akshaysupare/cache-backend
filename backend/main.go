package main

import (
	"fmt"
	"time"

	"backend-config.Cache/config"
	"backend-config.Cache/router"

	"github.com/dgraph-io/ristretto"
)

var err error
func main() {
	
	// Create a new config.Cache with a capacity of 1024.
	config.Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1024,    // maximum cost of config.Cache  in bytes
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic(err)
	}
	// set a value with a cost of 1 with no expiry
	config.Cache.Set("key", "value", 1)

	// set a value for test with a cost of 1 and for expiry time 5 Min 
	config.Cache.SetWithTTL("1", "akshay", 1, 600 * time.Second)

	// wait for value to pass through buffers
	config.Cache.Wait()

	// get value from config.Cache
	value, found := config.Cache.Get("1")
	if !found {
		fmt.Println("missing value")
	}
	fmt.Println("value for key 1 is " , value)

	defer config.Cache.Close()

	//Initalising router
	router.InitRoutes()

	
}
