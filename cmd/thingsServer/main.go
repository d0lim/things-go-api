package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/d0lim/things-go-api/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("THINGS_USERNAME")
	password := os.Getenv("THINGS_PASSWORD")

	fmt.Println("username: ", username, " // password: ", password)

	c := api.New(api.APIEndpoint, username, password)

	_, err = c.Verify()
	if err != nil {
		log.Fatalf("Login failed: %q\nPlease check your credentials.", err.Error())
	}
	fmt.Printf("User: %s\n", c.EMail)

	history, err := c.OwnHistory()
	if err != nil {
		log.Fatalf("Failed to lookup own history key: %q\n", err.Error())
	}
	fmt.Printf("Own History Key: %s\n", history.ID)

	history.Sync()

	// items, _, err := history.Items(thingscloud.ItemsOptions{StartIndex: 0})
	// if err != nil {
	// 	log.Fatalf("Failed to lookup items: %q\n", err.Error())
	// }

	// fmt.Println(items);

	fmt.Println("Testing ended")
}
