package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/bleemeo/bleemeo-go"
)

func main() {
	client, err := bleemeo.NewClient(
		bleemeo.WithCredentialsFromEnv(),
		bleemeo.WithEndpoint("http://localhost:8000"),
		bleemeo.WithOAuthClientID("5c31cbfc-254a-4fb9-822d-e55c681a3d4f"),
	)
	if err != nil {
		log.Fatalln("Failed to initialize client:", err)
	}

	defer func() {
		err := client.Logout(context.Background())
		if err != nil {
			log.Fatalln("Logout:", err)
		}
	}()

	// Retrieving only the id and label of each metric:
	// the fewer fields required, the faster the query.
	iter := client.Iterator(bleemeo.ResourceMetric, bleemeo.Params{"fields": "id,label"})
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

	if err = iter.Err(); err != nil {
		if authErr := new(bleemeo.AuthError); errors.As(err, &authErr) {
			log.Fatalln("Authentication error:", authErr.ErrorCode, "/", authErr.Message)
		}

		log.Fatalln("Iteration error:", err)
	}

	fmt.Printf("Successfully listed %d metrics\n", count)
}
