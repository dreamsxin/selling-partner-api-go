# Refinements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | [**[]BrandRefinement**](BrandRefinement.md) | A list of brands you can use to refine your search. | 
**Classifications** | [**[]ClassificationRefinement**](ClassificationRefinement.md) | A list of classifications you can use to refine your search. | 

## Methods

### NewRefinements

`func NewRefinements(brands []BrandRefinement, classifications []ClassificationRefinement, ) *Refinements`

NewRefinements instantiates a new Refinements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefinementsWithDefaults

`func NewRefinementsWithDefaults() *Refinements`

NewRefinementsWithDefaults instantiates a new Refinements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *Refinements) GetBrands() []BrandRefinement`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *Refinements) GetBrandsOk() (*[]BrandRefinement, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *Refinements) SetBrands(v []BrandRefinement)`

SetBrands sets Brands field to given value.


### GetClassifications

`func (o *Refinements) GetClassifications() []ClassificationRefinement`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *Refinements) GetClassificationsOk() (*[]ClassificationRefinement, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *Refinements) SetClassifications(v []ClassificationRefinement)`

SetClassifications sets Classifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


