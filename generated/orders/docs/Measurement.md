# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | The specific unit of measurement used to quantify this item. | 
**Value** | **float32** | The numerical quantity or amount being measured in the specified unit. | 

## Methods

### NewMeasurement

`func NewMeasurement(unit string, value float32, ) *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Measurement) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Measurement) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Measurement) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetValue

`func (o *Measurement) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Measurement) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Measurement) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


