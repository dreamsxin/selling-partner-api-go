# GetInventorySummariesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granularity** | [**Granularity**](Granularity.md) |  | 
**InventorySummaries** | [**[]InventorySummary**](InventorySummary.md) | A list of inventory summaries. | 

## Methods

### NewGetInventorySummariesResult

`func NewGetInventorySummariesResult(granularity Granularity, inventorySummaries []InventorySummary, ) *GetInventorySummariesResult`

NewGetInventorySummariesResult instantiates a new GetInventorySummariesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventorySummariesResultWithDefaults

`func NewGetInventorySummariesResultWithDefaults() *GetInventorySummariesResult`

NewGetInventorySummariesResultWithDefaults instantiates a new GetInventorySummariesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularity

`func (o *GetInventorySummariesResult) GetGranularity() Granularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GetInventorySummariesResult) GetGranularityOk() (*Granularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GetInventorySummariesResult) SetGranularity(v Granularity)`

SetGranularity sets Granularity field to given value.


### GetInventorySummaries

`func (o *GetInventorySummariesResult) GetInventorySummaries() []InventorySummary`

GetInventorySummaries returns the InventorySummaries field if non-nil, zero value otherwise.

### GetInventorySummariesOk

`func (o *GetInventorySummariesResult) GetInventorySummariesOk() (*[]InventorySummary, bool)`

GetInventorySummariesOk returns a tuple with the InventorySummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySummaries

`func (o *GetInventorySummariesResult) SetInventorySummaries(v []InventorySummary)`

SetInventorySummaries sets InventorySummaries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


