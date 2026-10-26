# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The subscription identifier generated when the subscription is created. | 
**PayloadVersion** | **string** | The version of the payload object to be used in the notification. | 
**DestinationId** | **string** | The identifier for the destination where notifications will be delivered. | 
**ProcessingDirective** | Pointer to [**ProcessingDirective**](ProcessingDirective.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription(subscriptionId string, payloadVersion string, destinationId string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *Subscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Subscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Subscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPayloadVersion

`func (o *Subscription) GetPayloadVersion() string`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *Subscription) GetPayloadVersionOk() (*string, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *Subscription) SetPayloadVersion(v string)`

SetPayloadVersion sets PayloadVersion field to given value.


### GetDestinationId

`func (o *Subscription) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *Subscription) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *Subscription) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetProcessingDirective

`func (o *Subscription) GetProcessingDirective() ProcessingDirective`

GetProcessingDirective returns the ProcessingDirective field if non-nil, zero value otherwise.

### GetProcessingDirectiveOk

`func (o *Subscription) GetProcessingDirectiveOk() (*ProcessingDirective, bool)`

GetProcessingDirectiveOk returns a tuple with the ProcessingDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDirective

`func (o *Subscription) SetProcessingDirective(v ProcessingDirective)`

SetProcessingDirective sets ProcessingDirective field to given value.

### HasProcessingDirective

`func (o *Subscription) HasProcessingDirective() bool`

HasProcessingDirective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


