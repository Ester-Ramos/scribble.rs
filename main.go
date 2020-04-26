package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	//Setting the seed in order for the petnames to be random.
	rand.Seed(time.Now().UnixNano())

	log.Println("Started.")
}
