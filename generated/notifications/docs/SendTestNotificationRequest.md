# SendTestNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | Pointer to **string** | The identifier for the destination where notifications will be delivered. | [optional] 
**TestNotification** | Pointer to [**TestNotification**](TestNotification.md) |  | [optional] 

## Methods

### NewSendTestNotificationRequest

`func NewSendTestNotificationRequest() *SendTestNotificationRequest`

NewSendTestNotificationRequest instantiates a new SendTestNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTestNotificationRequestWithDefaults

`func NewSendTestNotificationRequestWithDefaults() *SendTestNotificationRequest`

NewSendTestNotificationRequestWithDefaults instantiates a new SendTestNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *SendTestNotificationRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *SendTestNotificationRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *SendTestNotificationRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *SendTestNotificationRequest) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetTestNotification

`func (o *SendTestNotificationRequest) GetTestNotification() TestNotification`

GetTestNotification returns the TestNotification field if non-nil, zero value otherwise.

### GetTestNotificationOk

`func (o *SendTestNotificationRequest) GetTestNotificationOk() (*TestNotification, bool)`

GetTestNotificationOk returns a tuple with the TestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestNotification

`func (o *SendTestNotificationRequest) SetTestNotification(v TestNotification)`

SetTestNotification sets TestNotification field to given value.

### HasTestNotification

`func (o *SendTestNotificationRequest) HasTestNotification() bool`

HasTestNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


