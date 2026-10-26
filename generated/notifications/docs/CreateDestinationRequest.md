# CreateDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSpecification** | [**DestinationResourceSpecification**](DestinationResourceSpecification.md) |  | 
**Name** | **string** | A developer-defined name to help identify this destination. | 

## Methods

### NewCreateDestinationRequest

`func NewCreateDestinationRequest(resourceSpecification DestinationResourceSpecification, name string, ) *CreateDestinationRequest`

NewCreateDestinationRequest instantiates a new CreateDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationRequestWithDefaults

`func NewCreateDestinationRequestWithDefaults() *CreateDestinationRequest`

NewCreateDestinationRequestWithDefaults instantiates a new CreateDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSpecification

`func (o *CreateDestinationRequest) GetResourceSpecification() DestinationResourceSpecification`

GetResourceSpecification returns the ResourceSpecification field if non-nil, zero value otherwise.

### GetResourceSpecificationOk

`func (o *CreateDestinationRequest) GetResourceSpecificationOk() (*DestinationResourceSpecification, bool)`

GetResourceSpecificationOk returns a tuple with the ResourceSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpecification

`func (o *CreateDestinationRequest) SetResourceSpecification(v DestinationResourceSpecification)`

SetResourceSpecification sets ResourceSpecification field to given value.


### GetName

`func (o *CreateDestinationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDestinationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDestinationRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


