# ItemContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**ItemContributorRole**](ItemContributorRole.md) |  | 
**Value** | **string** | Name of the contributor, such as &#x60;Jane Austen&#x60;. | 

## Methods

### NewItemContributor

`func NewItemContributor(role ItemContributorRole, value string, ) *ItemContributor`

NewItemContributor instantiates a new ItemContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemContributorWithDefaults

`func NewItemContributorWithDefaults() *ItemContributor`

NewItemContributorWithDefaults instantiates a new ItemContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ItemContributor) GetRole() ItemContributorRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ItemContributor) GetRoleOk() (*ItemContributorRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ItemContributor) SetRole(v ItemContributorRole)`

SetRole sets Role field to given value.


### GetValue

`func (o *ItemContributor) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemContributor) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemContributor) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


