package main

import (
	"context"
	"encoding/json"
	"log"
	"net/url"

	"github.com/bleemeo/bleemeo-go"
)

// Updating the content of a widget.
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

	pageNumber, pageSize := 1, 1

	resultPage, err := client.GetPage(
		context.Background(),
		bleemeo.ResourceWidget,
		pageNumber, pageSize,
		url.Values{"title": {"My widget"}, "fields": {"id,dashboard"}},
	)
	if err != nil {
		log.Panicln("Failed to fetch widget:", err)
	}

	if len(resultPage.Results) == 0 {
		log.Panicln("Widget not found")
	}

	var widgetObj struct {
		ID          string        `json:"id"`
		Title       string        `json:"title"`
		Graph       bleemeo.Graph `json:"graph"`
		DashboardID string        `json:"dashboard"`
	}

	err = json.Unmarshal(resultPage.Results[0], &widgetObj)
	if err != nil {
		log.Panicln("Error unmarshalling widget:", err)
	}

	widget, err := client.Update(
		context.Background(),
		bleemeo.ResourceWidget,
		widgetObj.ID,
		map[string]any{"title": "This is my widget"},
	)
	if err != nil {
		log.Panicln("Failed to update widget:", err)
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Panicln("Error unmarshalling updated widget:", err)
	}

	log.Println("Successfully updated widget:", widgetObj)
	log.Println("Check it on https://panel.bleemeo.com/dashboard/" + widgetObj.DashboardID)
}
