# \GetOrderAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrder**](GetOrderAPI.md#GetOrder) | **Get** /orders/2026-01-01/orders/{orderId} | 



## GetOrder

> GetOrderResponse GetOrder(ctx, orderId).IncludedData(includedData).Execute()





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
	orderId := "orderId_example" // string | An Amazon-defined order identifier.
	includedData := []string{"IncludedData_example"} // []string | A list of datasets to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetOrderAPI.GetOrder(context.Background(), orderId).IncludedData(includedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOrderAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: GetOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GetOrderAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | An Amazon-defined order identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includedData** | **[]string** | A list of datasets to include in the response. | 

### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

