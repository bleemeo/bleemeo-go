package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"bleemeo"
)

func main() {
	client := bleemeo.NewClient(
		bleemeo.WithCredentialsFromEnv(),
		bleemeo.WithEndpoint("http://localhost:8000"),
		bleemeo.WithOAuthClientID("5c31cbfc-254a-4fb9-822d-e55c681a3d4f"),
	)

	dashboard, err := client.Create(context.Background(), bleemeo.Dashboard, bleemeo.Body{"name": "My dashboard"})
	if err != nil {
		log.Fatalln("Error creating dashboard:", err)
	}

	var dashboardObj struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	err = json.Unmarshal(dashboard, &dashboardObj)
	if err != nil {
		log.Fatalln("Error unmarshalling dashboard:", err)
	}

	fmt.Println("Successfully created dashboard:", dashboardObj)

	widget, err := client.Create(context.Background(), "widget", bleemeo.Body{"dashboard": dashboardObj.ID, "name": "My widget"})
	if err != nil {
		log.Fatalln("Error creating widget:", err)
	}

	var widgetObj struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling widget:", err)
	}

	fmt.Println("Successfully created widget:", widgetObj)
}
