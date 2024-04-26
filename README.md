# Bleemeo Go

Go library for the Bleemeo API

### Basic usage

Retrieving a metric by ID, only interested in the `id` and `label` fields.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bleemeo/bleemeo-go"
)

func main() {
	client, err := bleemeo.NewClient(bleemeo.WithCredentialsFromEnv())
	if err != nil {
		log.Fatalln("Failed to initialize client:", err)
	}

	metric, err := client.Get(context.Background(), bleemeo.Metric, "<the metric UUID>", bleemeo.Fields{"id", "label"})
	if err != nil {
		log.Fatalln("Failed to retrieve metric:", err)
	}

	var metricObj struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}

	err = json.Unmarshal(metric, &metricObj)
	if err != nil {
		log.Fatalln("Failed to unmarshal metric:", err)
	}

	fmt.Printf("The metric with the id %s is %q\n", metricObj.ID, metricObj.Label)
}
```

> More examples can be found in [examples](./examples)
