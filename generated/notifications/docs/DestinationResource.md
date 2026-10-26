# DestinationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sqs** | Pointer to [**SqsResource**](SqsResource.md) |  | [optional] 
**EventBridge** | Pointer to [**EventBridgeResource**](EventBridgeResource.md) |  | [optional] 

## Methods

### NewDestinationResource

`func NewDestinationResource() *DestinationResource`

NewDestinationResource instantiates a new DestinationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationResourceWithDefaults

`func NewDestinationResourceWithDefaults() *DestinationResource`

NewDestinationResourceWithDefaults instantiates a new DestinationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqs

`func (o *DestinationResource) GetSqs() SqsResource`

GetSqs returns the Sqs field if non-nil, zero value otherwise.

### GetSqsOk

`func (o *DestinationResource) GetSqsOk() (*SqsResource, bool)`

GetSqsOk returns a tuple with the Sqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqs

`func (o *DestinationResource) SetSqs(v SqsResource)`

SetSqs sets Sqs field to given value.

### HasSqs

`func (o *DestinationResource) HasSqs() bool`

HasSqs returns a boolean if a field has been set.

### GetEventBridge

`func (o *DestinationResource) GetEventBridge() EventBridgeResource`

GetEventBridge returns the EventBridge field if non-nil, zero value otherwise.

### GetEventBridgeOk

`func (o *DestinationResource) GetEventBridgeOk() (*EventBridgeResource, bool)`

GetEventBridgeOk returns a tuple with the EventBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBridge

`func (o *DestinationResource) SetEventBridge(v EventBridgeResource)`

SetEventBridge sets EventBridge field to given value.

### HasEventBridge

`func (o *DestinationResource) HasEventBridge() bool`

HasEventBridge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


