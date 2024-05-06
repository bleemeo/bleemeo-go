package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bleemeo/bleemeo-go"
)

// Finding a metric and counting its data points
func main() {
	client, err := bleemeo.NewClient(
		bleemeo.WithConfigurationFromEnv(),
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

	metricPage, err := client.GetPage(context.Background(), bleemeo.ResourceMetric, 1, 1, bleemeo.Params{"fields": "id,label", "active": "True"})
	if err != nil {
		log.Fatalln("Failed to fetch metric page:", err)
	}

	if len(metricPage.Results) == 0 {
		log.Fatalln("No metric found")
	}

	var metricObj struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}

	err = json.Unmarshal(metricPage.Results[0], &metricObj)
	if err != nil {
		log.Fatalln("Failed to unmarshal metric:", err)
	}

	resource := fmt.Sprintf("%s/%s/data/", bleemeo.ResourceMetric, metricObj.ID)

	statusCode, resp, err := client.Do(context.Background(), http.MethodGet, resource, nil, true, nil)
	if err != nil {
		log.Fatalln("Failed to fetch metric data:", err)
	}

	if statusCode != http.StatusOK {
		log.Fatalln("Unexpected status code:", statusCode)
	}

	var metricData struct {
		Values []struct {
			Time  time.Time `json:"time"`
			Value float64   `json:"value"`
		} `json:"values"`
	}

	err = json.Unmarshal(resp, &metricData)
	if err != nil {
		log.Fatalln("Failed to unmarshal metric data:", err)
	}

	fmt.Printf("Found %d data points for metric %q\n", len(metricData.Values), metricObj.Label)
}
