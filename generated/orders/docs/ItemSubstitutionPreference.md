# ItemSubstitutionPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubstitutionType** | **string** | Source and nature of the substitution preferences for this item. | 
**SubstitutionOptions** | Pointer to [**[]ItemSubstitutionOption**](ItemSubstitutionOption.md) | List of alternative products that can be substituted for the original item if it becomes unavailable. | [optional] 

## Methods

### NewItemSubstitutionPreference

`func NewItemSubstitutionPreference(substitutionType string, ) *ItemSubstitutionPreference`

NewItemSubstitutionPreference instantiates a new ItemSubstitutionPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSubstitutionPreferenceWithDefaults

`func NewItemSubstitutionPreferenceWithDefaults() *ItemSubstitutionPreference`

NewItemSubstitutionPreferenceWithDefaults instantiates a new ItemSubstitutionPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubstitutionType

`func (o *ItemSubstitutionPreference) GetSubstitutionType() string`

GetSubstitutionType returns the SubstitutionType field if non-nil, zero value otherwise.

### GetSubstitutionTypeOk

`func (o *ItemSubstitutionPreference) GetSubstitutionTypeOk() (*string, bool)`

GetSubstitutionTypeOk returns a tuple with the SubstitutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionType

`func (o *ItemSubstitutionPreference) SetSubstitutionType(v string)`

SetSubstitutionType sets SubstitutionType field to given value.


### GetSubstitutionOptions

`func (o *ItemSubstitutionPreference) GetSubstitutionOptions() []ItemSubstitutionOption`

GetSubstitutionOptions returns the SubstitutionOptions field if non-nil, zero value otherwise.

### GetSubstitutionOptionsOk

`func (o *ItemSubstitutionPreference) GetSubstitutionOptionsOk() (*[]ItemSubstitutionOption, bool)`

GetSubstitutionOptionsOk returns a tuple with the SubstitutionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionOptions

`func (o *ItemSubstitutionPreference) SetSubstitutionOptions(v []ItemSubstitutionOption)`

SetSubstitutionOptions sets SubstitutionOptions field to given value.

### HasSubstitutionOptions

`func (o *ItemSubstitutionPreference) HasSubstitutionOptions() bool`

HasSubstitutionOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


