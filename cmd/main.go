package main

import (
	"encoding/json"
	"fmt"
	Client "github.com/codescalersinternships/DateTime-HTTP-Client-Mohamed-Riyad/pkg"
	"log"
	"time"
)

func main() {

	config := *Client.NewConfig("http://localhost:", "8090")
	client := Client.NewClient(config)
	output, err := client.GetDateTime("/datetime")
	if err != nil {
		log.Fatal(err)
	}
	var responseTime time.Time
	err = json.NewDecoder(output.Body).Decode(&responseTime)

	if err != nil {
		log.Fatalf("error decoding response body: %v", err)
	}

	got := responseTime.Round(time.Minute)
	fmt.Println(got)
}
