# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the items in a transaction. | [optional] 
**RelatedIdentifiers** | Pointer to [**[]ItemRelatedIdentifier**](ItemRelatedIdentifier.md) | Related business identifiers of the item in the transaction. | [optional] 
**TotalAmount** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]Breakdown**](Breakdown.md) | A list of breakdowns that detail how the total amount is calculated for the transaction. | [optional] 
**Contexts** | Pointer to [**[]Context**](Context.md) | A list of additional information about the item. | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Item) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRelatedIdentifiers

`func (o *Item) GetRelatedIdentifiers() []ItemRelatedIdentifier`

GetRelatedIdentifiers returns the RelatedIdentifiers field if non-nil, zero value otherwise.

### GetRelatedIdentifiersOk

`func (o *Item) GetRelatedIdentifiersOk() (*[]ItemRelatedIdentifier, bool)`

GetRelatedIdentifiersOk returns a tuple with the RelatedIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIdentifiers

`func (o *Item) SetRelatedIdentifiers(v []ItemRelatedIdentifier)`

SetRelatedIdentifiers sets RelatedIdentifiers field to given value.

### HasRelatedIdentifiers

`func (o *Item) HasRelatedIdentifiers() bool`

HasRelatedIdentifiers returns a boolean if a field has been set.

### GetTotalAmount

`func (o *Item) GetTotalAmount() Currency`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Item) GetTotalAmountOk() (*Currency, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Item) SetTotalAmount(v Currency)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Item) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetBreakdowns

`func (o *Item) GetBreakdowns() []Breakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *Item) GetBreakdownsOk() (*[]Breakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *Item) SetBreakdowns(v []Breakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *Item) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### GetContexts

`func (o *Item) GetContexts() []Context`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *Item) GetContextsOk() (*[]Context, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *Item) SetContexts(v []Context)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *Item) HasContexts() bool`

HasContexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


