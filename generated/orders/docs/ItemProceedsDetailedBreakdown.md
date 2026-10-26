# ItemProceedsDetailedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtype** | Pointer to **string** | Specific classification of the further granular breakdown.   **Possible values**: &#x60;ITEM&#x60;, &#x60;SHIPPING&#x60;, &#x60;GIFT_WRAP&#x60;, &#x60;COD_FEE&#x60;, &#x60;OTHER&#x60;, &#x60;DISCOUNT&#x60; | [optional] 
**Value** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewItemProceedsDetailedBreakdown

`func NewItemProceedsDetailedBreakdown() *ItemProceedsDetailedBreakdown`

NewItemProceedsDetailedBreakdown instantiates a new ItemProceedsDetailedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemProceedsDetailedBreakdownWithDefaults

`func NewItemProceedsDetailedBreakdownWithDefaults() *ItemProceedsDetailedBreakdown`

NewItemProceedsDetailedBreakdownWithDefaults instantiates a new ItemProceedsDetailedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtype

`func (o *ItemProceedsDetailedBreakdown) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ItemProceedsDetailedBreakdown) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ItemProceedsDetailedBreakdown) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ItemProceedsDetailedBreakdown) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetValue

`func (o *ItemProceedsDetailedBreakdown) GetValue() Money`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemProceedsDetailedBreakdown) GetValueOk() (*Money, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemProceedsDetailedBreakdown) SetValue(v Money)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemProceedsDetailedBreakdown) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


