# ItemContributorRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the role in the requested locale, such as &#x60;Author&#x60; or &#x60;Actor&#x60;. | [optional] 
**Value** | **string** | Role value for the Amazon catalog item, such as &#x60;author&#x60; or &#x60;actor&#x60;. | 

## Methods

### NewItemContributorRole

`func NewItemContributorRole(value string, ) *ItemContributorRole`

NewItemContributorRole instantiates a new ItemContributorRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemContributorRoleWithDefaults

`func NewItemContributorRoleWithDefaults() *ItemContributorRole`

NewItemContributorRoleWithDefaults instantiates a new ItemContributorRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ItemContributorRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ItemContributorRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ItemContributorRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ItemContributorRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *ItemContributorRole) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemContributorRole) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemContributorRole) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


