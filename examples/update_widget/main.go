package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"bleemeo"
)

func main() {
	client, err := bleemeo.NewClient(
		bleemeo.WithCredentialsFromEnv(),
		bleemeo.WithEndpoint("http://localhost:8000"),
		bleemeo.WithOAuthClientID("5c31cbfc-254a-4fb9-822d-e55c681a3d4f"),
	)
	if err != nil {
		log.Fatalln("Failed to initialize client", err)
	}

	pageNumber, pageSize := 1, 1

	resultPage, err := client.GetPage(context.Background(), bleemeo.Widget, pageNumber, pageSize, bleemeo.Params{"title": "My widget", "fields": "id"})
	if err != nil {
		log.Fatalln("Failed to fetch widget:", err)
	}

	if len(resultPage.Results) == 0 {
		log.Fatalln("Widget not found")
	}

	var widgetObj struct {
		ID    string            `json:"id"`
		Title string            `json:"title"`
		Graph bleemeo.GraphEnum `json:"graph"`
	}

	err = json.Unmarshal(resultPage.Results[0], &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling widget:", err)
	}

	widget, err := client.Update(context.Background(), bleemeo.Widget, widgetObj.ID, bleemeo.Body{"title": "This is my widget"})
	if err != nil {
		log.Fatalln("Failed to update widget:", err)
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling updated widget:", err)
	}

	fmt.Println("Successfully updated widget:", widgetObj)
}
