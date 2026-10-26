# OrderProceedsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The proceeds category.   **Possible values**: &#x60;ITEM&#x60;, &#x60;SHIPPING&#x60;, &#x60;GIFT_WRAP&#x60;, &#x60;COD_FEE&#x60;, &#x60;TAX&#x60;, &#x60;DISCOUNT&#x60;, &#x60;DELIVERY_TIP&#x60;, &#x60;OTHER&#x60;. **Note:** &#x60;DELIVERY_TIP&#x60; is charged separately and not attributed to a specific item. The remaining categories are aggregated across all order items. | 
**Status** | Pointer to **string** | The processing status of the charge. Only present for categories processed separately after checkout, such as &#x60;DELIVERY_TIP&#x60;.  **Possible values**: &#x60;PENDING&#x60;, &#x60;FINALIZED&#x60;. | [optional] 
**Subtotal** | [**Money**](Money.md) |  | 

## Methods

### NewOrderProceedsBreakdown

`func NewOrderProceedsBreakdown(type_ string, subtotal Money, ) *OrderProceedsBreakdown`

NewOrderProceedsBreakdown instantiates a new OrderProceedsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderProceedsBreakdownWithDefaults

`func NewOrderProceedsBreakdownWithDefaults() *OrderProceedsBreakdown`

NewOrderProceedsBreakdownWithDefaults instantiates a new OrderProceedsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderProceedsBreakdown) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderProceedsBreakdown) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderProceedsBreakdown) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *OrderProceedsBreakdown) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderProceedsBreakdown) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderProceedsBreakdown) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderProceedsBreakdown) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtotal

`func (o *OrderProceedsBreakdown) GetSubtotal() Money`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *OrderProceedsBreakdown) GetSubtotalOk() (*Money, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *OrderProceedsBreakdown) SetSubtotal(v Money)`

SetSubtotal sets Subtotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


