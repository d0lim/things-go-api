package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if os.Getenv("TEST_THINGS_API") != "" {
		log.Println("Testing golang things api")
	}

	fmt.Println("Testing ended")
}
