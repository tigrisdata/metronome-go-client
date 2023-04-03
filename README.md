[![Go Reference](https://pkg.go.dev/badge/github.com/adilansari/metronome-go-client/cobra.svg)](https://pkg.go.dev/github.com/adilansari/metronome-go-client)
[![rcard](https://goreportcard.com/badge/github.com/adilansari/metronome-go-client)](https://goreportcard.com/report/github.com/adilansari/metronome-go-client)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/adilansari/metronome-go-client/main/LICENSE)

[//]: # ([![Sourcegraph]&#40;https://sourcegraph.com/github.com/adilansari/metronome-go-client/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/adilansari/metronome-go-client?badge&#41;)

# [Metronome](https://docs.metronome.com/api/) Go client
Go client for Metronome generated from its OpenAPI spec

> This is not an official client, and was generated for personal use from their [OpenAPI spec](https://docs.metronome.com/api/).

## Installation

By using this client you'll avoid a lot of boilerplate code to perform all the 
marshalling and unmarshalling into objects when using Metronome HTTP APIs.

Install in your project to have it as a dependency in `go.mod`:
```shell
go get github.com/adilansari/metronome-go-client@latest
```

## Usage

##### Authentication

Metronome requires all HTTP requests to [include a authorization header](https://docs.metronome.com/using-the-api/authorization/)
for authentication:

```shell
GET /v1/... HTTP/1.1
Host: api.metronome.com
Authorization: Bearer MY_TOKEN
```

Instead of adding authorization header to every request when using this go client, you can
simply register a callback as:

```golang
authProvider, err := securityprovider.NewSecurityProviderBearerToken("MY_TOKEN")
if err != nil {
	panic(err)
}

client, err := metronome.NewClient("https://api.metronome.com/v1", metronome.WithRequestEditorFn(authProvider.Intercept))
```

##### Create customer example

Following is a complete example you can try in your project. Replace the `REPLACE_ME` with your
own token generated from Metronome account.

```golang
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"

	"github.com/adilansari/metronome-go-client"
)

func main() {
	// API key based authentication. Add your metronome auth token here.
	// Learn more about auth: https://docs.metronome.com/using-the-api/authorization/
	authProvider, err := securityprovider.NewSecurityProviderBearerToken("REPLACE_ME")
	if err != nil {
		panic(err)
	}

	// Client will be used to perform all operations on metronome.
	// Auth provider generated above will implicitly add an authorization token to all requests.
	client, err := metronome.NewClient("https://api.metronome.com/v1", metronome.WithRequestEditorFn(authProvider.Intercept))
	if err != nil {
		panic(err)
	}

	// JSON request for CreateCustomer
	createCustomerBody := metronome.CreateCustomerJSONRequestBody{
		IngestAliases: &[]string{"my_customer_alias"},
		Name:          "my_customer_id",
	}

	// HTTP POST call to "/customers" endpoint
	resp, err := client.CreateCustomer(context.TODO(), createCustomerBody)
	if err != nil {
		panic(err)
	}

	// Checking if request succeeded
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("metronome request failed: %s", resp.Status))
	}

	// Using built-in parsing to convert response to a native Go struct
	parsed, err := metronome.ParseCreateCustomerResponse(resp)
	if err != nil {
		panic(err)
	}

	// Sample output:
	// request succeeded: {ExternalId:my_customer_alias Id:7d8e8341-f271-4738-8a7f-cf7d6c2418f3 IngestAliases:[my_customer_alias] Name:my_customer_id}
	fmt.Printf("request succeeded: %+v", parsed.JSON200.Data)
}
```

## Requesting updates
- Since this is not official client, the API may be out of date
- Please open an issue in the repo and I'll try to update the client to latest


