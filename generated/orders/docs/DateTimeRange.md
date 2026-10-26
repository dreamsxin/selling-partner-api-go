# DateTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EarliestDateTime** | Pointer to **time.Time** | The beginning of the time period, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 
**LatestDateTime** | Pointer to **time.Time** | The end of the time period, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | [optional] 

## Methods

### NewDateTimeRange

`func NewDateTimeRange() *DateTimeRange`

NewDateTimeRange instantiates a new DateTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeRangeWithDefaults

`func NewDateTimeRangeWithDefaults() *DateTimeRange`

NewDateTimeRangeWithDefaults instantiates a new DateTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliestDateTime

`func (o *DateTimeRange) GetEarliestDateTime() time.Time`

GetEarliestDateTime returns the EarliestDateTime field if non-nil, zero value otherwise.

### GetEarliestDateTimeOk

`func (o *DateTimeRange) GetEarliestDateTimeOk() (*time.Time, bool)`

GetEarliestDateTimeOk returns a tuple with the EarliestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTime

`func (o *DateTimeRange) SetEarliestDateTime(v time.Time)`

SetEarliestDateTime sets EarliestDateTime field to given value.

### HasEarliestDateTime

`func (o *DateTimeRange) HasEarliestDateTime() bool`

HasEarliestDateTime returns a boolean if a field has been set.

### GetLatestDateTime

`func (o *DateTimeRange) GetLatestDateTime() time.Time`

GetLatestDateTime returns the LatestDateTime field if non-nil, zero value otherwise.

### GetLatestDateTimeOk

`func (o *DateTimeRange) GetLatestDateTimeOk() (*time.Time, bool)`

GetLatestDateTimeOk returns a tuple with the LatestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTime

`func (o *DateTimeRange) SetLatestDateTime(v time.Time)`

SetLatestDateTime sets LatestDateTime field to given value.

### HasLatestDateTime

`func (o *DateTimeRange) HasLatestDateTime() bool`

HasLatestDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


