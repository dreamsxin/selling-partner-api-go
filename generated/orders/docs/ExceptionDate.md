# ExceptionDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionDate** | Pointer to **string** | Specific calendar date when normal operating hours do not apply. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format at day granularity. | [optional] 
**ExceptionDateType** | Pointer to **string** | Operational status of the business on the specified exception date. | [optional] 
**TimeWindows** | Pointer to [**[]TimeWindow**](TimeWindow.md) | Alternative operating hours that apply specifically to this exception date. | [optional] 

## Methods

### NewExceptionDate

`func NewExceptionDate() *ExceptionDate`

NewExceptionDate instantiates a new ExceptionDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionDateWithDefaults

`func NewExceptionDateWithDefaults() *ExceptionDate`

NewExceptionDateWithDefaults instantiates a new ExceptionDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionDate

`func (o *ExceptionDate) GetExceptionDate() string`

GetExceptionDate returns the ExceptionDate field if non-nil, zero value otherwise.

### GetExceptionDateOk

`func (o *ExceptionDate) GetExceptionDateOk() (*string, bool)`

GetExceptionDateOk returns a tuple with the ExceptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionDate

`func (o *ExceptionDate) SetExceptionDate(v string)`

SetExceptionDate sets ExceptionDate field to given value.

### HasExceptionDate

`func (o *ExceptionDate) HasExceptionDate() bool`

HasExceptionDate returns a boolean if a field has been set.

### GetExceptionDateType

`func (o *ExceptionDate) GetExceptionDateType() string`

GetExceptionDateType returns the ExceptionDateType field if non-nil, zero value otherwise.

### GetExceptionDateTypeOk

`func (o *ExceptionDate) GetExceptionDateTypeOk() (*string, bool)`

GetExceptionDateTypeOk returns a tuple with the ExceptionDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionDateType

`func (o *ExceptionDate) SetExceptionDateType(v string)`

SetExceptionDateType sets ExceptionDateType field to given value.

### HasExceptionDateType

`func (o *ExceptionDate) HasExceptionDateType() bool`

HasExceptionDateType returns a boolean if a field has been set.

### GetTimeWindows

`func (o *ExceptionDate) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *ExceptionDate) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *ExceptionDate) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.

### HasTimeWindows

`func (o *ExceptionDate) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


