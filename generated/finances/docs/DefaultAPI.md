# \DefaultAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBalances**](DefaultAPI.md#ListBalances) | **Get** /finances/2024-06-19/balances | 
[**ListSummary**](DefaultAPI.md#ListSummary) | **Get** /finances/2024-06-19/summary | 
[**ListTransactions**](DefaultAPI.md#ListTransactions) | **Get** /finances/2024-06-19/transactions | 



## ListBalances

> ListBalancesResponse ListBalances(ctx).MarketplaceIds(marketplaceIds).BalanceType(balanceType).AccountType(accountType).AsOfDate(asOfDate).NextToken(nextToken).Execute()





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
	marketplaceIds := []string{"Inner_example"} // []string | The marketplaces from which to retrieve balances. If omitted, balances from all applicable marketplaces may be returned. To find the marketplace ID for a region, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). (optional)
	balanceType := "TOTAL" // string | The type of balance to include in the response. If omitted, all balance types may be included in the response.  **Possible values:** `AVAILABLE`, `RESERVED`, `TOTAL`, `DEFERRED`, `ACCOUNT_LEVEL_RESERVE` (optional)
	accountType := "Standard Orders" // string | The type of account to include in the response. (optional)
	asOfDate := time.Now() // string | The date from which you want to retrieve balances. If provided, the response includes historical balances at the specified date. The value must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. If omitted, the point in time balance is provided. (optional)
	nextToken := "jehgri34yo7jr9e8f984tr9i4o" // string | A token that you use to retrieve subsequent pages of results. When there are more than 500 results available, the response will include a `nextToken` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. Repeat this process until the `nextToken` value is null to retrieve all results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListBalances(context.Background()).MarketplaceIds(marketplaceIds).BalanceType(balanceType).AccountType(accountType).AsOfDate(asOfDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBalances`: ListBalancesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceIds** | **[]string** | The marketplaces from which to retrieve balances. If omitted, balances from all applicable marketplaces may be returned. To find the marketplace ID for a region, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
 **balanceType** | **string** | The type of balance to include in the response. If omitted, all balance types may be included in the response.  **Possible values:** &#x60;AVAILABLE&#x60;, &#x60;RESERVED&#x60;, &#x60;TOTAL&#x60;, &#x60;DEFERRED&#x60;, &#x60;ACCOUNT_LEVEL_RESERVE&#x60; | 
 **accountType** | **string** | The type of account to include in the response. | 
 **asOfDate** | **string** | The date from which you want to retrieve balances. If provided, the response includes historical balances at the specified date. The value must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. If omitted, the point in time balance is provided. | 
 **nextToken** | **string** | A token that you use to retrieve subsequent pages of results. When there are more than 500 results available, the response will include a &#x60;nextToken&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. Repeat this process until the &#x60;nextToken&#x60; value is null to retrieve all results. | 

### Return type

[**ListBalancesResponse**](ListBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSummary

> SummaryResponse ListSummary(ctx).MarketplaceIds(marketplaceIds).AccountType(accountType).RelatedIdentifierName(relatedIdentifierName).RelatedIdentifierValue(relatedIdentifierValue).PeriodStart(periodStart).PeriodEnd(periodEnd).NextToken(nextToken).Execute()





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
	marketplaceIds := []string{"Inner_example"} // []string | The marketplaces from which to retrieve summaries. If omitted, summaries from all applicable marketplaces may be returned. To find the marketplace ID for a region, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). (optional)
	accountType := "Standard Orders" // string | The type of account to include in the response. (optional)
	relatedIdentifierName := "SETTLEMENT_ID" // string | The name of the `relatedIdentifier`. The only possible value is `SETTLEMENT_ID`, the settlement ID associated with the summary. (optional)
	relatedIdentifierValue := "8129762527551" // string | The value of the `relatedIdentifier`. (optional)
	periodStart := time.Now() // string | The start of the period for which to retrieve summaries. When provided, the response will only include summaries with transactions that occurred on or after the specified date. The value must be formatted in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. (optional)
	periodEnd := time.Now() // string | The end of the period for which to retrieve summaries. When provided, the response will only include summaries with transactions that occurred on or before the specified date. The value must be formatted in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. (optional)
	nextToken := "jehgri34yo7jr9e8f984tr9i4o" // string | A token that you use to retrieve subsequent pages of results. When there are more results available, the response will include a `nextToken` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. Repeat this process until the `nextToken` value is null to retrieve all results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListSummary(context.Background()).MarketplaceIds(marketplaceIds).AccountType(accountType).RelatedIdentifierName(relatedIdentifierName).RelatedIdentifierValue(relatedIdentifierValue).PeriodStart(periodStart).PeriodEnd(periodEnd).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSummary`: SummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceIds** | **[]string** | The marketplaces from which to retrieve summaries. If omitted, summaries from all applicable marketplaces may be returned. To find the marketplace ID for a region, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
 **accountType** | **string** | The type of account to include in the response. | 
 **relatedIdentifierName** | **string** | The name of the &#x60;relatedIdentifier&#x60;. The only possible value is &#x60;SETTLEMENT_ID&#x60;, the settlement ID associated with the summary. | 
 **relatedIdentifierValue** | **string** | The value of the &#x60;relatedIdentifier&#x60;. | 
 **periodStart** | **string** | The start of the period for which to retrieve summaries. When provided, the response will only include summaries with transactions that occurred on or after the specified date. The value must be formatted in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. | 
 **periodEnd** | **string** | The end of the period for which to retrieve summaries. When provided, the response will only include summaries with transactions that occurred on or before the specified date. The value must be formatted in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date format. | 
 **nextToken** | **string** | A token that you use to retrieve subsequent pages of results. When there are more results available, the response will include a &#x60;nextToken&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. Repeat this process until the &#x60;nextToken&#x60; value is null to retrieve all results. | 

### Return type

[**SummaryResponse**](SummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactionsResponse ListTransactions(ctx).PostedAfter(postedAfter).PostedBefore(postedBefore).MarketplaceId(marketplaceId).TransactionStatus(transactionStatus).RelatedIdentifierName(relatedIdentifierName).RelatedIdentifierValue(relatedIdentifierValue).NextToken(nextToken).Execute()





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
	postedAfter := time.Now() // time.Time | The response includes financial events posted on or after this date. This date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. The date-time must be more than two minutes before the time of the request.  This field is required if you do not specify a related identifier. (optional)
	postedBefore := time.Now() // time.Time | The response includes financial events posted before (but not on) this date. This date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.  The date-time must be later than `PostedAfter` and more than two minutes before the request was submitted. If `PostedAfter` and `PostedBefore` are more than 180 days apart, the response is empty.  **Default:** Two minutes before the time of the request. (optional)
	marketplaceId := "ATIV93840DER" // string | The identifier of the marketplace from which you want to retrieve transactions. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for a marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). (optional)
	transactionStatus := "RELEASED" // string | The status of the transaction.  **Possible values:**  * `DEFERRED`: the transaction is currently deferred. * `RELEASED`: the transaction is currently released. * `DEFERRED_RELEASED`: the transaction was deferred in the past, but is now released. The status of a deferred transaction is updated to `DEFERRED_RELEASED` when the transaction is released. (optional)
	relatedIdentifierName := "FINANCIAL_EVENT_GROUP_ID" // string | The name of the `relatedIdentifier`.  **Possible values:**  * `FINANCIAL_EVENT_GROUP_ID`: the financial event group ID associated with the transaction.  * `ORDER_ID`: the order ID associated with the transaction.    **Note:**  `FINANCIAL_EVENT_GROUP_ID` and `ORDER_ID` are the only `relatedIdentifier` with filtering capabilities at the moment. While other `relatedIdentifier` values will be included in the response when available, they cannot be used for filtering purposes. (optional)
	relatedIdentifierValue := "8129762527551" // string | The value of the `relatedIdentifier`. (optional)
	nextToken := "jehgri34yo7jr9e8f984tr9i4o" // string | The response includes `nextToken` when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTransactions(context.Background()).PostedAfter(postedAfter).PostedBefore(postedBefore).MarketplaceId(marketplaceId).TransactionStatus(transactionStatus).RelatedIdentifierName(relatedIdentifierName).RelatedIdentifierValue(relatedIdentifierValue).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: ListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postedAfter** | **time.Time** | The response includes financial events posted on or after this date. This date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. The date-time must be more than two minutes before the time of the request.  This field is required if you do not specify a related identifier. | 
 **postedBefore** | **time.Time** | The response includes financial events posted before (but not on) this date. This date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.  The date-time must be later than &#x60;PostedAfter&#x60; and more than two minutes before the request was submitted. If &#x60;PostedAfter&#x60; and &#x60;PostedBefore&#x60; are more than 180 days apart, the response is empty.  **Default:** Two minutes before the time of the request. | 
 **marketplaceId** | **string** | The identifier of the marketplace from which you want to retrieve transactions. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for a marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
 **transactionStatus** | **string** | The status of the transaction.  **Possible values:**  * &#x60;DEFERRED&#x60;: the transaction is currently deferred. * &#x60;RELEASED&#x60;: the transaction is currently released. * &#x60;DEFERRED_RELEASED&#x60;: the transaction was deferred in the past, but is now released. The status of a deferred transaction is updated to &#x60;DEFERRED_RELEASED&#x60; when the transaction is released. | 
 **relatedIdentifierName** | **string** | The name of the &#x60;relatedIdentifier&#x60;.  **Possible values:**  * &#x60;FINANCIAL_EVENT_GROUP_ID&#x60;: the financial event group ID associated with the transaction.  * &#x60;ORDER_ID&#x60;: the order ID associated with the transaction.    **Note:**  &#x60;FINANCIAL_EVENT_GROUP_ID&#x60; and &#x60;ORDER_ID&#x60; are the only &#x60;relatedIdentifier&#x60; with filtering capabilities at the moment. While other &#x60;relatedIdentifier&#x60; values will be included in the response when available, they cannot be used for filtering purposes. | 
 **relatedIdentifierValue** | **string** | The value of the &#x60;relatedIdentifier&#x60;. | 
 **nextToken** | **string** | The response includes &#x60;nextToken&#x60; when the number of results exceeds the specified &#x60;pageSize&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;nextToken&#x60; is null. Note that this operation can return empty pages. | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

