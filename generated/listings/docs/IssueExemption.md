# IssueExemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Indicates the current exemption status for the listed enforcement actions. It can take values such as &#x60;EXEMPT&#x60;, signifying permanent exemption, &#x60;EXEMPT_UNTIL_EXPIRY_DATE&#x60; indicating temporary exemption until a specified date, or &#x60;NOT_EXEMPT&#x60; signifying no exemptions, and enforcement actions were already applied. | 
**ExpiryDate** | Pointer to **time.Time** | Represents the timestamp, in ISO 8601 format, that specifies the date when the temporary exemptions expires, and Amazon begins enforcing the listed actions. | [optional] 

## Methods

### NewIssueExemption

`func NewIssueExemption(status string, ) *IssueExemption`

NewIssueExemption instantiates a new IssueExemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueExemptionWithDefaults

`func NewIssueExemptionWithDefaults() *IssueExemption`

NewIssueExemptionWithDefaults instantiates a new IssueExemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IssueExemption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssueExemption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssueExemption) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExpiryDate

`func (o *IssueExemption) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IssueExemption) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IssueExemption) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IssueExemption) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


