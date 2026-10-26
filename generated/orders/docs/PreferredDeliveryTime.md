# PreferredDeliveryTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessHours** | Pointer to [**[]BusinessHour**](BusinessHour.md) | Business hours when the business is open for deliveries. | [optional] 
**ExceptionDates** | Pointer to [**[]ExceptionDate**](ExceptionDate.md) | Specific dates within the next 30 days when normal business hours do not apply. | [optional] 

## Methods

### NewPreferredDeliveryTime

`func NewPreferredDeliveryTime() *PreferredDeliveryTime`

NewPreferredDeliveryTime instantiates a new PreferredDeliveryTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferredDeliveryTimeWithDefaults

`func NewPreferredDeliveryTimeWithDefaults() *PreferredDeliveryTime`

NewPreferredDeliveryTimeWithDefaults instantiates a new PreferredDeliveryTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessHours

`func (o *PreferredDeliveryTime) GetBusinessHours() []BusinessHour`

GetBusinessHours returns the BusinessHours field if non-nil, zero value otherwise.

### GetBusinessHoursOk

`func (o *PreferredDeliveryTime) GetBusinessHoursOk() (*[]BusinessHour, bool)`

GetBusinessHoursOk returns a tuple with the BusinessHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHours

`func (o *PreferredDeliveryTime) SetBusinessHours(v []BusinessHour)`

SetBusinessHours sets BusinessHours field to given value.

### HasBusinessHours

`func (o *PreferredDeliveryTime) HasBusinessHours() bool`

HasBusinessHours returns a boolean if a field has been set.

### GetExceptionDates

`func (o *PreferredDeliveryTime) GetExceptionDates() []ExceptionDate`

GetExceptionDates returns the ExceptionDates field if non-nil, zero value otherwise.

### GetExceptionDatesOk

`func (o *PreferredDeliveryTime) GetExceptionDatesOk() (*[]ExceptionDate, bool)`

GetExceptionDatesOk returns a tuple with the ExceptionDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionDates

`func (o *PreferredDeliveryTime) SetExceptionDates(v []ExceptionDate)`

SetExceptionDates sets ExceptionDates field to given value.

### HasExceptionDates

`func (o *PreferredDeliveryTime) HasExceptionDates() bool`

HasExceptionDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


