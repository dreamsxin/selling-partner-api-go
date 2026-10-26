# PropertyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display label of the property group. | [optional] 
**Description** | Pointer to **string** | The description of the property group. | [optional] 
**PropertyNames** | Pointer to **[]string** | The names of the schema properties for the property group. | [optional] 

## Methods

### NewPropertyGroup

`func NewPropertyGroup() *PropertyGroup`

NewPropertyGroup instantiates a new PropertyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupWithDefaults

`func NewPropertyGroupWithDefaults() *PropertyGroup`

NewPropertyGroupWithDefaults instantiates a new PropertyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PropertyGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PropertyGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PropertyGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PropertyGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPropertyNames

`func (o *PropertyGroup) GetPropertyNames() []string`

GetPropertyNames returns the PropertyNames field if non-nil, zero value otherwise.

### GetPropertyNamesOk

`func (o *PropertyGroup) GetPropertyNamesOk() (*[]string, bool)`

GetPropertyNamesOk returns a tuple with the PropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNames

`func (o *PropertyGroup) SetPropertyNames(v []string)`

SetPropertyNames sets PropertyNames field to given value.

### HasPropertyNames

`func (o *PropertyGroup) HasPropertyNames() bool`

HasPropertyNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


