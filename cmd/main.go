package main

import (
	"fmt"
	Client "github.com/codescalersinternships/DateTime-HTTP-Client-Mohamed-Riyad/pkg"
	"log"
)

func main() {

	config := *Client.NewConfig("http://localhost:8090", "400")
	client := Client.NewClient(config)
	output, err := client.GetDateTime("/datetime")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
