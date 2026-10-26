# DestinationResourceSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sqs** | Pointer to [**SqsResource**](SqsResource.md) |  | [optional] 
**EventBridge** | Pointer to [**EventBridgeResourceSpecification**](EventBridgeResourceSpecification.md) |  | [optional] 

## Methods

### NewDestinationResourceSpecification

`func NewDestinationResourceSpecification() *DestinationResourceSpecification`

NewDestinationResourceSpecification instantiates a new DestinationResourceSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationResourceSpecificationWithDefaults

`func NewDestinationResourceSpecificationWithDefaults() *DestinationResourceSpecification`

NewDestinationResourceSpecificationWithDefaults instantiates a new DestinationResourceSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqs

`func (o *DestinationResourceSpecification) GetSqs() SqsResource`

GetSqs returns the Sqs field if non-nil, zero value otherwise.

### GetSqsOk

`func (o *DestinationResourceSpecification) GetSqsOk() (*SqsResource, bool)`

GetSqsOk returns a tuple with the Sqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqs

`func (o *DestinationResourceSpecification) SetSqs(v SqsResource)`

SetSqs sets Sqs field to given value.

### HasSqs

`func (o *DestinationResourceSpecification) HasSqs() bool`

HasSqs returns a boolean if a field has been set.

### GetEventBridge

`func (o *DestinationResourceSpecification) GetEventBridge() EventBridgeResourceSpecification`

GetEventBridge returns the EventBridge field if non-nil, zero value otherwise.

### GetEventBridgeOk

`func (o *DestinationResourceSpecification) GetEventBridgeOk() (*EventBridgeResourceSpecification, bool)`

GetEventBridgeOk returns a tuple with the EventBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBridge

`func (o *DestinationResourceSpecification) SetEventBridge(v EventBridgeResourceSpecification)`

SetEventBridge sets EventBridge field to given value.

### HasEventBridge

`func (o *DestinationResourceSpecification) HasEventBridge() bool`

HasEventBridge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


