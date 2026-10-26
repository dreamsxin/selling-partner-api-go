# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** | A unique identifier for this specific item within the order. | 
**QuantityOrdered** | **int32** | The number of units of this item that the customer ordered. | 
**Measurement** | Pointer to [**Measurement**](Measurement.md) |  | [optional] 
**AssociatedOrderItems** | Pointer to [**[]AssociatedOrderItem**](AssociatedOrderItem.md) | A list of order items associated with this item. For example, a value-add service purchased with the product, or a product covered by a warranty protection plan. | [optional] 
**Programs** | Pointer to **[]string** | Special programs that apply specifically to this item within the order.  **Possible values**: &#x60;TRANSPARENCY&#x60;, &#x60;SUBSCRIBE_AND_SAVE&#x60; | [optional] 
**Product** | [**ItemProduct**](ItemProduct.md) |  | 
**Proceeds** | Pointer to [**ItemProceeds**](ItemProceeds.md) |  | [optional] 
**Expense** | Pointer to [**ItemExpense**](ItemExpense.md) |  | [optional] 
**Promotion** | Pointer to [**ItemPromotion**](ItemPromotion.md) |  | [optional] 
**Cancellation** | Pointer to [**ItemCancellation**](ItemCancellation.md) |  | [optional] 
**Fulfillment** | Pointer to [**ItemFulfillment**](ItemFulfillment.md) |  | [optional] 
**Tax** | Pointer to [**ItemTax**](ItemTax.md) |  | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem(orderItemId string, quantityOrdered int32, product ItemProduct, ) *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *OrderItem) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *OrderItem) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *OrderItem) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetQuantityOrdered

`func (o *OrderItem) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderItem) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderItem) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.


### GetMeasurement

`func (o *OrderItem) GetMeasurement() Measurement`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *OrderItem) GetMeasurementOk() (*Measurement, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *OrderItem) SetMeasurement(v Measurement)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *OrderItem) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.

### GetAssociatedOrderItems

`func (o *OrderItem) GetAssociatedOrderItems() []AssociatedOrderItem`

GetAssociatedOrderItems returns the AssociatedOrderItems field if non-nil, zero value otherwise.

### GetAssociatedOrderItemsOk

`func (o *OrderItem) GetAssociatedOrderItemsOk() (*[]AssociatedOrderItem, bool)`

GetAssociatedOrderItemsOk returns a tuple with the AssociatedOrderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedOrderItems

`func (o *OrderItem) SetAssociatedOrderItems(v []AssociatedOrderItem)`

SetAssociatedOrderItems sets AssociatedOrderItems field to given value.

### HasAssociatedOrderItems

`func (o *OrderItem) HasAssociatedOrderItems() bool`

HasAssociatedOrderItems returns a boolean if a field has been set.

### GetPrograms

`func (o *OrderItem) GetPrograms() []string`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *OrderItem) GetProgramsOk() (*[]string, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *OrderItem) SetPrograms(v []string)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *OrderItem) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetProduct

`func (o *OrderItem) GetProduct() ItemProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *OrderItem) GetProductOk() (*ItemProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *OrderItem) SetProduct(v ItemProduct)`

SetProduct sets Product field to given value.


### GetProceeds

`func (o *OrderItem) GetProceeds() ItemProceeds`

GetProceeds returns the Proceeds field if non-nil, zero value otherwise.

### GetProceedsOk

`func (o *OrderItem) GetProceedsOk() (*ItemProceeds, bool)`

GetProceedsOk returns a tuple with the Proceeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceeds

`func (o *OrderItem) SetProceeds(v ItemProceeds)`

SetProceeds sets Proceeds field to given value.

### HasProceeds

`func (o *OrderItem) HasProceeds() bool`

HasProceeds returns a boolean if a field has been set.

### GetExpense

`func (o *OrderItem) GetExpense() ItemExpense`

GetExpense returns the Expense field if non-nil, zero value otherwise.

### GetExpenseOk

`func (o *OrderItem) GetExpenseOk() (*ItemExpense, bool)`

GetExpenseOk returns a tuple with the Expense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpense

`func (o *OrderItem) SetExpense(v ItemExpense)`

SetExpense sets Expense field to given value.

### HasExpense

`func (o *OrderItem) HasExpense() bool`

HasExpense returns a boolean if a field has been set.

### GetPromotion

`func (o *OrderItem) GetPromotion() ItemPromotion`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *OrderItem) GetPromotionOk() (*ItemPromotion, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *OrderItem) SetPromotion(v ItemPromotion)`

SetPromotion sets Promotion field to given value.

### HasPromotion

`func (o *OrderItem) HasPromotion() bool`

HasPromotion returns a boolean if a field has been set.

### GetCancellation

`func (o *OrderItem) GetCancellation() ItemCancellation`

GetCancellation returns the Cancellation field if non-nil, zero value otherwise.

### GetCancellationOk

`func (o *OrderItem) GetCancellationOk() (*ItemCancellation, bool)`

GetCancellationOk returns a tuple with the Cancellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellation

`func (o *OrderItem) SetCancellation(v ItemCancellation)`

SetCancellation sets Cancellation field to given value.

### HasCancellation

`func (o *OrderItem) HasCancellation() bool`

HasCancellation returns a boolean if a field has been set.

### GetFulfillment

`func (o *OrderItem) GetFulfillment() ItemFulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *OrderItem) GetFulfillmentOk() (*ItemFulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *OrderItem) SetFulfillment(v ItemFulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *OrderItem) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetTax

`func (o *OrderItem) GetTax() ItemTax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderItem) GetTaxOk() (*ItemTax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderItem) SetTax(v ItemTax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderItem) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


