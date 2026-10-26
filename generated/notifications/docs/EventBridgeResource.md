# EventBridgeResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the partner event source associated with the destination. | 
**Region** | **string** | The AWS region in which you receive the notifications. For AWS regions that Amazon EventBridge supports, refer to [Amazon EventBridge endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/ev.html). | 
**AccountId** | **string** | The identifier for the AWS account that is responsible for charges related to receiving notifications. | 

## Methods

### NewEventBridgeResource

`func NewEventBridgeResource(name string, region string, accountId string, ) *EventBridgeResource`

NewEventBridgeResource instantiates a new EventBridgeResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBridgeResourceWithDefaults

`func NewEventBridgeResourceWithDefaults() *EventBridgeResource`

NewEventBridgeResourceWithDefaults instantiates a new EventBridgeResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventBridgeResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventBridgeResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventBridgeResource) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *EventBridgeResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EventBridgeResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EventBridgeResource) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccountId

`func (o *EventBridgeResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EventBridgeResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EventBridgeResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


