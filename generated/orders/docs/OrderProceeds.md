# OrderProceeds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrandTotal** | Pointer to [**Money**](Money.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]OrderProceedsBreakdown**](OrderProceedsBreakdown.md) | Categorized proceeds for the order. Proceed categories are either aggregated across all order items (such as &#x60;ITEM&#x60;, &#x60;SHIPPING&#x60;, and &#x60;TAX&#x60;) or applied at the order level (such as &#x60;DELIVERY_TIP&#x60;). | [optional] 

## Methods

### NewOrderProceeds

`func NewOrderProceeds() *OrderProceeds`

NewOrderProceeds instantiates a new OrderProceeds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderProceedsWithDefaults

`func NewOrderProceedsWithDefaults() *OrderProceeds`

NewOrderProceedsWithDefaults instantiates a new OrderProceeds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrandTotal

`func (o *OrderProceeds) GetGrandTotal() Money`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *OrderProceeds) GetGrandTotalOk() (*Money, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *OrderProceeds) SetGrandTotal(v Money)`

SetGrandTotal sets GrandTotal field to given value.

### HasGrandTotal

`func (o *OrderProceeds) HasGrandTotal() bool`

HasGrandTotal returns a boolean if a field has been set.

### GetBreakdowns

`func (o *OrderProceeds) GetBreakdowns() []OrderProceedsBreakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *OrderProceeds) GetBreakdownsOk() (*[]OrderProceedsBreakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *OrderProceeds) SetBreakdowns(v []OrderProceedsBreakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *OrderProceeds) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


