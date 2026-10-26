# HourMinute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** | The hour when the business opens or closes, in 24-hour format (0-23). | [optional] 
**Minute** | Pointer to **int32** | The minute when the business opens or closes. | [optional] 

## Methods

### NewHourMinute

`func NewHourMinute() *HourMinute`

NewHourMinute instantiates a new HourMinute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHourMinuteWithDefaults

`func NewHourMinuteWithDefaults() *HourMinute`

NewHourMinuteWithDefaults instantiates a new HourMinute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *HourMinute) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *HourMinute) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *HourMinute) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *HourMinute) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *HourMinute) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *HourMinute) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *HourMinute) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *HourMinute) HasMinute() bool`

HasMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


