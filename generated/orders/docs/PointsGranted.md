# PointsGranted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PointsNumber** | Pointer to **int32** | Total number of Amazon Points granted to the customer&#39;s account for this item purchase. | [optional] 
**PointsMonetaryValue** | Pointer to [**Money**](Money.md) |  | [optional] 

## Methods

### NewPointsGranted

`func NewPointsGranted() *PointsGranted`

NewPointsGranted instantiates a new PointsGranted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointsGrantedWithDefaults

`func NewPointsGrantedWithDefaults() *PointsGranted`

NewPointsGrantedWithDefaults instantiates a new PointsGranted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointsNumber

`func (o *PointsGranted) GetPointsNumber() int32`

GetPointsNumber returns the PointsNumber field if non-nil, zero value otherwise.

### GetPointsNumberOk

`func (o *PointsGranted) GetPointsNumberOk() (*int32, bool)`

GetPointsNumberOk returns a tuple with the PointsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsNumber

`func (o *PointsGranted) SetPointsNumber(v int32)`

SetPointsNumber sets PointsNumber field to given value.

### HasPointsNumber

`func (o *PointsGranted) HasPointsNumber() bool`

HasPointsNumber returns a boolean if a field has been set.

### GetPointsMonetaryValue

`func (o *PointsGranted) GetPointsMonetaryValue() Money`

GetPointsMonetaryValue returns the PointsMonetaryValue field if non-nil, zero value otherwise.

### GetPointsMonetaryValueOk

`func (o *PointsGranted) GetPointsMonetaryValueOk() (*Money, bool)`

GetPointsMonetaryValueOk returns a tuple with the PointsMonetaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsMonetaryValue

`func (o *PointsGranted) SetPointsMonetaryValue(v Money)`

SetPointsMonetaryValue sets PointsMonetaryValue field to given value.

### HasPointsMonetaryValue

`func (o *PointsGranted) HasPointsMonetaryValue() bool`

HasPointsMonetaryValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


