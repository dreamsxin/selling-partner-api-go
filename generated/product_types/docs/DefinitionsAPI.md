# \DefinitionsAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefinitionsProductType**](DefinitionsAPI.md#GetDefinitionsProductType) | **Get** /definitions/2020-09-01/productTypes/{productType} | 
[**SearchDefinitionsProductTypes**](DefinitionsAPI.md#SearchDefinitionsProductTypes) | **Get** /definitions/2020-09-01/productTypes | 



## GetDefinitionsProductType

> ProductTypeDefinition GetDefinitionsProductType(ctx, productType).MarketplaceIds(marketplaceIds).SellerId(sellerId).ProductTypeVersion(productTypeVersion).Requirements(requirements).RequirementsEnforced(requirementsEnforced).Locale(locale).ParentageLevel(parentageLevel).Execute()





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
	productType := "LUGGAGE" // string | The Amazon product type name.
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request. Note: This parameter is limited to one marketplaceId at this time.
	sellerId := "sellerId_example" // string | A selling partner identifier. When provided, seller-specific requirements and values are populated within the product type definition schema, such as brand names associated with the selling partner. (optional)
	productTypeVersion := "LATEST" // string | The version of the Amazon product type to retrieve. Defaults to \"LATEST\". Prerelease versions of product type definitions may be retrieved with \"RELEASE_CANDIDATE\". If no prerelease version is currently available, the \"LATEST\" live version will be provided. (optional) (default to "LATEST")
	requirements := "LISTING" // string | The name of the requirements set to retrieve requirements for. (optional) (default to "LISTING")
	requirementsEnforced := "ENFORCED" // string | Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all the required attributes being present (such as for partial updates). (optional) (default to "ENFORCED")
	locale := "DEFAULT" // string | Locale for retrieving display labels and other presentation details. Defaults to the default language of the first marketplace in the request. (optional) (default to "DEFAULT")
	parentageLevel := "CHILD" // string | The parentage level of the listing to retrieve a schema for. When provided, the schema is simplified by resolving all conditional logic related to the specified parentage level, resulting in a smaller schema with fewer conditions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.GetDefinitionsProductType(context.Background(), productType).MarketplaceIds(marketplaceIds).SellerId(sellerId).ProductTypeVersion(productTypeVersion).Requirements(requirements).RequirementsEnforced(requirementsEnforced).Locale(locale).ParentageLevel(parentageLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.GetDefinitionsProductType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefinitionsProductType`: ProductTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.GetDefinitionsProductType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productType** | **string** | The Amazon product type name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitionsProductTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. Note: This parameter is limited to one marketplaceId at this time. | 
 **sellerId** | **string** | A selling partner identifier. When provided, seller-specific requirements and values are populated within the product type definition schema, such as brand names associated with the selling partner. | 
 **productTypeVersion** | **string** | The version of the Amazon product type to retrieve. Defaults to \&quot;LATEST\&quot;. Prerelease versions of product type definitions may be retrieved with \&quot;RELEASE_CANDIDATE\&quot;. If no prerelease version is currently available, the \&quot;LATEST\&quot; live version will be provided. | [default to &quot;LATEST&quot;]
 **requirements** | **string** | The name of the requirements set to retrieve requirements for. | [default to &quot;LISTING&quot;]
 **requirementsEnforced** | **string** | Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all the required attributes being present (such as for partial updates). | [default to &quot;ENFORCED&quot;]
 **locale** | **string** | Locale for retrieving display labels and other presentation details. Defaults to the default language of the first marketplace in the request. | [default to &quot;DEFAULT&quot;]
 **parentageLevel** | **string** | The parentage level of the listing to retrieve a schema for. When provided, the schema is simplified by resolving all conditional logic related to the specified parentage level, resulting in a smaller schema with fewer conditions. | 

### Return type

[**ProductTypeDefinition**](ProductTypeDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDefinitionsProductTypes

> ProductTypeList SearchDefinitionsProductTypes(ctx).MarketplaceIds(marketplaceIds).Keywords(keywords).ItemName(itemName).Locale(locale).SearchLocale(searchLocale).Execute()





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
	marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request.
	keywords := []string{"Inner_example"} // []string | A comma-delimited list of keywords to search product types. **Note:** Cannot be used with `itemName`. (optional)
	itemName := "Running shoes" // string | Title of ASIN to get product type recommendation. **Note:** Cannot be used with `keywords`. (optional)
	locale := "en_US" // string | Locale for display names in response. Defaults to primary locale of the marketplace. (optional)
	searchLocale := "en_US" // string | Language used for `keywords` or `itemName` parameters. Defaults to primary locale of the marketplace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionsAPI.SearchDefinitionsProductTypes(context.Background()).MarketplaceIds(marketplaceIds).Keywords(keywords).ItemName(itemName).Locale(locale).SearchLocale(searchLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionsAPI.SearchDefinitionsProductTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDefinitionsProductTypes`: ProductTypeList
	fmt.Fprintf(os.Stdout, "Response from `DefinitionsAPI.SearchDefinitionsProductTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDefinitionsProductTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **keywords** | **[]string** | A comma-delimited list of keywords to search product types. **Note:** Cannot be used with &#x60;itemName&#x60;. | 
 **itemName** | **string** | Title of ASIN to get product type recommendation. **Note:** Cannot be used with &#x60;keywords&#x60;. | 
 **locale** | **string** | Locale for display names in response. Defaults to primary locale of the marketplace. | 
 **searchLocale** | **string** | Language used for &#x60;keywords&#x60; or &#x60;itemName&#x60; parameters. Defaults to primary locale of the marketplace. | 

### Return type

[**ProductTypeList**](ProductTypeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

