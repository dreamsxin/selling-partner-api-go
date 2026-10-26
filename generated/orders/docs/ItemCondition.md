# ItemCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | Pointer to **string** | The primary condition category that broadly describes the item&#39;s state.  **Possible values**: &#x60;NEW&#x60;, &#x60;USED&#x60;, &#x60;COLLECTIBLE&#x60;, &#x60;REFURBISHED&#x60;, &#x60;PREORDER&#x60;, &#x60;CLUB&#x60;. | [optional] 
**ConditionSubtype** | Pointer to **string** | A more specific condition classification that provides additional detail about the item&#39;s quality within the main condition type.  **Possible values**: &#x60;NEW&#x60;, &#x60;MINT&#x60;, &#x60;VERY_GOOD&#x60;, &#x60;GOOD&#x60;, &#x60;ACCEPTABLE&#x60;, &#x60;POOR&#x60;, &#x60;CLUB&#x60;, &#x60;OEM&#x60;, &#x60;WARRANTY&#x60;, &#x60;REFURBISHED_WARRANTY&#x60;, &#x60;REFURBISHED&#x60;, &#x60;OPEN_BOX&#x60;, &#x60;ANY&#x60;, &#x60;OTHER&#x60;. | [optional] 
**ConditionNote** | Pointer to **string** | Additional details provided by the seller to describe the specific condition of this particular item. | [optional] 

## Methods

### NewItemCondition

`func NewItemCondition() *ItemCondition`

NewItemCondition instantiates a new ItemCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemConditionWithDefaults

`func NewItemConditionWithDefaults() *ItemCondition`

NewItemConditionWithDefaults instantiates a new ItemCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *ItemCondition) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ItemCondition) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ItemCondition) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ItemCondition) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetConditionSubtype

`func (o *ItemCondition) GetConditionSubtype() string`

GetConditionSubtype returns the ConditionSubtype field if non-nil, zero value otherwise.

### GetConditionSubtypeOk

`func (o *ItemCondition) GetConditionSubtypeOk() (*string, bool)`

GetConditionSubtypeOk returns a tuple with the ConditionSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSubtype

`func (o *ItemCondition) SetConditionSubtype(v string)`

SetConditionSubtype sets ConditionSubtype field to given value.

### HasConditionSubtype

`func (o *ItemCondition) HasConditionSubtype() bool`

HasConditionSubtype returns a boolean if a field has been set.

### GetConditionNote

`func (o *ItemCondition) GetConditionNote() string`

GetConditionNote returns the ConditionNote field if non-nil, zero value otherwise.

### GetConditionNoteOk

`func (o *ItemCondition) GetConditionNoteOk() (*string, bool)`

GetConditionNoteOk returns a tuple with the ConditionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionNote

`func (o *ItemCondition) SetConditionNote(v string)`

SetConditionNote sets ConditionNote field to given value.

### HasConditionNote

`func (o *ItemCondition) HasConditionNote() bool`

HasConditionNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


