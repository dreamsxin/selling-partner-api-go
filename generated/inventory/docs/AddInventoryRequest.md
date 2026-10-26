# AddInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryItems** | Pointer to [**[]InventoryItem**](InventoryItem.md) | List of Inventory to be added | [optional] 

## Methods

### NewAddInventoryRequest

`func NewAddInventoryRequest() *AddInventoryRequest`

NewAddInventoryRequest instantiates a new AddInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInventoryRequestWithDefaults

`func NewAddInventoryRequestWithDefaults() *AddInventoryRequest`

NewAddInventoryRequestWithDefaults instantiates a new AddInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryItems

`func (o *AddInventoryRequest) GetInventoryItems() []InventoryItem`

GetInventoryItems returns the InventoryItems field if non-nil, zero value otherwise.

### GetInventoryItemsOk

`func (o *AddInventoryRequest) GetInventoryItemsOk() (*[]InventoryItem, bool)`

GetInventoryItemsOk returns a tuple with the InventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItems

`func (o *AddInventoryRequest) SetInventoryItems(v []InventoryItem)`

SetInventoryItems sets InventoryItems field to given value.

### HasInventoryItems

`func (o *AddInventoryRequest) HasInventoryItems() bool`

HasInventoryItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


