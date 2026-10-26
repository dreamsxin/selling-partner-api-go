# FulfillmentOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentOrderId** | **string** | The Fulfillment Order ID assigned by Amazon after fulfillment planning. This identifier is identical to the Shipment ID required by External Fulfillment APIs. | 

## Methods

### NewFulfillmentOrder

`func NewFulfillmentOrder(fulfillmentOrderId string, ) *FulfillmentOrder`

NewFulfillmentOrder instantiates a new FulfillmentOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentOrderWithDefaults

`func NewFulfillmentOrderWithDefaults() *FulfillmentOrder`

NewFulfillmentOrderWithDefaults instantiates a new FulfillmentOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentOrderId

`func (o *FulfillmentOrder) GetFulfillmentOrderId() string`

GetFulfillmentOrderId returns the FulfillmentOrderId field if non-nil, zero value otherwise.

### GetFulfillmentOrderIdOk

`func (o *FulfillmentOrder) GetFulfillmentOrderIdOk() (*string, bool)`

GetFulfillmentOrderIdOk returns a tuple with the FulfillmentOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentOrderId

`func (o *FulfillmentOrder) SetFulfillmentOrderId(v string)`

SetFulfillmentOrderId sets FulfillmentOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


