# AssociatedOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | The order identifier of the associated order item. | [optional] 
**OrderItemId** | Pointer to **string** | The order item identifier of the associated order item. | [optional] 
**AssociationType** | Pointer to **string** | The type of association between the order items.  **Possible values**: - &#x60;VALUE_ADD_SERVICE&#x60; (The associated item is a service order) - &#x60;WARRANTY&#x60; (The associated item is a product covered by a warranty or protection plan) | [optional] 
**Asin** | Pointer to **string** | The ASIN of the associated order item. | [optional] 
**PurchasePrice** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewAssociatedOrderItem

`func NewAssociatedOrderItem() *AssociatedOrderItem`

NewAssociatedOrderItem instantiates a new AssociatedOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociatedOrderItemWithDefaults

`func NewAssociatedOrderItemWithDefaults() *AssociatedOrderItem`

NewAssociatedOrderItemWithDefaults instantiates a new AssociatedOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *AssociatedOrderItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AssociatedOrderItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AssociatedOrderItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AssociatedOrderItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderItemId

`func (o *AssociatedOrderItem) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *AssociatedOrderItem) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *AssociatedOrderItem) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.

### HasOrderItemId

`func (o *AssociatedOrderItem) HasOrderItemId() bool`

HasOrderItemId returns a boolean if a field has been set.

### GetAssociationType

`func (o *AssociatedOrderItem) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *AssociatedOrderItem) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *AssociatedOrderItem) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *AssociatedOrderItem) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetAsin

`func (o *AssociatedOrderItem) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *AssociatedOrderItem) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *AssociatedOrderItem) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *AssociatedOrderItem) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *AssociatedOrderItem) GetPurchasePrice() Money`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *AssociatedOrderItem) GetPurchasePriceOk() (*Money, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *AssociatedOrderItem) SetPurchasePrice(v Money)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *AssociatedOrderItem) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


