# Dimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Unit of measurement for the dimension value. | [optional] 
**Value** | Pointer to **float32** | Numeric value of the dimension. | [optional] 

## Methods

### NewDimension

`func NewDimension() *Dimension`

NewDimension instantiates a new Dimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionWithDefaults

`func NewDimensionWithDefaults() *Dimension`

NewDimensionWithDefaults instantiates a new Dimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Dimension) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Dimension) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Dimension) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Dimension) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *Dimension) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Dimension) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Dimension) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Dimension) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


