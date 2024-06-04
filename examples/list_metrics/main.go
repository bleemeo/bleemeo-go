package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/bleemeo/bleemeo-go"
)

// Listing the first 200 metrics of the account.
func main() {
	client, err := bleemeo.NewClient(
		bleemeo.WithConfigurationFromEnv(),
	)
	if err != nil {
		log.Panicln("Failed to initialize client:", err)
	}

	defer func() {
		err := client.Logout(context.Background())
		if err != nil {
			log.Panicln("Logout:", err)
		}
	}()

	// Retrieving only the id and label of each metric:
	// the fewer fields required, the faster the query.
	iter := client.Iterator(bleemeo.ResourceMetric, url.Values{"fields": {"id,label"}, "active": {"True"}})
	count := 0

	type metricType struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}

	for iter.Next(context.Background()) {
		var metricObj metricType

		err = json.Unmarshal(iter.At(), &metricObj)
		if err != nil {
			log.Panicln("Failed to unmarshal metric:", err)
		}

		count++
		if count <= 200 {
			log.Println("-> ", metricObj)
		} else if count == 201 {
			log.Println("Listing has more than 200 metrics, only the first 200 metrics are shown")
		}
	}

	if err = iter.Err(); err != nil {
		if authErr := new(bleemeo.AuthError); errors.As(err, &authErr) {
			// An AuthError is also an APIError
			log.Panicln("Authentication error:", authErr.ErrorCode, "/", authErr.Message)
		}

		if apiErr := new(bleemeo.APIError); errors.As(err, &apiErr) {
			log.Panicln("API error:", apiErr.StatusCode, "-", apiErr.Message)
		}

		log.Panicln("Iteration error:", err)
	}

	log.Printf("Successfully retrieved %d metrics from API\n", count)
}
