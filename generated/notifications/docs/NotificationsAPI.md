# \NotificationsAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestination**](NotificationsAPI.md#CreateDestination) | **Post** /notifications/v1/destinations | 
[**CreateSubscription**](NotificationsAPI.md#CreateSubscription) | **Post** /notifications/v1/subscriptions/{notificationType} | 
[**DeleteDestination**](NotificationsAPI.md#DeleteDestination) | **Delete** /notifications/v1/destinations/{destinationId} | 
[**DeleteSubscriptionById**](NotificationsAPI.md#DeleteSubscriptionById) | **Delete** /notifications/v1/subscriptions/{notificationType}/{subscriptionId} | 
[**GetDestination**](NotificationsAPI.md#GetDestination) | **Get** /notifications/v1/destinations/{destinationId} | 
[**GetDestinations**](NotificationsAPI.md#GetDestinations) | **Get** /notifications/v1/destinations | 
[**GetSubscription**](NotificationsAPI.md#GetSubscription) | **Get** /notifications/v1/subscriptions/{notificationType} | 
[**GetSubscriptionById**](NotificationsAPI.md#GetSubscriptionById) | **Get** /notifications/v1/subscriptions/{notificationType}/{subscriptionId} | 
[**GetSubscriptions**](NotificationsAPI.md#GetSubscriptions) | **Get** /notifications/v1/subscriptions | 
[**SendTestNotification**](NotificationsAPI.md#SendTestNotification) | **Post** /notifications/v1/subscriptions/{notificationType}/testNotification | 



## CreateDestination

> CreateDestinationResponse CreateDestination(ctx).Body(body).Execute()





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
	body := *openapiclient.NewCreateDestinationRequest(*openapiclient.NewDestinationResourceSpecification(), "Name_example") // CreateDestinationRequest | The request schema for the `createDestination` operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateDestination(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestination`: CreateDestinationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDestinationRequest**](CreateDestinationRequest.md) | The request schema for the &#x60;createDestination&#x60; operation. | 

### Return type

[**CreateDestinationResponse**](CreateDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> CreateSubscriptionResponse CreateSubscription(ctx, notificationType).Body(body).Execute()





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
	notificationType := "notificationType_example" // string | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).
	body := *openapiclient.NewCreateSubscriptionRequest("PayloadVersion_example", "DestinationId_example") // CreateSubscriptionRequest | The request schema for the `createSubscription` operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateSubscription(context.Background(), notificationType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: CreateSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md) | The request schema for the &#x60;createSubscription&#x60; operation. | 

### Return type

[**CreateSubscriptionResponse**](CreateSubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestination

> DeleteDestinationResponse DeleteDestination(ctx, destinationId).Execute()





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
	destinationId := "destinationId_example" // string | The identifier for the destination that you want to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteDestination(context.Background(), destinationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDestination`: DeleteDestinationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** | The identifier for the destination that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDestinationResponse**](DeleteDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionById

> DeleteSubscriptionByIdResponse DeleteSubscriptionById(ctx, subscriptionId, notificationType).Execute()





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
	subscriptionId := "subscriptionId_example" // string | The identifier for the subscription that you want to delete.
	notificationType := "notificationType_example" // string | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteSubscriptionById(context.Background(), subscriptionId, notificationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteSubscriptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubscriptionById`: DeleteSubscriptionByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The identifier for the subscription that you want to delete. | 
**notificationType** | **string** | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteSubscriptionByIdResponse**](DeleteSubscriptionByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Operation Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestination

> GetDestinationResponse GetDestination(ctx, destinationId).Execute()





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
	destinationId := "destinationId_example" // string | The identifier generated when you created the destination.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetDestination(context.Background(), destinationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestination`: GetDestinationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** | The identifier generated when you created the destination. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDestinationResponse**](GetDestinationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinations

> GetDestinationsResponse GetDestinations(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetDestinations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinations`: GetDestinationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetDestinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsRequest struct via the builder pattern


### Return type

[**GetDestinationsResponse**](GetDestinationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> GetSubscriptionResponse GetSubscription(ctx, notificationType).PayloadVersion(payloadVersion).Execute()





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
	notificationType := "notificationType_example" // string | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).
	payloadVersion := "payloadVersion_example" // string | The version of the payload object to be used in the notification. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetSubscription(context.Background(), notificationType).PayloadVersion(payloadVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: GetSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payloadVersion** | **string** | The version of the payload object to be used in the notification. | 

### Return type

[**GetSubscriptionResponse**](GetSubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionById

> GetSubscriptionByIdResponse GetSubscriptionById(ctx, subscriptionId, notificationType).Execute()





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
	subscriptionId := "subscriptionId_example" // string | The identifier for the subscription that you want to get.
	notificationType := "notificationType_example" // string | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetSubscriptionById(context.Background(), subscriptionId, notificationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSubscriptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionById`: GetSubscriptionByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The identifier for the subscription that you want to get. | 
**notificationType** | **string** | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSubscriptionByIdResponse**](GetSubscriptionByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, Successful Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> GetSubscriptionsResponse GetSubscriptions(ctx).NotificationTypes(notificationTypes).PayloadVersion(payloadVersion).PageSize(pageSize).NextToken(nextToken).Execute()





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
	notificationTypes := []string{"Inner_example"} // []string | A list of notification types to retrieve subscriptions for. Currently limited to a single notification type per request.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).
	payloadVersion := "payloadVersion_example" // string | The version of the payload object to be used in the notification. (optional)
	pageSize := int32(56) // int32 | The maximum number of subscriptions to return per page. Minimum value is 30. Maximum value is 100. Default is 30. (optional) (default to 30)
	nextToken := "nextToken_example" // string | A token to retrieve the next page of results. If this field is not empty in a response, pass its value in the next request to retrieve the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetSubscriptions(context.Background()).NotificationTypes(notificationTypes).PayloadVersion(payloadVersion).PageSize(pageSize).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: GetSubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTypes** | **[]string** | A list of notification types to retrieve subscriptions for. Currently limited to a single notification type per request.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 
 **payloadVersion** | **string** | The version of the payload object to be used in the notification. | 
 **pageSize** | **int32** | The maximum number of subscriptions to return per page. Minimum value is 30. Maximum value is 100. Default is 30. | [default to 30]
 **nextToken** | **string** | A token to retrieve the next page of results. If this field is not empty in a response, pass its value in the next request to retrieve the next page. | 

### Return type

[**GetSubscriptionsResponse**](GetSubscriptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestNotification

> SendTestNotificationResponse SendTestNotification(ctx, notificationType).Body(body).Execute()





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
	notificationType := "notificationType_example" // string | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide).
	body := *openapiclient.NewSendTestNotificationRequest() // SendTestNotificationRequest | The request schema for the `sendTestNotification` operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.SendTestNotification(context.Background(), notificationType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.SendTestNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestNotification`: SendTestNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.SendTestNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** | The type of notification.   For more information about notification types, refer to the [Notifications API v1 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/notifications-api-v1-use-case-guide). | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SendTestNotificationRequest**](SendTestNotificationRequest.md) | The request schema for the &#x60;sendTestNotification&#x60; operation. | 

### Return type

[**SendTestNotificationResponse**](SendTestNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, Successful Response, Response

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

