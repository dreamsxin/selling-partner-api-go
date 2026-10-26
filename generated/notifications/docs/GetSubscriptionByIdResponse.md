# GetSubscriptionByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | A list of error responses returned when a request is unsuccessful. | [optional] 

## Methods

### NewGetSubscriptionByIdResponse

`func NewGetSubscriptionByIdResponse() *GetSubscriptionByIdResponse`

NewGetSubscriptionByIdResponse instantiates a new GetSubscriptionByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscriptionByIdResponseWithDefaults

`func NewGetSubscriptionByIdResponseWithDefaults() *GetSubscriptionByIdResponse`

NewGetSubscriptionByIdResponseWithDefaults instantiates a new GetSubscriptionByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *GetSubscriptionByIdResponse) GetPayload() Subscription`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetSubscriptionByIdResponse) GetPayloadOk() (*Subscription, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetSubscriptionByIdResponse) SetPayload(v Subscription)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GetSubscriptionByIdResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetErrors

`func (o *GetSubscriptionByIdResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetSubscriptionByIdResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetSubscriptionByIdResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetSubscriptionByIdResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


