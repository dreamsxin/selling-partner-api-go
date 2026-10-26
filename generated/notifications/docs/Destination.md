# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The developer-defined name for this destination. | 
**DestinationId** | **string** | The destination identifier generated when you created the destination. | 
**Resource** | [**DestinationResource**](DestinationResource.md) |  | 

## Methods

### NewDestination

`func NewDestination(name string, destinationId string, resource DestinationResource, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Destination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Destination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Destination) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationId

`func (o *Destination) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *Destination) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *Destination) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetResource

`func (o *Destination) GetResource() DestinationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Destination) GetResourceOk() (*DestinationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Destination) SetResource(v DestinationResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


