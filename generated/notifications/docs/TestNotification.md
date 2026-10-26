# TestNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadVersion** | **string** | The version of the payload object to be used in the notification. | 
**TestScenario** | Pointer to **string** | The scenario of the specified notification to be used in the notification payload. If testScenario is empty, a 400 response will be returned back to the developer. The scenarios supported for each notification type can be found in the Selling Partner API Developer Guide. | [optional] 

## Methods

### NewTestNotification

`func NewTestNotification(payloadVersion string, ) *TestNotification`

NewTestNotification instantiates a new TestNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestNotificationWithDefaults

`func NewTestNotificationWithDefaults() *TestNotification`

NewTestNotificationWithDefaults instantiates a new TestNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadVersion

`func (o *TestNotification) GetPayloadVersion() string`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *TestNotification) GetPayloadVersionOk() (*string, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *TestNotification) SetPayloadVersion(v string)`

SetPayloadVersion sets PayloadVersion field to given value.


### GetTestScenario

`func (o *TestNotification) GetTestScenario() string`

GetTestScenario returns the TestScenario field if non-nil, zero value otherwise.

### GetTestScenarioOk

`func (o *TestNotification) GetTestScenarioOk() (*string, bool)`

GetTestScenarioOk returns a tuple with the TestScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestScenario

`func (o *TestNotification) SetTestScenario(v string)`

SetTestScenario sets TestScenario field to given value.

### HasTestScenario

`func (o *TestNotification) HasTestScenario() bool`

HasTestScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


