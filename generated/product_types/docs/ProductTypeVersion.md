# ProductTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version identifier. | 
**Latest** | **bool** | When true, the version indicated by the version identifier is the latest available for the Amazon product type. | 
**ReleaseCandidate** | Pointer to **bool** | When true, the version indicated by the version identifier is the prerelease (release candidate) for the Amazon product type. | [optional] 

## Methods

### NewProductTypeVersion

`func NewProductTypeVersion(version string, latest bool, ) *ProductTypeVersion`

NewProductTypeVersion instantiates a new ProductTypeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeVersionWithDefaults

`func NewProductTypeVersionWithDefaults() *ProductTypeVersion`

NewProductTypeVersionWithDefaults instantiates a new ProductTypeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ProductTypeVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProductTypeVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProductTypeVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLatest

`func (o *ProductTypeVersion) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *ProductTypeVersion) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *ProductTypeVersion) SetLatest(v bool)`

SetLatest sets Latest field to given value.


### GetReleaseCandidate

`func (o *ProductTypeVersion) GetReleaseCandidate() bool`

GetReleaseCandidate returns the ReleaseCandidate field if non-nil, zero value otherwise.

### GetReleaseCandidateOk

`func (o *ProductTypeVersion) GetReleaseCandidateOk() (*bool, bool)`

GetReleaseCandidateOk returns a tuple with the ReleaseCandidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCandidate

`func (o *ProductTypeVersion) SetReleaseCandidate(v bool)`

SetReleaseCandidate sets ReleaseCandidate field to given value.

### HasReleaseCandidate

`func (o *ProductTypeVersion) HasReleaseCandidate() bool`

HasReleaseCandidate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


