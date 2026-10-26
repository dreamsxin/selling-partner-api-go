# ClassificationRefinement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfResults** | **int32** | The estimated number of results that would be returned if you refine your search by the specified &#x60;classificationId&#x60;. | 
**DisplayName** | **string** | Display name for the classification. | 
**ClassificationId** | **string** | The identifier of the classification that you can use to refine your search. | 

## Methods

### NewClassificationRefinement

`func NewClassificationRefinement(numberOfResults int32, displayName string, classificationId string, ) *ClassificationRefinement`

NewClassificationRefinement instantiates a new ClassificationRefinement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationRefinementWithDefaults

`func NewClassificationRefinementWithDefaults() *ClassificationRefinement`

NewClassificationRefinementWithDefaults instantiates a new ClassificationRefinement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfResults

`func (o *ClassificationRefinement) GetNumberOfResults() int32`

GetNumberOfResults returns the NumberOfResults field if non-nil, zero value otherwise.

### GetNumberOfResultsOk

`func (o *ClassificationRefinement) GetNumberOfResultsOk() (*int32, bool)`

GetNumberOfResultsOk returns a tuple with the NumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResults

`func (o *ClassificationRefinement) SetNumberOfResults(v int32)`

SetNumberOfResults sets NumberOfResults field to given value.


### GetDisplayName

`func (o *ClassificationRefinement) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClassificationRefinement) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClassificationRefinement) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetClassificationId

`func (o *ClassificationRefinement) GetClassificationId() string`

GetClassificationId returns the ClassificationId field if non-nil, zero value otherwise.

### GetClassificationIdOk

`func (o *ClassificationRefinement) GetClassificationIdOk() (*string, bool)`

GetClassificationIdOk returns a tuple with the ClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationId

`func (o *ClassificationRefinement) SetClassificationId(v string)`

SetClassificationId sets ClassificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


