# \FbaInventoryAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventory**](FbaInventoryAPI.md#AddInventory) | **Post** /fba/inventory/v1/items/inventory | 
[**CreateInventoryItem**](FbaInventoryAPI.md#CreateInventoryItem) | **Post** /fba/inventory/v1/items | 
[**DeleteInventoryItem**](FbaInventoryAPI.md#DeleteInventoryItem) | **Delete** /fba/inventory/v1/items/{sellerSku} | 
[**GetInventorySummaries**](FbaInventoryAPI.md#GetInventorySummaries) | **Get** /fba/inventory/v1/summaries | 



## AddInventory

> AddInventoryResponse AddInventory(ctx).XAmznIdempotencyToken(xAmznIdempotencyToken).AddInventoryRequestBody(addInventoryRequestBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xAmznIdempotencyToken := "xAmznIdempotencyToken_example" // string | A unique token/requestId provided with each call to ensure idempotency.
	addInventoryRequestBody := *openapiclient.NewAddInventoryRequest() // AddInventoryRequest | List of items to add to Sandbox inventory.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FbaInventoryAPI.AddInventory(context.Background()).XAmznIdempotencyToken(xAmznIdempotencyToken).AddInventoryRequestBody(addInventoryRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FbaInventoryAPI.AddInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInventory`: AddInventoryResponse
	fmt.Fprintf(os.Stdout, "Response from `FbaInventoryAPI.AddInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAmznIdempotencyToken** | **string** | A unique token/requestId provided with each call to ensure idempotency. | 
 **addInventoryRequestBody** | [**AddInventoryRequest**](AddInventoryRequest.md) | List of items to add to Sandbox inventory. | 

### Return type

[**AddInventoryResponse**](AddInventoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInventoryItem

> CreateInventoryItemResponse CreateInventoryItem(ctx).CreateInventoryItemRequestBody(createInventoryItemRequestBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createInventoryItemRequestBody := *openapiclient.NewCreateInventoryItemRequest("SellerSku_example", "MarketplaceId_example", "ProductName_example") // CreateInventoryItemRequest | CreateInventoryItem Request Body Parameter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FbaInventoryAPI.CreateInventoryItem(context.Background()).CreateInventoryItemRequestBody(createInventoryItemRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FbaInventoryAPI.CreateInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInventoryItem`: CreateInventoryItemResponse
	fmt.Fprintf(os.Stdout, "Response from `FbaInventoryAPI.CreateInventoryItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInventoryItemRequestBody** | [**CreateInventoryItemRequest**](CreateInventoryItemRequest.md) | CreateInventoryItem Request Body Parameter. | 

### Return type

[**CreateInventoryItemResponse**](CreateInventoryItemResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInventoryItem

> DeleteInventoryItemResponse DeleteInventoryItem(ctx, sellerSku).MarketplaceId(marketplaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sellerSku := "sellerSku_example" // string | A single seller SKU used for querying the specified seller SKU inventory summaries.
	marketplaceId := "marketplaceId_example" // string | The marketplace ID for the marketplace for which the sellerSku is to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FbaInventoryAPI.DeleteInventoryItem(context.Background(), sellerSku).MarketplaceId(marketplaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FbaInventoryAPI.DeleteInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInventoryItem`: DeleteInventoryItemResponse
	fmt.Fprintf(os.Stdout, "Response from `FbaInventoryAPI.DeleteInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerSku** | **string** | A single seller SKU used for querying the specified seller SKU inventory summaries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceId** | **string** | The marketplace ID for the marketplace for which the sellerSku is to be deleted. | 

### Return type

[**DeleteInventoryItemResponse**](DeleteInventoryItemResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventorySummaries

> GetInventorySummariesResponse GetInventorySummaries(ctx).GranularityType(granularityType).GranularityId(granularityId).MarketplaceIds(marketplaceIds).Details(details).StartDateTime(startDateTime).SellerSkus(sellerSkus).SellerSku(sellerSku).NextToken(nextToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	granularityType := "granularityType_example" // string | The granularity type for the inventory aggregation level.
	granularityId := "granularityId_example" // string | The granularity ID for the inventory aggregation level.
	marketplaceIds := []string{"Inner_example"} // []string | The marketplace ID for the marketplace for which to return inventory summaries.
	details := true // bool | true to return inventory summaries with additional summarized inventory details and quantities. Otherwise, returns inventory summaries only (default value). (optional) (default to false)
	startDateTime := time.Now() // time.Time | A start date and time in ISO8601 format. If specified, all inventory summaries that have changed since then are returned. You must specify a date and time that is no earlier than 18 months prior to the date and time when you call the API. Note: Changes in inboundWorkingQuantity, inboundShippedQuantity and inboundReceivingQuantity are not detected. (optional)
	sellerSkus := []string{"Inner_example"} // []string | A list of seller SKUs for which to return inventory summaries. You may specify up to 50 SKUs. (optional)
	sellerSku := "sellerSku_example" // string | A single seller SKU used for querying the specified seller SKU inventory summaries. (optional)
	nextToken := "nextToken_example" // string | String token returned in the response of your previous request. The string token will expire 30 seconds after being created. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FbaInventoryAPI.GetInventorySummaries(context.Background()).GranularityType(granularityType).GranularityId(granularityId).MarketplaceIds(marketplaceIds).Details(details).StartDateTime(startDateTime).SellerSkus(sellerSkus).SellerSku(sellerSku).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FbaInventoryAPI.GetInventorySummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventorySummaries`: GetInventorySummariesResponse
	fmt.Fprintf(os.Stdout, "Response from `FbaInventoryAPI.GetInventorySummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInventorySummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **granularityType** | **string** | The granularity type for the inventory aggregation level. | 
 **granularityId** | **string** | The granularity ID for the inventory aggregation level. | 
 **marketplaceIds** | **[]string** | The marketplace ID for the marketplace for which to return inventory summaries. | 
 **details** | **bool** | true to return inventory summaries with additional summarized inventory details and quantities. Otherwise, returns inventory summaries only (default value). | [default to false]
 **startDateTime** | **time.Time** | A start date and time in ISO8601 format. If specified, all inventory summaries that have changed since then are returned. You must specify a date and time that is no earlier than 18 months prior to the date and time when you call the API. Note: Changes in inboundWorkingQuantity, inboundShippedQuantity and inboundReceivingQuantity are not detected. | 
 **sellerSkus** | **[]string** | A list of seller SKUs for which to return inventory summaries. You may specify up to 50 SKUs. | 
 **sellerSku** | **string** | A single seller SKU used for querying the specified seller SKU inventory summaries. | 
 **nextToken** | **string** | String token returned in the response of your previous request. The string token will expire 30 seconds after being created. | 

### Return type

[**GetInventorySummariesResponse**](GetInventorySummariesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

