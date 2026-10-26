# CreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadVersion** | **string** | The version of the payload object to be used in the notification. | 
**DestinationId** | **string** | The identifier for the destination where notifications will be delivered. | 
**ProcessingDirective** | Pointer to [**ProcessingDirective**](ProcessingDirective.md) |  | [optional] 

## Methods

### NewCreateSubscriptionRequest

`func NewCreateSubscriptionRequest(payloadVersion string, destinationId string, ) *CreateSubscriptionRequest`

NewCreateSubscriptionRequest instantiates a new CreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestWithDefaults

`func NewCreateSubscriptionRequestWithDefaults() *CreateSubscriptionRequest`

NewCreateSubscriptionRequestWithDefaults instantiates a new CreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadVersion

`func (o *CreateSubscriptionRequest) GetPayloadVersion() string`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *CreateSubscriptionRequest) GetPayloadVersionOk() (*string, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *CreateSubscriptionRequest) SetPayloadVersion(v string)`

SetPayloadVersion sets PayloadVersion field to given value.


### GetDestinationId

`func (o *CreateSubscriptionRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateSubscriptionRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateSubscriptionRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetProcessingDirective

`func (o *CreateSubscriptionRequest) GetProcessingDirective() ProcessingDirective`

GetProcessingDirective returns the ProcessingDirective field if non-nil, zero value otherwise.

### GetProcessingDirectiveOk

`func (o *CreateSubscriptionRequest) GetProcessingDirectiveOk() (*ProcessingDirective, bool)`

GetProcessingDirectiveOk returns a tuple with the ProcessingDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDirective

`func (o *CreateSubscriptionRequest) SetProcessingDirective(v ProcessingDirective)`

SetProcessingDirective sets ProcessingDirective field to given value.

### HasProcessingDirective

`func (o *CreateSubscriptionRequest) HasProcessingDirective() bool`

HasProcessingDirective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


