# \ReportsAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](ReportsAPI.md#CancelReport) | **Delete** /reports/2021-06-30/reports/{reportId} | 
[**CancelReportSchedule**](ReportsAPI.md#CancelReportSchedule) | **Delete** /reports/2021-06-30/schedules/{reportScheduleId} | 
[**CreateReport**](ReportsAPI.md#CreateReport) | **Post** /reports/2021-06-30/reports | 
[**CreateReportSchedule**](ReportsAPI.md#CreateReportSchedule) | **Post** /reports/2021-06-30/schedules | 
[**GetReport**](ReportsAPI.md#GetReport) | **Get** /reports/2021-06-30/reports/{reportId} | 
[**GetReportDocument**](ReportsAPI.md#GetReportDocument) | **Get** /reports/2021-06-30/documents/{reportDocumentId} | 
[**GetReportSchedule**](ReportsAPI.md#GetReportSchedule) | **Get** /reports/2021-06-30/schedules/{reportScheduleId} | 
[**GetReportSchedules**](ReportsAPI.md#GetReportSchedules) | **Get** /reports/2021-06-30/schedules | 
[**GetReports**](ReportsAPI.md#GetReports) | **Get** /reports/2021-06-30/reports | 



## CancelReport

> CancelReport(ctx, reportId).Execute()





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
	reportId := "reportId_example" // string | The identifier for the report. This identifier is unique only in combination with a seller ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.CancelReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CancelReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The identifier for the report. This identifier is unique only in combination with a seller ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelReportSchedule

> CancelReportSchedule(ctx, reportScheduleId).Execute()





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
	reportScheduleId := "reportScheduleId_example" // string | The identifier for the report schedule. This identifier is unique only in combination with a seller ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.CancelReportSchedule(context.Background(), reportScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CancelReportSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportScheduleId** | **string** | The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReport

> CreateReportResponse CreateReport(ctx).Body(body).Execute()





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
	body := *openapiclient.NewCreateReportSpecification("ReportType_example", []string{"MarketplaceIds_example"}) // CreateReportSpecification | Information required to create the report.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.CreateReport(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CreateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReport`: CreateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.CreateReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateReportSpecification**](CreateReportSpecification.md) | Information required to create the report. | 

### Return type

[**CreateReportResponse**](CreateReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportSchedule

> CreateReportScheduleResponse CreateReportSchedule(ctx).Body(body).Execute()





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
	body := *openapiclient.NewCreateReportScheduleSpecification("ReportType_example", []string{"MarketplaceIds_example"}, "Period_example") // CreateReportScheduleSpecification | Information required to create the report schedule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.CreateReportSchedule(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CreateReportSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportSchedule`: CreateReportScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.CreateReportSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateReportScheduleSpecification**](CreateReportScheduleSpecification.md) | Information required to create the report schedule. | 

### Return type

[**CreateReportScheduleResponse**](CreateReportScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReport

> Report GetReport(ctx, reportId).Execute()





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
	reportId := "reportId_example" // string | The identifier for the report. This identifier is unique only in combination with a seller ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The identifier for the report. This identifier is unique only in combination with a seller ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportDocument

> ReportDocument GetReportDocument(ctx, reportDocumentId).EnableContentEncodingUrlHeader(enableContentEncodingUrlHeader).Execute()





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
	reportDocumentId := "reportDocumentId_example" // string | The identifier for the report document.
	enableContentEncodingUrlHeader := true // bool | When `true`, the Content-Encoding header on the returned URL is set to `gzip` instead of the default `identity` when `compressionAlgorithm` is `GZIP`. This allows automatic decompression by HTTP clients. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportDocument(context.Background(), reportDocumentId).EnableContentEncodingUrlHeader(enableContentEncodingUrlHeader).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportDocument`: ReportDocument
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportDocumentId** | **string** | The identifier for the report document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableContentEncodingUrlHeader** | **bool** | When &#x60;true&#x60;, the Content-Encoding header on the returned URL is set to &#x60;gzip&#x60; instead of the default &#x60;identity&#x60; when &#x60;compressionAlgorithm&#x60; is &#x60;GZIP&#x60;. This allows automatic decompression by HTTP clients. | 

### Return type

[**ReportDocument**](ReportDocument.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportSchedule

> ReportSchedule GetReportSchedule(ctx, reportScheduleId).Execute()





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
	reportScheduleId := "reportScheduleId_example" // string | The identifier for the report schedule. This identifier is unique only in combination with a seller ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportSchedule(context.Background(), reportScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportSchedule`: ReportSchedule
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportScheduleId** | **string** | The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportSchedule**](ReportSchedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportSchedules

> ReportScheduleList GetReportSchedules(ctx).ReportTypes(reportTypes).Execute()





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
	reportTypes := []string{"Inner_example"} // []string | A list of report types used to filter report schedules. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportSchedules(context.Background()).ReportTypes(reportTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportSchedules`: ReportScheduleList
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportTypes** | **[]string** | A list of report types used to filter report schedules. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | 

### Return type

[**ReportScheduleList**](ReportScheduleList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReports

> GetReportsResponse GetReports(ctx).ReportTypes(reportTypes).ProcessingStatuses(processingStatuses).MarketplaceIds(marketplaceIds).PageSize(pageSize).CreatedSince(createdSince).CreatedUntil(createdUntil).NextToken(nextToken).Execute()





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
	reportTypes := []string{"Inner_example"} // []string | A list of report types used to filter reports. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required. (optional)
	processingStatuses := []string{"ProcessingStatuses_example"} // []string | A list of processing statuses used to filter reports. (optional)
	marketplaceIds := []string{"Inner_example"} // []string | A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify. (optional)
	pageSize := int32(56) // int32 | The maximum number of reports to return in a single call. (optional) (default to 10)
	createdSince := time.Now() // time.Time | The earliest report creation date and time for reports to include in the response, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days. (optional)
	createdUntil := time.Now() // time.Time | The latest report creation date and time for reports to include in the response, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. The default is now. (optional)
	nextToken := "nextToken_example" // string | A string token returned in the response to your previous request. `nextToken` is returned when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the `getReports` operation and include this token as the only parameter. Specifying `nextToken` with any other parameters will cause the request to fail. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReports(context.Background()).ReportTypes(reportTypes).ProcessingStatuses(processingStatuses).MarketplaceIds(marketplaceIds).PageSize(pageSize).CreatedSince(createdSince).CreatedUntil(createdUntil).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReports`: GetReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportTypes** | **[]string** | A list of report types used to filter reports. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required. | 
 **processingStatuses** | **[]string** | A list of processing statuses used to filter reports. | 
 **marketplaceIds** | **[]string** | A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify. | 
 **pageSize** | **int32** | The maximum number of reports to return in a single call. | [default to 10]
 **createdSince** | **time.Time** | The earliest report creation date and time for reports to include in the response, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days. | 
 **createdUntil** | **time.Time** | The latest report creation date and time for reports to include in the response, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. The default is now. | 
 **nextToken** | **string** | A string token returned in the response to your previous request. &#x60;nextToken&#x60; is returned when the number of results exceeds the specified &#x60;pageSize&#x60; value. To get the next page of results, call the &#x60;getReports&#x60; operation and include this token as the only parameter. Specifying &#x60;nextToken&#x60; with any other parameters will cause the request to fail. | 

### Return type

[**GetReportsResponse**](GetReportsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

