# CreateSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | A list of error responses returned when a request is unsuccessful. | [optional] 

## Methods

### NewCreateSubscriptionResponse

`func NewCreateSubscriptionResponse() *CreateSubscriptionResponse`

NewCreateSubscriptionResponse instantiates a new CreateSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionResponseWithDefaults

`func NewCreateSubscriptionResponseWithDefaults() *CreateSubscriptionResponse`

NewCreateSubscriptionResponseWithDefaults instantiates a new CreateSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *CreateSubscriptionResponse) GetPayload() Subscription`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateSubscriptionResponse) GetPayloadOk() (*Subscription, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateSubscriptionResponse) SetPayload(v Subscription)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CreateSubscriptionResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetErrors

`func (o *CreateSubscriptionResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateSubscriptionResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateSubscriptionResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateSubscriptionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


