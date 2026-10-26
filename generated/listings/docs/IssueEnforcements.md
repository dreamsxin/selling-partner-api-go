# IssueEnforcements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]IssueEnforcementAction**](IssueEnforcementAction.md) | List of enforcement actions taken by Amazon that affect the publishing or status of a listing. | 
**Exemption** | [**IssueExemption**](IssueExemption.md) |  | 

## Methods

### NewIssueEnforcements

`func NewIssueEnforcements(actions []IssueEnforcementAction, exemption IssueExemption, ) *IssueEnforcements`

NewIssueEnforcements instantiates a new IssueEnforcements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEnforcementsWithDefaults

`func NewIssueEnforcementsWithDefaults() *IssueEnforcements`

NewIssueEnforcementsWithDefaults instantiates a new IssueEnforcements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *IssueEnforcements) GetActions() []IssueEnforcementAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IssueEnforcements) GetActionsOk() (*[]IssueEnforcementAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IssueEnforcements) SetActions(v []IssueEnforcementAction)`

SetActions sets Actions field to given value.


### GetExemption

`func (o *IssueEnforcements) GetExemption() IssueExemption`

GetExemption returns the Exemption field if non-nil, zero value otherwise.

### GetExemptionOk

`func (o *IssueEnforcements) GetExemptionOk() (*IssueExemption, bool)`

GetExemptionOk returns a tuple with the Exemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemption

`func (o *IssueEnforcements) SetExemption(v IssueExemption)`

SetExemption sets Exemption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


