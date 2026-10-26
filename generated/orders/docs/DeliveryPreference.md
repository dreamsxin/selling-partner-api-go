# DeliveryPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropOffLocation** | Pointer to **string** | The drop-off location selected by the customer. | [optional] 
**AddressInstruction** | Pointer to **string** | Building instructions, nearby landmark, or navigation instructions. | [optional] 
**DeliveryTime** | Pointer to [**PreferredDeliveryTime**](PreferredDeliveryTime.md) |  | [optional] 
**DeliveryCapabilities** | Pointer to **[]string** | A list of miscellaneous delivery capabilities associated with the shipping address. | [optional] 

## Methods

### NewDeliveryPreference

`func NewDeliveryPreference() *DeliveryPreference`

NewDeliveryPreference instantiates a new DeliveryPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryPreferenceWithDefaults

`func NewDeliveryPreferenceWithDefaults() *DeliveryPreference`

NewDeliveryPreferenceWithDefaults instantiates a new DeliveryPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropOffLocation

`func (o *DeliveryPreference) GetDropOffLocation() string`

GetDropOffLocation returns the DropOffLocation field if non-nil, zero value otherwise.

### GetDropOffLocationOk

`func (o *DeliveryPreference) GetDropOffLocationOk() (*string, bool)`

GetDropOffLocationOk returns a tuple with the DropOffLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropOffLocation

`func (o *DeliveryPreference) SetDropOffLocation(v string)`

SetDropOffLocation sets DropOffLocation field to given value.

### HasDropOffLocation

`func (o *DeliveryPreference) HasDropOffLocation() bool`

HasDropOffLocation returns a boolean if a field has been set.

### GetAddressInstruction

`func (o *DeliveryPreference) GetAddressInstruction() string`

GetAddressInstruction returns the AddressInstruction field if non-nil, zero value otherwise.

### GetAddressInstructionOk

`func (o *DeliveryPreference) GetAddressInstructionOk() (*string, bool)`

GetAddressInstructionOk returns a tuple with the AddressInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInstruction

`func (o *DeliveryPreference) SetAddressInstruction(v string)`

SetAddressInstruction sets AddressInstruction field to given value.

### HasAddressInstruction

`func (o *DeliveryPreference) HasAddressInstruction() bool`

HasAddressInstruction returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *DeliveryPreference) GetDeliveryTime() PreferredDeliveryTime`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *DeliveryPreference) GetDeliveryTimeOk() (*PreferredDeliveryTime, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *DeliveryPreference) SetDeliveryTime(v PreferredDeliveryTime)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *DeliveryPreference) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetDeliveryCapabilities

`func (o *DeliveryPreference) GetDeliveryCapabilities() []string`

GetDeliveryCapabilities returns the DeliveryCapabilities field if non-nil, zero value otherwise.

### GetDeliveryCapabilitiesOk

`func (o *DeliveryPreference) GetDeliveryCapabilitiesOk() (*[]string, bool)`

GetDeliveryCapabilitiesOk returns a tuple with the DeliveryCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCapabilities

`func (o *DeliveryPreference) SetDeliveryCapabilities(v []string)`

SetDeliveryCapabilities sets DeliveryCapabilities field to given value.

### HasDeliveryCapabilities

`func (o *DeliveryPreference) HasDeliveryCapabilities() bool`

HasDeliveryCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


