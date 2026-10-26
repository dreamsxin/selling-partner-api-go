# TrackingIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Amazon unique tracking identifier. | [optional] 
**Acsin** | Pointer to **string** | Air Cargo Shipment Identification Number. | [optional] 
**Aftn** | Pointer to **string** | Amazon Fulfillment Tracking Number. | [optional] 
**ContainerNumber** | Pointer to **string** | Container number provided by the Logistics Service Provider. | [optional] 
**HouseBillOfLadingNumber** | Pointer to **string** | House Bill of Lading number. | [optional] 
**CarrierTracking** | Pointer to [**TrackingIdentifierCarrierTracking**](TrackingIdentifierCarrierTracking.md) |  | [optional] 

## Methods

### NewTrackingIdentifier

`func NewTrackingIdentifier() *TrackingIdentifier`

NewTrackingIdentifier instantiates a new TrackingIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingIdentifierWithDefaults

`func NewTrackingIdentifierWithDefaults() *TrackingIdentifier`

NewTrackingIdentifierWithDefaults instantiates a new TrackingIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackingIdentifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingIdentifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingIdentifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingIdentifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAcsin

`func (o *TrackingIdentifier) GetAcsin() string`

GetAcsin returns the Acsin field if non-nil, zero value otherwise.

### GetAcsinOk

`func (o *TrackingIdentifier) GetAcsinOk() (*string, bool)`

GetAcsinOk returns a tuple with the Acsin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsin

`func (o *TrackingIdentifier) SetAcsin(v string)`

SetAcsin sets Acsin field to given value.

### HasAcsin

`func (o *TrackingIdentifier) HasAcsin() bool`

HasAcsin returns a boolean if a field has been set.

### GetAftn

`func (o *TrackingIdentifier) GetAftn() string`

GetAftn returns the Aftn field if non-nil, zero value otherwise.

### GetAftnOk

`func (o *TrackingIdentifier) GetAftnOk() (*string, bool)`

GetAftnOk returns a tuple with the Aftn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAftn

`func (o *TrackingIdentifier) SetAftn(v string)`

SetAftn sets Aftn field to given value.

### HasAftn

`func (o *TrackingIdentifier) HasAftn() bool`

HasAftn returns a boolean if a field has been set.

### GetContainerNumber

`func (o *TrackingIdentifier) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *TrackingIdentifier) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *TrackingIdentifier) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *TrackingIdentifier) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### GetHouseBillOfLadingNumber

`func (o *TrackingIdentifier) GetHouseBillOfLadingNumber() string`

GetHouseBillOfLadingNumber returns the HouseBillOfLadingNumber field if non-nil, zero value otherwise.

### GetHouseBillOfLadingNumberOk

`func (o *TrackingIdentifier) GetHouseBillOfLadingNumberOk() (*string, bool)`

GetHouseBillOfLadingNumberOk returns a tuple with the HouseBillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseBillOfLadingNumber

`func (o *TrackingIdentifier) SetHouseBillOfLadingNumber(v string)`

SetHouseBillOfLadingNumber sets HouseBillOfLadingNumber field to given value.

### HasHouseBillOfLadingNumber

`func (o *TrackingIdentifier) HasHouseBillOfLadingNumber() bool`

HasHouseBillOfLadingNumber returns a boolean if a field has been set.

### GetCarrierTracking

`func (o *TrackingIdentifier) GetCarrierTracking() TrackingIdentifierCarrierTracking`

GetCarrierTracking returns the CarrierTracking field if non-nil, zero value otherwise.

### GetCarrierTrackingOk

`func (o *TrackingIdentifier) GetCarrierTrackingOk() (*TrackingIdentifierCarrierTracking, bool)`

GetCarrierTrackingOk returns a tuple with the CarrierTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierTracking

`func (o *TrackingIdentifier) SetCarrierTracking(v TrackingIdentifierCarrierTracking)`

SetCarrierTracking sets CarrierTracking field to given value.

### HasCarrierTracking

`func (o *TrackingIdentifier) HasCarrierTracking() bool`

HasCarrierTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


