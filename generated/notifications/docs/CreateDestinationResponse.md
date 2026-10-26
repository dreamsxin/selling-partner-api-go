# CreateDestinationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**Destination**](Destination.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | A list of error responses returned when a request is unsuccessful. | [optional] 

## Methods

### NewCreateDestinationResponse

`func NewCreateDestinationResponse() *CreateDestinationResponse`

NewCreateDestinationResponse instantiates a new CreateDestinationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationResponseWithDefaults

`func NewCreateDestinationResponseWithDefaults() *CreateDestinationResponse`

NewCreateDestinationResponseWithDefaults instantiates a new CreateDestinationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *CreateDestinationResponse) GetPayload() Destination`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateDestinationResponse) GetPayloadOk() (*Destination, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateDestinationResponse) SetPayload(v Destination)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CreateDestinationResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetErrors

`func (o *CreateDestinationResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateDestinationResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateDestinationResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateDestinationResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


