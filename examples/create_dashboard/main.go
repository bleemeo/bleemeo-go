package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bleemeo/bleemeo-go"
)

// Creating a dashboard named "My dashboard",
// and inside a text widget with the content "My widget".
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

	body := map[string]any{"name": "My dashboard"}

	dashboard, err := client.Create(context.Background(), bleemeo.ResourceDashboard, body)
	if err != nil {
		log.Panicln("Error creating dashboard:", err)
	}

	var dashboardObj struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	err = json.Unmarshal(dashboard, &dashboardObj)
	if err != nil {
		log.Panicln("Error unmarshalling dashboard:", err)
	}

	log.Println("Successfully created dashboard:", dashboardObj)
	log.Println("View it on https://panel.bleemeo.com/dashboard/" + dashboardObj.ID)

	widget, err := client.Create(
		context.Background(),
		bleemeo.ResourceWidget,
		map[string]any{"dashboard": dashboardObj.ID, "title": "My widget", "graph": bleemeo.Graph_Text},
	)
	if err != nil {
		log.Panicln("Error creating widget:", err)
	}

	var widgetObj struct {
		ID    string        `json:"id"`
		Title string        `json:"title"`
		Graph bleemeo.Graph `json:"graph"`
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Panicln("Error unmarshalling widget:", err)
	}

	log.Println("Successfully created widget:", widgetObj)
}
