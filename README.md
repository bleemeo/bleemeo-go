# Bleemeo Go [![Go Report Card](https://goreportcard.com/badge/github.com/bleemeo/bleemeo-go)](https://goreportcard.com/report/github.com/bleemeo/bleemeo-go) [![Go Reference](https://pkg.go.dev/badge/github.com/bleemeo/bleemeo-go.svg)](https://pkg.go.dev/github.com/bleemeo/bleemeo-go) [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/bleemeo/bleemeo-go/blob/main/LICENSE)

Go library for interacting with the Bleemeo API

## Requirements

- Go1.23 or later
- An account on [Bleemeo](https://bleemeo.com/)

## Installation

On your existing Go project, run:

```
go get -u github.com/bleemeo/bleemeo-go
```

If you start a new project, use go mod init to bootstrap the Go project:

```
go mod init github.com/my-username/my-project
```

Then do the first command.

## Documentation

The Go library documentation can be found [here](https://pkg.go.dev/github.com/bleemeo/bleemeo-go).

Some examples of library usage can be found in [examples](./examples).

## Basic usage

Listing the first 10 agents in your account:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	const pageSize = 10

	page, err := client.GetPage(context.Background(), bleemeo.ResourceAgent, 1, pageSize, url.Values{"active": {"true"}})
	if err != nil {
		log.Fatalln("Failed to list agents:", err)
	}

	for _, row := range page.Results {
		var agentObj struct {
			ID          string `json:"id"`
			FQDN        string `json:"fqdn"`
			DisplayName string `json:"display_name"`
		}

		err = json.Unmarshal(row, &agentObj)
		if err != nil {
			log.Fatalln("Failed to unmarshal agent:", err)
		}

		fmt.Printf("* Agent %s (fqdn = %s, id = %s)\n", agentObj.DisplayName, agentObj.FQDN, agentObj.ID)
	}
}
```

Save this file in list-agents.go.

If not already in an existing Go project, create the project:

```
go mod init github.com/my-username/my-project
```

Run with:

```
go get -u github.com/bleemeo/bleemeo-go
BLEEMEO_USER=user-email@domain.com BLEEMEO_PASSWORD=password go run list-agents.go
```

> More examples can be found in [examples](./examples)

To run an example, from a clone of this repository run the following:

```
BLEEMEO_USER=user-email@domain.com BLEEMEO_PASSWORD=password go run ./examples/list_metrics/
```

## Environment

At least the following options must be configured (as environment variables or with options):

- Credentials OR initial refresh token
- All other configuration options are optional and may be omitted

> Ways to provide those options are referenced in the [Configuration](#configuration) section.

## Configuration

**For environment variables to be taken into account, the option `WithConfigurationFromEnv()` must be provided.**

| Property                      | Option function                        | Env variable(s)                                           | Default values                                                                                   |
|-------------------------------|----------------------------------------|-----------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| Credentials                   | `WithCredentials(username, password)`  | `BLEEMEO_USER` & `BLEEMEO_PASSWORD`                       | None. This option is required (unless initial refresh token is used)                             |
| Bleemeo account header        | `WithBleemeoAccountHeader(accountID)`  | `BLEEMEO_ACCOUNT_ID`                                      | The first account associated with used credentials.                                              |
| OAuth client ID/secret        | `WithOAuthClient(id, secret)`          | `BLEEMEO_OAUTH_CLIENT_ID` & `BLEEMEO_OAUTH_CLIENT_SECRET` | The default SDK OAuth client ID                                                                  |
| Endpoint URL                  | `WithEndpoint(endpoint)`               | `BLEEMEO_API_URL`                                         | `https://api.bleemeo.com`                                                                        |
| Initial refresh token         | `WithInitialOAuthRefreshToken(token)`  | `BLEEMEO_OAUTH_INITIAL_REFRESH_TOKEN`                     | None. This is an alternative to username & password credentials.                                 |
| HTTP client                   | `WithHTTPClient(client)`               | -                                                         | None. This option allow to customize behavior of the HTTP client.                                |
| New OAuth token callback      | `WithNewOAuthTokenCallback(callback)`  | -                                                         | None. This option allow to get access to refresh token, useful for initial refresh token option. |
| Throttle max auto retry delay | `WithThrottleMaxAutoRetryDelay(delay)` | -                                                         | 1 minute.                                                                                        |
