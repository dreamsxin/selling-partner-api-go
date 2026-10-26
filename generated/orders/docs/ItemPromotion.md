# ItemPromotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdowns** | Pointer to [**[]ItemPromotionBreakdown**](ItemPromotionBreakdown.md) | A list of promotions applied to the order item. | [optional] 

## Methods

### NewItemPromotion

`func NewItemPromotion() *ItemPromotion`

NewItemPromotion instantiates a new ItemPromotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPromotionWithDefaults

`func NewItemPromotionWithDefaults() *ItemPromotion`

NewItemPromotionWithDefaults instantiates a new ItemPromotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdowns

`func (o *ItemPromotion) GetBreakdowns() []ItemPromotionBreakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *ItemPromotion) GetBreakdownsOk() (*[]ItemPromotionBreakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *ItemPromotion) SetBreakdowns(v []ItemPromotionBreakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *ItemPromotion) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


