# BrandRefinement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfResults** | **int32** | The estimated number of results that would be returned if you refine your search by the specified &#x60;brandName&#x60;. | 
**BrandName** | **string** | The brand name that you can use to refine your search. | 

## Methods

### NewBrandRefinement

`func NewBrandRefinement(numberOfResults int32, brandName string, ) *BrandRefinement`

NewBrandRefinement instantiates a new BrandRefinement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandRefinementWithDefaults

`func NewBrandRefinementWithDefaults() *BrandRefinement`

NewBrandRefinementWithDefaults instantiates a new BrandRefinement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfResults

`func (o *BrandRefinement) GetNumberOfResults() int32`

GetNumberOfResults returns the NumberOfResults field if non-nil, zero value otherwise.

### GetNumberOfResultsOk

`func (o *BrandRefinement) GetNumberOfResultsOk() (*int32, bool)`

GetNumberOfResultsOk returns a tuple with the NumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResults

`func (o *BrandRefinement) SetNumberOfResults(v int32)`

SetNumberOfResults sets NumberOfResults field to given value.


### GetBrandName

`func (o *BrandRefinement) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *BrandRefinement) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *BrandRefinement) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


