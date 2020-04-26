package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/scribble-rs/scribble.rs/communication"
)

func main() {

	//Setting the seed in order for the petnames to be random.
	rand.Seed(time.Now().UnixNano())

	log.Println("Started.")
}
