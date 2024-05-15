package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/bleemeo/bleemeo-go"
)

// Updating the content of a widget
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

	pageNumber, pageSize := 1, 1

	resultPage, err := client.GetPage(
		context.Background(),
		bleemeo.ResourceWidget,
		pageNumber, pageSize,
		url.Values{"title": {"My widget"}, "fields": {"id,dashboard"}},
	)
	if err != nil {
		log.Fatalln("Failed to fetch widget:", err)
	}

	if len(resultPage.Results) == 0 {
		log.Fatalln("Widget not found")
	}

	var widgetObj struct {
		ID          string        `json:"id"`
		Title       string        `json:"title"`
		Graph       bleemeo.Graph `json:"graph"`
		DashboardID string        `json:"dashboard"`
	}

	err = json.Unmarshal(resultPage.Results[0], &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling widget:", err)
	}

	widget, err := client.Update(
		context.Background(),
		bleemeo.ResourceWidget,
		widgetObj.ID,
		map[string]any{"title": "This is my widget"},
	)
	if err != nil {
		log.Fatalln("Failed to update widget:", err)
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling updated widget:", err)
	}

	fmt.Println("Successfully updated widget:", widgetObj)
	fmt.Println("Check it on https://panel.bleemeo.com/dashboard/" + widgetObj.DashboardID)
}
