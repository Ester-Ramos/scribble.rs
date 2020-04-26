package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/scribble-rs/scribble.rs/communication"
)

func main() {

	//Setting the seed in order for the petnames to be random.
	rand.Seed(time.Now().UnixNano())

	log.Println("Started.")

	//If this ever fails, it will return and print a fatal logger message
	log.Fatal(communication.Serve(portHTTP))
}
