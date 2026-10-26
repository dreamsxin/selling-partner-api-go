# IssueEnforcementAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The enforcement action name.   Possible values:   * &#x60;LISTING_SUPPRESSED&#x60; - This enforcement takes down the current listing item&#39;s buyability.   * &#x60;ATTRIBUTE_SUPPRESSED&#x60; - An attribute&#39;s value on the listing item is invalid, which causes it to be rejected by Amazon.   * &#x60;CATALOG_ITEM_REMOVED&#x60; - This catalog item is inactive on Amazon, and all offers against it in the applicable Amazon store are non-buyable.   * &#x60;SEARCH_SUPPRESSED&#x60; - This value indicates that the catalog item is hidden from search results. | 

## Methods

### NewIssueEnforcementAction

`func NewIssueEnforcementAction(action string, ) *IssueEnforcementAction`

NewIssueEnforcementAction instantiates a new IssueEnforcementAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEnforcementActionWithDefaults

`func NewIssueEnforcementActionWithDefaults() *IssueEnforcementAction`

NewIssueEnforcementActionWithDefaults instantiates a new IssueEnforcementAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IssueEnforcementAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IssueEnforcementAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IssueEnforcementAction) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


