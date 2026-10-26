package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dreamsxin/selling-partner-api-go/auth"
	"github.com/dreamsxin/selling-partner-api-go/client"
	"github.com/dreamsxin/selling-partner-api-go/generated/sellers"
)

func main() {
	credentials := auth.Credentials{
		ClientID:     "${SPAPI_LWA_CLIENT_ID}",
		ClientSecret: "${SPAPI_LWA_CLIENT_SECRET}",
		RefreshToken: "${SPAPI_LWA_REFRESH_TOKEN}",
	}
	tokens := auth.NewLWAProvider(credentials)
	httpClient := client.NewHTTPClient(client.HTTPClientOptions{
		TokenSource: tokens,
		UserAgent:   "MyERP/1.0 (Language=Go)",
		MaxRetries:  3,
	})

	cfg := sellers.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers[0].URL, _ = client.Endpoint(client.RegionNA)
	api := sellers.NewAPIClient(cfg)

	response, _, err := api.SellersAPI.GetMarketplaceParticipations(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marketplaces: %+v\n", response)
}
