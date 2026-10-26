# ItemTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxCalculationBreakdowns** | Pointer to [**[]ItemTaxCalculationBreakdown**](ItemTaxCalculationBreakdown.md) | A list of tax calculation breakdowns for the order item. | [optional] 
**TaxCollections** | Pointer to [**[]ItemTaxCollection**](ItemTaxCollection.md) | A list of tax collections for the order item. | [optional] 

## Methods

### NewItemTax

`func NewItemTax() *ItemTax`

NewItemTax instantiates a new ItemTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTaxWithDefaults

`func NewItemTaxWithDefaults() *ItemTax`

NewItemTaxWithDefaults instantiates a new ItemTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxCalculationBreakdowns

`func (o *ItemTax) GetTaxCalculationBreakdowns() []ItemTaxCalculationBreakdown`

GetTaxCalculationBreakdowns returns the TaxCalculationBreakdowns field if non-nil, zero value otherwise.

### GetTaxCalculationBreakdownsOk

`func (o *ItemTax) GetTaxCalculationBreakdownsOk() (*[]ItemTaxCalculationBreakdown, bool)`

GetTaxCalculationBreakdownsOk returns a tuple with the TaxCalculationBreakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationBreakdowns

`func (o *ItemTax) SetTaxCalculationBreakdowns(v []ItemTaxCalculationBreakdown)`

SetTaxCalculationBreakdowns sets TaxCalculationBreakdowns field to given value.

### HasTaxCalculationBreakdowns

`func (o *ItemTax) HasTaxCalculationBreakdowns() bool`

HasTaxCalculationBreakdowns returns a boolean if a field has been set.

### GetTaxCollections

`func (o *ItemTax) GetTaxCollections() []ItemTaxCollection`

GetTaxCollections returns the TaxCollections field if non-nil, zero value otherwise.

### GetTaxCollectionsOk

`func (o *ItemTax) GetTaxCollectionsOk() (*[]ItemTaxCollection, bool)`

GetTaxCollectionsOk returns a tuple with the TaxCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCollections

`func (o *ItemTax) SetTaxCollections(v []ItemTaxCollection)`

SetTaxCollections sets TaxCollections field to given value.

### HasTaxCollections

`func (o *ItemTax) HasTaxCollections() bool`

HasTaxCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


