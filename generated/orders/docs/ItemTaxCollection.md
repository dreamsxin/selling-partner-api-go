# ItemTaxCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | The tax collection model applied to the item.  **Possible values**: - &#x60;MARKETPLACE_FACILITATOR&#x60; (Tax is withheld and remitted to the taxing authority by Amazon on behalf of the seller) | [optional] 
**ResponsibleParty** | Pointer to **string** | The party responsible for withholding the taxes and remitting them to the taxing authority. | [optional] 

## Methods

### NewItemTaxCollection

`func NewItemTaxCollection() *ItemTaxCollection`

NewItemTaxCollection instantiates a new ItemTaxCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTaxCollectionWithDefaults

`func NewItemTaxCollectionWithDefaults() *ItemTaxCollection`

NewItemTaxCollectionWithDefaults instantiates a new ItemTaxCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ItemTaxCollection) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ItemTaxCollection) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ItemTaxCollection) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ItemTaxCollection) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResponsibleParty

`func (o *ItemTaxCollection) GetResponsibleParty() string`

GetResponsibleParty returns the ResponsibleParty field if non-nil, zero value otherwise.

### GetResponsiblePartyOk

`func (o *ItemTaxCollection) GetResponsiblePartyOk() (*string, bool)`

GetResponsiblePartyOk returns a tuple with the ResponsibleParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleParty

`func (o *ItemTaxCollection) SetResponsibleParty(v string)`

SetResponsibleParty sets ResponsibleParty field to given value.

### HasResponsibleParty

`func (o *ItemTaxCollection) HasResponsibleParty() bool`

HasResponsibleParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


