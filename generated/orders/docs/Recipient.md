# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryAddress** | Pointer to [**CustomerAddress**](CustomerAddress.md) |  | [optional] 
**DeliveryPreference** | Pointer to [**DeliveryPreference**](DeliveryPreference.md) |  | [optional] 

## Methods

### NewRecipient

`func NewRecipient() *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryAddress

`func (o *Recipient) GetDeliveryAddress() CustomerAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *Recipient) GetDeliveryAddressOk() (*CustomerAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *Recipient) SetDeliveryAddress(v CustomerAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *Recipient) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryPreference

`func (o *Recipient) GetDeliveryPreference() DeliveryPreference`

GetDeliveryPreference returns the DeliveryPreference field if non-nil, zero value otherwise.

### GetDeliveryPreferenceOk

`func (o *Recipient) GetDeliveryPreferenceOk() (*DeliveryPreference, bool)`

GetDeliveryPreferenceOk returns a tuple with the DeliveryPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPreference

`func (o *Recipient) SetDeliveryPreference(v DeliveryPreference)`

SetDeliveryPreference sets DeliveryPreference field to given value.

### HasDeliveryPreference

`func (o *Recipient) HasDeliveryPreference() bool`

HasDeliveryPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


