# TransitTimeInputDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | The carrier name. For example, &#x60;Correios&#x60;. | [optional] 
**ShippingService** | Pointer to **string** | The shipping service. For example, &#x60;SEDEX&#x60;. | [optional] 

## Methods

### NewTransitTimeInputDetails

`func NewTransitTimeInputDetails() *TransitTimeInputDetails`

NewTransitTimeInputDetails instantiates a new TransitTimeInputDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitTimeInputDetailsWithDefaults

`func NewTransitTimeInputDetailsWithDefaults() *TransitTimeInputDetails`

NewTransitTimeInputDetailsWithDefaults instantiates a new TransitTimeInputDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *TransitTimeInputDetails) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *TransitTimeInputDetails) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *TransitTimeInputDetails) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *TransitTimeInputDetails) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetShippingService

`func (o *TransitTimeInputDetails) GetShippingService() string`

GetShippingService returns the ShippingService field if non-nil, zero value otherwise.

### GetShippingServiceOk

`func (o *TransitTimeInputDetails) GetShippingServiceOk() (*string, bool)`

GetShippingServiceOk returns a tuple with the ShippingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingService

`func (o *TransitTimeInputDetails) SetShippingService(v string)`

SetShippingService sets ShippingService field to given value.

### HasShippingService

`func (o *TransitTimeInputDetails) HasShippingService() bool`

HasShippingService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


