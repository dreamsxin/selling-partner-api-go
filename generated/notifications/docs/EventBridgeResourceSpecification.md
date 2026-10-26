# EventBridgeResourceSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The AWS region in which you will be receiving the notifications. | 
**AccountId** | **string** | The identifier for the AWS account that is responsible for charges related to receiving notifications. | 

## Methods

### NewEventBridgeResourceSpecification

`func NewEventBridgeResourceSpecification(region string, accountId string, ) *EventBridgeResourceSpecification`

NewEventBridgeResourceSpecification instantiates a new EventBridgeResourceSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBridgeResourceSpecificationWithDefaults

`func NewEventBridgeResourceSpecificationWithDefaults() *EventBridgeResourceSpecification`

NewEventBridgeResourceSpecificationWithDefaults instantiates a new EventBridgeResourceSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *EventBridgeResourceSpecification) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EventBridgeResourceSpecification) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EventBridgeResourceSpecification) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccountId

`func (o *EventBridgeResourceSpecification) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EventBridgeResourceSpecification) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EventBridgeResourceSpecification) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


