# Bleemeo Go

### Go library for interacting with the Bleemeo API

## Requirements

- Go1.18 or later
- An account on [Bleemeo](https://bleemeo.com/)

### Environment

At least the following options should be configured (as environment variables or with options):

- Credentials OR initial refresh token

> Ways to provide those options are referenced in the [Configuration](#configuration) section.

## Basic usage

Retrieving a metric by ID, expecting the model's default fields.

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
	client, err := bleemeo.NewClient(bleemeo.WithConfigurationFromEnv())
	if err != nil {
		log.Fatalln("Failed to initialize client:", err)
	}

	defer func() {
		err := client.Logout(context.Background())
		if err != nil {
			log.Fatalln("Logout:", err)
		}
	}()

	metric, err := client.Get(context.Background(), bleemeo.ResourceMetric, "<the metric UUID>")
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

## Configuration

**For environment variables to be taken into account, the option `WithConfigurationFromEnv()` must be provided.**

| Property                      | Option function                        | Env variable(s)                                           |
|-------------------------------|----------------------------------------|-----------------------------------------------------------|
| Credentials                   | `WithCredentials(username, password)`  | `BLEEMEO_USER` & `BLEEMEO_PASSWORD`                       |
| Bleemeo account header        | `WithBleemeoAccountHeader(accountID)`  | `BLEEMEO_ACCOUNT_ID`                                      |
| OAuth client ID/secret        | `WithOAuthClient(id, secret)`          | `BLEEMEO_OAUTH_CLIENT_ID` & `BLEEMEO_OAUTH_CLIENT_SECRET` |
| Endpoint URL                  | `WithEndpoint(endpoint)`               | `BLEEMEO_API_URL`                                         |
| Initial refresh token         | `WithInitialOAuthRefreshToken(token)`  | `BLEEMEO_OAUTH_INITIAL_REFRESH_TOKEN`                     |
| HTTP client                   | `WithHTTPClient(client)`               | -                                                         |
| New OAuth token callback      | `WithNewOAuthTokenCallback(callback)`  | -                                                         |
| Throttle max auto retry delay | `WithThrottleMaxAutoRetryDelay(delay)` | -                                                         |
