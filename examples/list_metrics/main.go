package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"bleemeo"
	"bleemeo/examples"
)

func main() {
	username, password := examples.ParseArguments()
	client := bleemeo.NewClient(bleemeo.WithCredentials(username, password))

	iter := client.Iterator("metrics", bleemeo.Params{"fields": "id,label"})
	count := 0

	type metricType struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}

	for iter.Next(context.Background()) {
		var metricObj metricType

		err := json.Unmarshal(iter.At(), &metricObj)
		if err != nil {
			log.Fatalln("Failed to unmarshal metric:", err)
		}

		fmt.Println("->", metricObj)

		count++
		if count == 200 {
			fmt.Println("Listing has at least 200 metrics, only the first 200 metrics are shown")

			break
		}
	}

	if err := iter.Err(); err != nil {
		log.Fatalln("Iteration error:", err)
	}

	fmt.Printf("Successfully listed %d metrics\n", count)
}
