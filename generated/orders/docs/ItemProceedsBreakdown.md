# ItemProceedsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Category classification of the proceeds breakdown.   **Possible values**: &#x60;ITEM&#x60;, &#x60;SHIPPING&#x60;, &#x60;GIFT_WRAP&#x60;, &#x60;COD_FEE&#x60;, &#x60;OTHER&#x60;, &#x60;TAX&#x60;, &#x60;DISCOUNT&#x60; | 
**Subtotal** | [**Money**](Money.md) |  | 
**DetailedBreakdowns** | Pointer to [**[]ItemProceedsDetailedBreakdown**](ItemProceedsDetailedBreakdown.md) | Further granular breakdown of the subtotal. | [optional] 

## Methods

### NewItemProceedsBreakdown

`func NewItemProceedsBreakdown(type_ string, subtotal Money, ) *ItemProceedsBreakdown`

NewItemProceedsBreakdown instantiates a new ItemProceedsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemProceedsBreakdownWithDefaults

`func NewItemProceedsBreakdownWithDefaults() *ItemProceedsBreakdown`

NewItemProceedsBreakdownWithDefaults instantiates a new ItemProceedsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ItemProceedsBreakdown) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemProceedsBreakdown) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemProceedsBreakdown) SetType(v string)`

SetType sets Type field to given value.


### GetSubtotal

`func (o *ItemProceedsBreakdown) GetSubtotal() Money`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *ItemProceedsBreakdown) GetSubtotalOk() (*Money, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *ItemProceedsBreakdown) SetSubtotal(v Money)`

SetSubtotal sets Subtotal field to given value.


### GetDetailedBreakdowns

`func (o *ItemProceedsBreakdown) GetDetailedBreakdowns() []ItemProceedsDetailedBreakdown`

GetDetailedBreakdowns returns the DetailedBreakdowns field if non-nil, zero value otherwise.

### GetDetailedBreakdownsOk

`func (o *ItemProceedsBreakdown) GetDetailedBreakdownsOk() (*[]ItemProceedsDetailedBreakdown, bool)`

GetDetailedBreakdownsOk returns a tuple with the DetailedBreakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedBreakdowns

`func (o *ItemProceedsBreakdown) SetDetailedBreakdowns(v []ItemProceedsDetailedBreakdown)`

SetDetailedBreakdowns sets DetailedBreakdowns field to given value.

### HasDetailedBreakdowns

`func (o *ItemProceedsBreakdown) HasDetailedBreakdowns() bool`

HasDetailedBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


