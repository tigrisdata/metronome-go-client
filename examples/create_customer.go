package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"

	"github.com/tigrisdata/metronome-go-client"
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
