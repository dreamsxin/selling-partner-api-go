# \CatalogAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogItem**](CatalogAPI.md#GetCatalogItem) | **Get** /catalog/2022-04-01/items/{asin} | 
[**SearchCatalogItems**](CatalogAPI.md#SearchCatalogItems) | **Get** /catalog/2022-04-01/items | 



## GetCatalogItem

> Item GetCatalogItem(ctx, asin).MarketplaceIds(marketplaceIds).IncludedData(includedData).Locale(locale).Execute()





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
	asin := "asin_example" // string | The Amazon Standard Identification Number (ASIN) of the item.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of datasets to include in the response. (optional) (default to ["summaries"])
	locale := "en_US" // string | The locale for which you want to retrieve localized summaries. Defaults to the primary locale of the marketplace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.GetCatalogItem(context.Background(), asin).MarketplaceIds(marketplaceIds).IncludedData(includedData).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetCatalogItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogItem`: Item
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | The Amazon Standard Identification Number (ASIN) of the item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
 **includedData** | **[]string** | A comma-delimited list of datasets to include in the response. | [default to [&quot;summaries&quot;]]
 **locale** | **string** | The locale for which you want to retrieve localized summaries. Defaults to the primary locale of the marketplace. | 

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


## SearchCatalogItems

> ItemSearchResults SearchCatalogItems(ctx).MarketplaceIds(marketplaceIds).Identifiers(identifiers).IdentifiersType(identifiersType).IncludedData(includedData).Locale(locale).SellerId(sellerId).Keywords(keywords).BrandNames(brandNames).ClassificationIds(classificationIds).PageSize(pageSize).PageToken(pageToken).KeywordsLocale(keywordsLocale).Execute()





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
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	identifiers := []string{"Inner_example"} // []string | A comma-delimited list of product identifiers that you can use to search the Amazon catalog. **Note:** You cannot include `identifiers` and `keywords` in the same request. (optional)
	identifiersType := "ASIN" // string | The type of product identifiers that you can use to search the Amazon catalog. **Note:** `identifiersType` is required when `identifiers` is in the request. (optional)
	includedData := []string{"IncludedData_example"} // []string | A comma-delimited list of datasets to include in the response. (optional) (default to ["summaries"])
	locale := "en_US" // string | The locale for which you want to retrieve localized summaries. Defaults to the primary locale of the marketplace. (optional)
	sellerId := "sellerId_example" // string | A selling partner identifier, such as a seller account or vendor code. **Note:** Required when `identifiersType` is `SKU`. (optional)
	keywords := []string{"Inner_example"} // []string | A comma-delimited list of keywords that you can use to search the Amazon catalog. **Note:** You cannot include `keywords` and `identifiers` in the same request. (optional)
	brandNames := []string{"Inner_example"} // []string | A comma-delimited list of brand names that you can use to limit the search in queries based on `keywords`. **Note:** Cannot be used with `identifiers`. (optional)
	classificationIds := []string{"Inner_example"} // []string | A comma-delimited list of classification identifiers that you can use to limit the search in queries based on `keywords`. **Note:** Cannot be used with `identifiers`. (optional)
	pageSize := int32(9) // int32 | The number of results to include on each page. (optional) (default to 10)
	pageToken := "sdlkj234lkj234lksjdflkjwdflkjsfdlkj234234234234" // string | A token that you can use to fetch a specific page when there are multiple pages of results. (optional)
	keywordsLocale := "en_US" // string | The language of the keywords that are included in queries based on `keywords`. Defaults to the primary locale of the marketplace. **Note:** Cannot be used with `identifiers`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.SearchCatalogItems(context.Background()).MarketplaceIds(marketplaceIds).Identifiers(identifiers).IdentifiersType(identifiersType).IncludedData(includedData).Locale(locale).SellerId(sellerId).Keywords(keywords).BrandNames(brandNames).ClassificationIds(classificationIds).PageSize(pageSize).PageToken(pageToken).KeywordsLocale(keywordsLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.SearchCatalogItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCatalogItems`: ItemSearchResults
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.SearchCatalogItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
 **identifiers** | **[]string** | A comma-delimited list of product identifiers that you can use to search the Amazon catalog. **Note:** You cannot include &#x60;identifiers&#x60; and &#x60;keywords&#x60; in the same request. | 
 **identifiersType** | **string** | The type of product identifiers that you can use to search the Amazon catalog. **Note:** &#x60;identifiersType&#x60; is required when &#x60;identifiers&#x60; is in the request. | 
 **includedData** | **[]string** | A comma-delimited list of datasets to include in the response. | [default to [&quot;summaries&quot;]]
 **locale** | **string** | The locale for which you want to retrieve localized summaries. Defaults to the primary locale of the marketplace. | 
 **sellerId** | **string** | A selling partner identifier, such as a seller account or vendor code. **Note:** Required when &#x60;identifiersType&#x60; is &#x60;SKU&#x60;. | 
 **keywords** | **[]string** | A comma-delimited list of keywords that you can use to search the Amazon catalog. **Note:** You cannot include &#x60;keywords&#x60; and &#x60;identifiers&#x60; in the same request. | 
 **brandNames** | **[]string** | A comma-delimited list of brand names that you can use to limit the search in queries based on &#x60;keywords&#x60;. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 
 **classificationIds** | **[]string** | A comma-delimited list of classification identifiers that you can use to limit the search in queries based on &#x60;keywords&#x60;. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 
 **pageSize** | **int32** | The number of results to include on each page. | [default to 10]
 **pageToken** | **string** | A token that you can use to fetch a specific page when there are multiple pages of results. | 
 **keywordsLocale** | **string** | The language of the keywords that are included in queries based on &#x60;keywords&#x60;. Defaults to the primary locale of the marketplace. **Note:** Cannot be used with &#x60;identifiers&#x60;. | 

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

