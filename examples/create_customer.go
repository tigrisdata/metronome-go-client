package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adilansari/metronome-go-client"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
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

	createCustomerBody := metronome.CreateCustomerJSONRequestBody{
		IngestAliases: &[]string{"my_customer_alias"},
		Name:          "my_customer_id",
	}

	resp, err := client.CreateCustomer(context.TODO(), createCustomerBody)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("metronome request failed: %s", resp.Status))
	}

	parsed, err := metronome.ParseCreateCustomerResponse(resp)
	if err != nil {
		panic(err)
	}

	// Sample output:
	// request succeeded: {my_customer_alias 8d433a54-f281-499c-a4fa-6ec84f3d6157 [my_customer_alias] my_customer_id}
	fmt.Printf("request succeeded: %v", parsed.JSON200.Data)
}
