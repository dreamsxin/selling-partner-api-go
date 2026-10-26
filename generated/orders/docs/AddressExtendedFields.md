# AddressExtendedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetName** | Pointer to **string** | The name of the street. | [optional] 
**StreetNumber** | Pointer to **string** | The house, building, or property number associated with the location&#39;s street address. | [optional] 
**Complement** | Pointer to **string** | The floor number / unit number. | [optional] 
**Neighborhood** | Pointer to **string** | The neighborhood. This value is only used in some countries (such as Brazil). | [optional] 

## Methods

### NewAddressExtendedFields

`func NewAddressExtendedFields() *AddressExtendedFields`

NewAddressExtendedFields instantiates a new AddressExtendedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressExtendedFieldsWithDefaults

`func NewAddressExtendedFieldsWithDefaults() *AddressExtendedFields`

NewAddressExtendedFieldsWithDefaults instantiates a new AddressExtendedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetName

`func (o *AddressExtendedFields) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *AddressExtendedFields) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *AddressExtendedFields) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *AddressExtendedFields) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetStreetNumber

`func (o *AddressExtendedFields) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *AddressExtendedFields) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *AddressExtendedFields) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *AddressExtendedFields) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetComplement

`func (o *AddressExtendedFields) GetComplement() string`

GetComplement returns the Complement field if non-nil, zero value otherwise.

### GetComplementOk

`func (o *AddressExtendedFields) GetComplementOk() (*string, bool)`

GetComplementOk returns a tuple with the Complement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplement

`func (o *AddressExtendedFields) SetComplement(v string)`

SetComplement sets Complement field to given value.

### HasComplement

`func (o *AddressExtendedFields) HasComplement() bool`

HasComplement returns a boolean if a field has been set.

### GetNeighborhood

`func (o *AddressExtendedFields) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *AddressExtendedFields) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *AddressExtendedFields) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *AddressExtendedFields) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


