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
	client := bleemeo.NewClient(username, password)

	resultPage, err := client.List(context.Background(), "widget", 1, 1, bleemeo.Params{"title": "My widget", "fields": "id"})
	if err != nil {
		log.Fatalln("Failed to fetch widget:", err)
	}

	if len(resultPage.Results) == 0 {
		log.Fatalln("Widget not found")
	}

	var widgetObj struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	err = json.Unmarshal(resultPage.Results[0], &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling widget:", err)
	}

	widget, err := client.Update(context.Background(), "widget", widgetObj.ID, bleemeo.Body{"title": "This is my widget"}, bleemeo.Fields{"id", "title"})
	if err != nil {
		log.Fatalln("Failed to update widget:", err)
	}

	err = json.Unmarshal(widget, &widgetObj)
	if err != nil {
		log.Fatalln("Error unmarshalling updated widget:", err)
	}

	fmt.Println("Successfully updated widget:", widgetObj)
}
