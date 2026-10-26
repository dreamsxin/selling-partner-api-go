# ItemRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildAsins** | Pointer to **[]string** | ASINs of the related items that are children of this item. | [optional] 
**ParentAsins** | Pointer to **[]string** | ASINs of the related items that are parents of this item. | [optional] 
**VariationTheme** | Pointer to [**ItemVariationTheme**](ItemVariationTheme.md) |  | [optional] 
**Type** | **string** | Type of relationship. | 

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

### GetChildAsins

`func (o *ItemRelationship) GetChildAsins() []string`

GetChildAsins returns the ChildAsins field if non-nil, zero value otherwise.

### GetChildAsinsOk

`func (o *ItemRelationship) GetChildAsinsOk() (*[]string, bool)`

GetChildAsinsOk returns a tuple with the ChildAsins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAsins

`func (o *ItemRelationship) SetChildAsins(v []string)`

SetChildAsins sets ChildAsins field to given value.

### HasChildAsins

`func (o *ItemRelationship) HasChildAsins() bool`

HasChildAsins returns a boolean if a field has been set.

### GetParentAsins

`func (o *ItemRelationship) GetParentAsins() []string`

GetParentAsins returns the ParentAsins field if non-nil, zero value otherwise.

### GetParentAsinsOk

`func (o *ItemRelationship) GetParentAsinsOk() (*[]string, bool)`

GetParentAsinsOk returns a tuple with the ParentAsins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAsins

`func (o *ItemRelationship) SetParentAsins(v []string)`

SetParentAsins sets ParentAsins field to given value.

### HasParentAsins

`func (o *ItemRelationship) HasParentAsins() bool`

HasParentAsins returns a boolean if a field has been set.

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


