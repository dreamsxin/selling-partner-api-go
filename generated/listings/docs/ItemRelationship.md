# ItemRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildSkus** | Pointer to **[]string** | Identifiers (SKUs) of the related items that are children of this listing item. | [optional] 
**ParentSkus** | Pointer to **[]string** | Identifiers (SKUs) of the related items that are parents of this listing item. | [optional] 
**VariationTheme** | Pointer to [**ItemVariationTheme**](ItemVariationTheme.md) |  | [optional] 
**Type** | **string** | The type of relationship. | 

## Methods

### NewItemRelationship

`func NewItemRelationship(type_ string, ) *ItemRelationship`

NewItemRelationship instantiates a new ItemRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRelationshipWithDefaults

`func NewItemRelationshipWithDefaults() *ItemRelationship`

NewItemRelationshipWithDefaults instantiates a new ItemRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildSkus

`func (o *ItemRelationship) GetChildSkus() []string`

GetChildSkus returns the ChildSkus field if non-nil, zero value otherwise.

### GetChildSkusOk

`func (o *ItemRelationship) GetChildSkusOk() (*[]string, bool)`

GetChildSkusOk returns a tuple with the ChildSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSkus

`func (o *ItemRelationship) SetChildSkus(v []string)`

SetChildSkus sets ChildSkus field to given value.

### HasChildSkus

`func (o *ItemRelationship) HasChildSkus() bool`

HasChildSkus returns a boolean if a field has been set.

### GetParentSkus

`func (o *ItemRelationship) GetParentSkus() []string`

GetParentSkus returns the ParentSkus field if non-nil, zero value otherwise.

### GetParentSkusOk

`func (o *ItemRelationship) GetParentSkusOk() (*[]string, bool)`

GetParentSkusOk returns a tuple with the ParentSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSkus

`func (o *ItemRelationship) SetParentSkus(v []string)`

SetParentSkus sets ParentSkus field to given value.

### HasParentSkus

`func (o *ItemRelationship) HasParentSkus() bool`

HasParentSkus returns a boolean if a field has been set.

### GetVariationTheme

`func (o *ItemRelationship) GetVariationTheme() ItemVariationTheme`

GetVariationTheme returns the VariationTheme field if non-nil, zero value otherwise.

### GetVariationThemeOk

`func (o *ItemRelationship) GetVariationThemeOk() (*ItemVariationTheme, bool)`

GetVariationThemeOk returns a tuple with the VariationTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationTheme

`func (o *ItemRelationship) SetVariationTheme(v ItemVariationTheme)`

SetVariationTheme sets VariationTheme field to given value.

### HasVariationTheme

`func (o *ItemRelationship) HasVariationTheme() bool`

HasVariationTheme returns a boolean if a field has been set.

### GetType

`func (o *ItemRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemRelationship) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


