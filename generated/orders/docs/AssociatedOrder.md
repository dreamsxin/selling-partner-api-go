# AssociatedOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | The unique identifier of the related order that is associated with the current order. | [optional] 
**AssociationType** | Pointer to **string** | The relationship between the current order and the associated order.  **Possible values**: &#x60;REPLACEMENT_ORIGINAL_ID&#x60;, &#x60;EXCHANGE_ORIGINAL_ID&#x60; | [optional] 

## Methods

### NewAssociatedOrder

`func NewAssociatedOrder() *AssociatedOrder`

NewAssociatedOrder instantiates a new AssociatedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociatedOrderWithDefaults

`func NewAssociatedOrderWithDefaults() *AssociatedOrder`

NewAssociatedOrderWithDefaults instantiates a new AssociatedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *AssociatedOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AssociatedOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AssociatedOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AssociatedOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAssociationType

`func (o *AssociatedOrder) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *AssociatedOrder) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *AssociatedOrder) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *AssociatedOrder) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


