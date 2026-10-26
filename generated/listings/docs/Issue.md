# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | An issue code that identifies the type of issue. | 
**Message** | **string** | A message that describes the issue. | 
**Severity** | **string** | The severity of the issue. | 
**AttributeNames** | Pointer to **[]string** | Names of the attributes that are associated with the issue, if applicable. | [optional] 
**Categories** | **[]string** | List of issue categories. | 
**Enforcements** | Pointer to [**IssueEnforcements**](IssueEnforcements.md) |  | [optional] 
**MarketplaceIds** | Pointer to **[]string** | List of Amazon store identifiers. | [optional] 

## Methods

### NewIssue

`func NewIssue(code string, message string, severity string, categories []string, ) *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Issue) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Issue) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Issue) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *Issue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Issue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Issue) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeverity

`func (o *Issue) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Issue) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Issue) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetAttributeNames

`func (o *Issue) GetAttributeNames() []string`

GetAttributeNames returns the AttributeNames field if non-nil, zero value otherwise.

### GetAttributeNamesOk

`func (o *Issue) GetAttributeNamesOk() (*[]string, bool)`

GetAttributeNamesOk returns a tuple with the AttributeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeNames

`func (o *Issue) SetAttributeNames(v []string)`

SetAttributeNames sets AttributeNames field to given value.

### HasAttributeNames

`func (o *Issue) HasAttributeNames() bool`

HasAttributeNames returns a boolean if a field has been set.

### GetCategories

`func (o *Issue) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Issue) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Issue) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetEnforcements

`func (o *Issue) GetEnforcements() IssueEnforcements`

GetEnforcements returns the Enforcements field if non-nil, zero value otherwise.

### GetEnforcementsOk

`func (o *Issue) GetEnforcementsOk() (*IssueEnforcements, bool)`

GetEnforcementsOk returns a tuple with the Enforcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcements

`func (o *Issue) SetEnforcements(v IssueEnforcements)`

SetEnforcements sets Enforcements field to given value.

### HasEnforcements

`func (o *Issue) HasEnforcements() bool`

HasEnforcements returns a boolean if a field has been set.

### GetMarketplaceIds

`func (o *Issue) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *Issue) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *Issue) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.

### HasMarketplaceIds

`func (o *Issue) HasMarketplaceIds() bool`

HasMarketplaceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


