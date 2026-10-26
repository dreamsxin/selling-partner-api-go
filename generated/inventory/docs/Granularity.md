# Granularity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GranularityType** | Pointer to **string** | The granularity type for the inventory aggregation level. | [optional] 
**GranularityId** | Pointer to **string** | The granularity ID for the specified granularity type. When granularityType is Marketplace, specify the marketplaceId. | [optional] 

## Methods

### NewGranularity

`func NewGranularity() *Granularity`

NewGranularity instantiates a new Granularity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGranularityWithDefaults

`func NewGranularityWithDefaults() *Granularity`

NewGranularityWithDefaults instantiates a new Granularity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularityType

`func (o *Granularity) GetGranularityType() string`

GetGranularityType returns the GranularityType field if non-nil, zero value otherwise.

### GetGranularityTypeOk

`func (o *Granularity) GetGranularityTypeOk() (*string, bool)`

GetGranularityTypeOk returns a tuple with the GranularityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityType

`func (o *Granularity) SetGranularityType(v string)`

SetGranularityType sets GranularityType field to given value.

### HasGranularityType

`func (o *Granularity) HasGranularityType() bool`

HasGranularityType returns a boolean if a field has been set.

### GetGranularityId

`func (o *Granularity) GetGranularityId() string`

GetGranularityId returns the GranularityId field if non-nil, zero value otherwise.

### GetGranularityIdOk

`func (o *Granularity) GetGranularityIdOk() (*string, bool)`

GetGranularityIdOk returns a tuple with the GranularityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityId

`func (o *Granularity) SetGranularityId(v string)`

SetGranularityId sets GranularityId field to given value.

### HasGranularityId

`func (o *Granularity) HasGranularityId() bool`

HasGranularityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


