package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bleemeo/bleemeo-go"
)

// Creating a dashboard and a widget
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

	dashboard, err := client.Create(context.Background(), bleemeo.ResourceDashboard, bleemeo.Body{"name": "My dashboard"})
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
	fmt.Println("View it on https://panel.bleemeo.com/dashboard/" + dashboardObj.ID)

	widget, err := client.Create(context.Background(), bleemeo.ResourceWidget, bleemeo.Body{"dashboard": dashboardObj.ID, "title": "My widget", "graph": bleemeo.Graph_Text})
	if err != nil {
		log.Fatalln("Error creating widget:", err)
	}

	var widgetObj struct {
		ID    string        `json:"id"`
		Title string        `json:"title"`
		Graph bleemeo.Graph `json:"graph"`
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling widget:", err)
	}

	fmt.Println("Successfully created widget:", widgetObj)
}
