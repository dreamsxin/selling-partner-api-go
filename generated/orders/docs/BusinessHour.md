# BusinessHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **string** | Specific day of the week for which operating hours are being defined. | [optional] 
**TimeWindows** | Pointer to [**[]TimeWindow**](TimeWindow.md) | Collection of time windows during which the location is available for deliveries on the specified day. | [optional] 

## Methods

### NewBusinessHour

`func NewBusinessHour() *BusinessHour`

NewBusinessHour instantiates a new BusinessHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessHourWithDefaults

`func NewBusinessHourWithDefaults() *BusinessHour`

NewBusinessHourWithDefaults instantiates a new BusinessHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *BusinessHour) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *BusinessHour) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *BusinessHour) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *BusinessHour) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetTimeWindows

`func (o *BusinessHour) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *BusinessHour) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *BusinessHour) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.

### HasTimeWindows

`func (o *BusinessHour) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


