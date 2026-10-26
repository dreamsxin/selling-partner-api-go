# ItemProceeds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProceedsTotal** | Pointer to [**Money**](Money.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]ItemProceedsBreakdown**](ItemProceedsBreakdown.md) | The breakdown of proceeds. | [optional] 

## Methods

### NewItemProceeds

`func NewItemProceeds() *ItemProceeds`

NewItemProceeds instantiates a new ItemProceeds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemProceedsWithDefaults

`func NewItemProceedsWithDefaults() *ItemProceeds`

NewItemProceedsWithDefaults instantiates a new ItemProceeds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProceedsTotal

`func (o *ItemProceeds) GetProceedsTotal() Money`

GetProceedsTotal returns the ProceedsTotal field if non-nil, zero value otherwise.

### GetProceedsTotalOk

`func (o *ItemProceeds) GetProceedsTotalOk() (*Money, bool)`

GetProceedsTotalOk returns a tuple with the ProceedsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceedsTotal

`func (o *ItemProceeds) SetProceedsTotal(v Money)`

SetProceedsTotal sets ProceedsTotal field to given value.

### HasProceedsTotal

`func (o *ItemProceeds) HasProceedsTotal() bool`

HasProceedsTotal returns a boolean if a field has been set.

### GetBreakdowns

`func (o *ItemProceeds) GetBreakdowns() []ItemProceedsBreakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *ItemProceeds) GetBreakdownsOk() (*[]ItemProceedsBreakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *ItemProceeds) SetBreakdowns(v []ItemProceedsBreakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *ItemProceeds) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


