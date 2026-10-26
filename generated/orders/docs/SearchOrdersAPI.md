# \SearchOrdersAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOrders**](SearchOrdersAPI.md#SearchOrders) | **Get** /orders/2026-01-01/orders | 



## SearchOrders

> SearchOrdersResponse SearchOrders(ctx).CreatedAfter(createdAfter).CreatedBefore(createdBefore).LastUpdatedAfter(lastUpdatedAfter).LastUpdatedBefore(lastUpdatedBefore).FulfillmentStatuses(fulfillmentStatuses).MarketplaceIds(marketplaceIds).FulfilledBy(fulfilledBy).MaxResultsPerPage(maxResultsPerPage).PaginationToken(paginationToken).IncludedData(includedData).Execute()





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
	createdAfter := time.Now() // time.Time | The response includes orders created at or after this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of `createdAfter` and `lastUpdatedAfter` in your request. If `createdAfter` is provided, neither `lastUpdatedAfter` nor `lastUpdatedBefore` may be provided. (optional)
	createdBefore := time.Now() // time.Time | The response includes orders created at or before this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include `createdAfter` in the request, `createdBefore` is optional, and if provided must be equal to or after the `createdAfter` date and at least two minutes before the time of the request. If `createdBefore` is provided, neither `lastUpdatedAfter` nor `lastUpdatedBefore` may be provided. (optional)
	lastUpdatedAfter := time.Now() // time.Time | The response includes orders updated at or after this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of `createdAfter` and `lastUpdatedAfter`. If `lastUpdatedAfter` is provided, neither `createdAfter` nor `createdBefore` may be provided. (optional)
	lastUpdatedBefore := time.Now() // time.Time | The response includes orders updated at or before this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include `lastUpdatedAfter` in the request, `lastUpdatedBefore` is optional, and if provided must be equal to or after the `lastUpdatedAfter` date and at least two minutes before the time of the request. If `lastUpdatedBefore` is provided, neither `createdAfter` nor `createdBefore` may be provided. (optional)
	fulfillmentStatuses := []string{"FulfillmentStatuses_example"} // []string | A list of `FulfillmentStatus` values you can use to filter the results. (optional)
	marketplaceIds := []string{"Inner_example"} // []string | The response includes orders that were placed in marketplaces you include in this list.  Refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids) for a complete list of `marketplaceId` values. (optional)
	fulfilledBy := []string{"FulfilledBy_example"} // []string | The response includes orders that are fulfilled by the parties that you include in this list. (optional)
	maxResultsPerPage := int32(50) // int32 | The maximum number of orders that can be returned per page. The value must be between 1 and 100. **Default:** 100. (optional)
	paginationToken := "2YgYW55IGNhcm5hbCBwbGVhc3VyZS4" // string | Pagination occurs when a request produces a response that exceeds the `maxResultsPerPage`. This means that the response is divided into individual pages. To retrieve the next page, you must pass the `nextToken` value as the `paginationToken` query parameter in the next request. You will not receive a `nextToken` value on the last page. (optional)
	includedData := []string{"IncludedData_example"} // []string | A list of datasets to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchOrdersAPI.SearchOrders(context.Background()).CreatedAfter(createdAfter).CreatedBefore(createdBefore).LastUpdatedAfter(lastUpdatedAfter).LastUpdatedBefore(lastUpdatedBefore).FulfillmentStatuses(fulfillmentStatuses).MarketplaceIds(marketplaceIds).FulfilledBy(fulfilledBy).MaxResultsPerPage(maxResultsPerPage).PaginationToken(paginationToken).IncludedData(includedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchOrdersAPI.SearchOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrders`: SearchOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchOrdersAPI.SearchOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **time.Time** | The response includes orders created at or after this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of &#x60;createdAfter&#x60; and &#x60;lastUpdatedAfter&#x60; in your request. If &#x60;createdAfter&#x60; is provided, neither &#x60;lastUpdatedAfter&#x60; nor &#x60;lastUpdatedBefore&#x60; may be provided. | 
 **createdBefore** | **time.Time** | The response includes orders created at or before this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include &#x60;createdAfter&#x60; in the request, &#x60;createdBefore&#x60; is optional, and if provided must be equal to or after the &#x60;createdAfter&#x60; date and at least two minutes before the time of the request. If &#x60;createdBefore&#x60; is provided, neither &#x60;lastUpdatedAfter&#x60; nor &#x60;lastUpdatedBefore&#x60; may be provided. | 
 **lastUpdatedAfter** | **time.Time** | The response includes orders updated at or after this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of &#x60;createdAfter&#x60; and &#x60;lastUpdatedAfter&#x60;. If &#x60;lastUpdatedAfter&#x60; is provided, neither &#x60;createdAfter&#x60; nor &#x60;createdBefore&#x60; may be provided. | 
 **lastUpdatedBefore** | **time.Time** | The response includes orders updated at or before this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include &#x60;lastUpdatedAfter&#x60; in the request, &#x60;lastUpdatedBefore&#x60; is optional, and if provided must be equal to or after the &#x60;lastUpdatedAfter&#x60; date and at least two minutes before the time of the request. If &#x60;lastUpdatedBefore&#x60; is provided, neither &#x60;createdAfter&#x60; nor &#x60;createdBefore&#x60; may be provided. | 
 **fulfillmentStatuses** | **[]string** | A list of &#x60;FulfillmentStatus&#x60; values you can use to filter the results. | 
 **marketplaceIds** | **[]string** | The response includes orders that were placed in marketplaces you include in this list.  Refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids) for a complete list of &#x60;marketplaceId&#x60; values. | 
 **fulfilledBy** | **[]string** | The response includes orders that are fulfilled by the parties that you include in this list. | 
 **maxResultsPerPage** | **int32** | The maximum number of orders that can be returned per page. The value must be between 1 and 100. **Default:** 100. | 
 **paginationToken** | **string** | Pagination occurs when a request produces a response that exceeds the &#x60;maxResultsPerPage&#x60;. This means that the response is divided into individual pages. To retrieve the next page, you must pass the &#x60;nextToken&#x60; value as the &#x60;paginationToken&#x60; query parameter in the next request. You will not receive a &#x60;nextToken&#x60; value on the last page. | 
 **includedData** | **[]string** | A list of datasets to include in the response. | 

### Return type

[**SearchOrdersResponse**](SearchOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

