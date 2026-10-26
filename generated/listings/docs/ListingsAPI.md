# \ListingsAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListingsItem**](ListingsAPI.md#DeleteListingsItem) | **Delete** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**GetListingsItem**](ListingsAPI.md#GetListingsItem) | **Get** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PatchListingsItem**](ListingsAPI.md#PatchListingsItem) | **Patch** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PutListingsItem**](ListingsAPI.md#PutListingsItem) | **Put** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**SearchListingsItems**](ListingsAPI.md#SearchListingsItems) | **Get** /listings/2021-08-01/items/{sellerId} | 



## DeleteListingsItem

> ListingsItemSubmissionResponse DeleteListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).Execute()





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
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
	sku := "sku_example" // string | A selling partner-provided identifier for an Amazon listing.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon store identifiers for the request.
	issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when localization is not available for the specified locale. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingsAPI.DeleteListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.DeleteListingsItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteListingsItem`: ListingsItemSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.DeleteListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner-provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon store identifiers for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when localization is not available for the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsItem

> Item GetListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Execute()





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
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
	sku := "sku_example" // string | A selling partner-provided identifier for an Amazon listing.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon store identifiers for the request.
	issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when localization is not available for the specified locale. (optional)
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of data sets to include in the response. Default: `summaries`. (optional) (default to ["summaries"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingsAPI.GetListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.GetListingsItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsItem`: Item
	fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.GetListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner-provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon store identifiers for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when localization is not available for the specified locale. | 
 **includedData** | **[]string** | A comma-delimited list of data sets to include in the response. Default: &#x60;summaries&#x60;. | [default to [&quot;summaries&quot;]]

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListingsItem

> ListingsItemSubmissionResponse PatchListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IncludedData(includedData).Mode(mode).IssueLocale(issueLocale).Execute()





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
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
	sku := "sku_example" // string | A selling partner-provided identifier for an Amazon listing.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon store identifiers for the request.
	body := *openapiclient.NewListingsItemPatchRequest("ProductType_example", []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("Op_example", "Path_example")}) // ListingsItemPatchRequest | The request body schema for the `patchListingsItem` operation.
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of data sets to include in the response. Default: `issues`. (optional) (default to ["issues"])
	mode := "VALIDATION_PREVIEW" // string | Describes the mode of operation for the request. (optional)
	issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when localization is not available for the specified locale. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingsAPI.PatchListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IncludedData(includedData).Mode(mode).IssueLocale(issueLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PatchListingsItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListingsItem`: ListingsItemSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.PatchListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner-provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon store identifiers for the request. | 
 **body** | [**ListingsItemPatchRequest**](ListingsItemPatchRequest.md) | The request body schema for the &#x60;patchListingsItem&#x60; operation. | 
 **includedData** | **[]string** | A comma-delimited list of data sets to include in the response. Default: &#x60;issues&#x60;. | [default to [&quot;issues&quot;]]
 **mode** | **string** | Describes the mode of operation for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when localization is not available for the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutListingsItem

> ListingsItemSubmissionResponse PutListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IncludedData(includedData).Mode(mode).IssueLocale(issueLocale).Execute()





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
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
	sku := "sku_example" // string | A selling partner-provided identifier for an Amazon listing.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon store identifiers for the request.
	body := *openapiclient.NewListingsItemPutRequest("ProductType_example", map[string]interface{}{"key": interface{}(123)}) // ListingsItemPutRequest | The request body schema for the `putListingsItem` operation.
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of data sets to include in the response. Default: `issues`. (optional) (default to ["issues"])
	mode := "VALIDATION_PREVIEW" // string | Describes the mode of operation for the request. (optional)
	issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when localization is not available for the specified locale. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingsAPI.PutListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IncludedData(includedData).Mode(mode).IssueLocale(issueLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PutListingsItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutListingsItem`: ListingsItemSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.PutListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner-provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon store identifiers for the request. | 
 **body** | [**ListingsItemPutRequest**](ListingsItemPutRequest.md) | The request body schema for the &#x60;putListingsItem&#x60; operation. | 
 **includedData** | **[]string** | A comma-delimited list of data sets to include in the response. Default: &#x60;issues&#x60;. | [default to [&quot;issues&quot;]]
 **mode** | **string** | Describes the mode of operation for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first Amazon store is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when localization is not available for the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListingsItems

> ItemSearchResults SearchListingsItems(ctx, sellerId).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Identifiers(identifiers).IdentifiersType(identifiersType).VariationParentSku(variationParentSku).PackageHierarchySku(packageHierarchySku).CreatedAfter(createdAfter).CreatedBefore(createdBefore).LastUpdatedAfter(lastUpdatedAfter).LastUpdatedBefore(lastUpdatedBefore).WithIssueSeverity(withIssueSeverity).WithStatus(withStatus).WithoutStatus(withoutStatus).SortBy(sortBy).SortOrder(sortOrder).PageSize(pageSize).PageToken(pageToken).Execute()





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
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon store identifiers for the request.
	issueLocale := "en_US" // string | A locale that is used to localize issues. When not provided, the default language code of the first Amazon store is used. Examples: \"en_US\", \"fr_CA\", \"fr_FR\". When a localization is not available in the specified locale, localized messages default to \"en_US\". (optional)
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of data sets that you want to include in the response. Default: `summaries`. (optional) (default to ["summaries"])
	identifiers := []string{"Inner_example"} // []string | A comma-delimited list of product identifiers that you can use to search for listings items.   **Note**:  1. This is required when you specify `identifiersType`. 2. You cannot use 'identifiers' if you specify `variationParentSku` or `packageHierarchySku`. (optional)
	identifiersType := "SKU" // string | A type of product identifiers that you can use to search for listings items.   **Note**:  This is required when `identifiers` is provided. (optional)
	variationParentSku := "GM-ZDPI-9B4E" // string | Filters results to include listing items that are variation children of the specified SKU.   **Note**: You cannot use `variationParentSku` if you include `identifiers` or `packageHierarchySku` in your request. (optional)
	packageHierarchySku := "GM-ZDPI-9B4E" // string | Filter results to include listing items that contain or are contained by the specified SKU.   **Note**: You cannot use `packageHierarchySku` if you include `identifiers` or `variationParentSku` in your request. (optional)
	createdAfter := time.Now() // time.Time | A date-time that is used to filter listing items. The response includes listings items that were created at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. (optional)
	createdBefore := time.Now() // time.Time | A date-time that is used to filter listing items. The response includes listings items that were created at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. (optional)
	lastUpdatedAfter := time.Now() // time.Time | A date-time that is used to filter listing items. The response includes listings items that were last updated at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. (optional)
	lastUpdatedBefore := time.Now() // time.Time | A date-time that is used to filter listing items. The response includes listings items that were last updated at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. (optional)
	withIssueSeverity := []string{"WithIssueSeverity_example"} // []string | Filter results to include only listing items that have issues that match one or more of the specified severity levels. (optional)
	withStatus := []string{"WithStatus_example"} // []string | Filter results to include only listing items that have the specified status. (optional)
	withoutStatus := []string{"WithoutStatus_example"} // []string | Filter results to include only listing items that don't contain the specified statuses. (optional)
	sortBy := "lastUpdatedDate" // string | An attribute by which to sort the returned listing items. (optional) (default to "lastUpdatedDate")
	sortOrder := "DESC" // string | The order in which to sort the result items. (optional) (default to "DESC")
	pageSize := int32(9) // int32 | The number of results that you want to include on each page. (optional) (default to 10)
	pageToken := "sdlkj234lkj234lksjdflkjwdflkjsfdlkj234234234234" // string | A token that you can use to fetch a specific page when there are multiple pages of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListingsAPI.SearchListingsItems(context.Background(), sellerId).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Identifiers(identifiers).IdentifiersType(identifiersType).VariationParentSku(variationParentSku).PackageHierarchySku(packageHierarchySku).CreatedAfter(createdAfter).CreatedBefore(createdBefore).LastUpdatedAfter(lastUpdatedAfter).LastUpdatedBefore(lastUpdatedBefore).WithIssueSeverity(withIssueSeverity).WithStatus(withStatus).WithoutStatus(withoutStatus).SortBy(sortBy).SortOrder(sortOrder).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.SearchListingsItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListingsItems`: ItemSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.SearchListingsItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchListingsItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon store identifiers for the request. | 
 **issueLocale** | **string** | A locale that is used to localize issues. When not provided, the default language code of the first Amazon store is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. When a localization is not available in the specified locale, localized messages default to \&quot;en_US\&quot;. | 
 **includedData** | **[]string** | A comma-delimited list of data sets that you want to include in the response. Default: &#x60;summaries&#x60;. | [default to [&quot;summaries&quot;]]
 **identifiers** | **[]string** | A comma-delimited list of product identifiers that you can use to search for listings items.   **Note**:  1. This is required when you specify &#x60;identifiersType&#x60;. 2. You cannot use &#39;identifiers&#39; if you specify &#x60;variationParentSku&#x60; or &#x60;packageHierarchySku&#x60;. | 
 **identifiersType** | **string** | A type of product identifiers that you can use to search for listings items.   **Note**:  This is required when &#x60;identifiers&#x60; is provided. | 
 **variationParentSku** | **string** | Filters results to include listing items that are variation children of the specified SKU.   **Note**: You cannot use &#x60;variationParentSku&#x60; if you include &#x60;identifiers&#x60; or &#x60;packageHierarchySku&#x60; in your request. | 
 **packageHierarchySku** | **string** | Filter results to include listing items that contain or are contained by the specified SKU.   **Note**: You cannot use &#x60;packageHierarchySku&#x60; if you include &#x60;identifiers&#x60; or &#x60;variationParentSku&#x60; in your request. | 
 **createdAfter** | **time.Time** | A date-time that is used to filter listing items. The response includes listings items that were created at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | 
 **createdBefore** | **time.Time** | A date-time that is used to filter listing items. The response includes listings items that were created at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | 
 **lastUpdatedAfter** | **time.Time** | A date-time that is used to filter listing items. The response includes listings items that were last updated at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | 
 **lastUpdatedBefore** | **time.Time** | A date-time that is used to filter listing items. The response includes listings items that were last updated at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | 
 **withIssueSeverity** | **[]string** | Filter results to include only listing items that have issues that match one or more of the specified severity levels. | 
 **withStatus** | **[]string** | Filter results to include only listing items that have the specified status. | 
 **withoutStatus** | **[]string** | Filter results to include only listing items that don&#39;t contain the specified statuses. | 
 **sortBy** | **string** | An attribute by which to sort the returned listing items. | [default to &quot;lastUpdatedDate&quot;]
 **sortOrder** | **string** | The order in which to sort the result items. | [default to &quot;DESC&quot;]
 **pageSize** | **int32** | The number of results that you want to include on each page. | [default to 10]
 **pageToken** | **string** | A token that you can use to fetch a specific page when there are multiple pages of results. | 

### Return type

[**ItemSearchResults**](ItemSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

