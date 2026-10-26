# ResearchingQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResearchingQuantity** | Pointer to **int32** | The total number of units currently being researched in Amazon&#39;s fulfillment network. | [optional] 
**ResearchingQuantityBreakdown** | Pointer to [**[]ResearchingQuantityEntry**](ResearchingQuantityEntry.md) | A list of quantity details for items currently being researched. | [optional] 

## Methods

### NewResearchingQuantity

`func NewResearchingQuantity() *ResearchingQuantity`

NewResearchingQuantity instantiates a new ResearchingQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchingQuantityWithDefaults

`func NewResearchingQuantityWithDefaults() *ResearchingQuantity`

NewResearchingQuantityWithDefaults instantiates a new ResearchingQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResearchingQuantity

`func (o *ResearchingQuantity) GetTotalResearchingQuantity() int32`

GetTotalResearchingQuantity returns the TotalResearchingQuantity field if non-nil, zero value otherwise.

### GetTotalResearchingQuantityOk

`func (o *ResearchingQuantity) GetTotalResearchingQuantityOk() (*int32, bool)`

GetTotalResearchingQuantityOk returns a tuple with the TotalResearchingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResearchingQuantity

`func (o *ResearchingQuantity) SetTotalResearchingQuantity(v int32)`

SetTotalResearchingQuantity sets TotalResearchingQuantity field to given value.

### HasTotalResearchingQuantity

`func (o *ResearchingQuantity) HasTotalResearchingQuantity() bool`

HasTotalResearchingQuantity returns a boolean if a field has been set.

### GetResearchingQuantityBreakdown

`func (o *ResearchingQuantity) GetResearchingQuantityBreakdown() []ResearchingQuantityEntry`

GetResearchingQuantityBreakdown returns the ResearchingQuantityBreakdown field if non-nil, zero value otherwise.

### GetResearchingQuantityBreakdownOk

`func (o *ResearchingQuantity) GetResearchingQuantityBreakdownOk() (*[]ResearchingQuantityEntry, bool)`

GetResearchingQuantityBreakdownOk returns a tuple with the ResearchingQuantityBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchingQuantityBreakdown

`func (o *ResearchingQuantity) SetResearchingQuantityBreakdown(v []ResearchingQuantityEntry)`

SetResearchingQuantityBreakdown sets ResearchingQuantityBreakdown field to given value.

### HasResearchingQuantityBreakdown

`func (o *ResearchingQuantity) HasResearchingQuantityBreakdown() bool`

HasResearchingQuantityBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


