# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street address information. | 
**AddressLine2** | Pointer to **string** | Additional street address information. | [optional] 
**CountryCode** | **string** | The country code in two-character ISO 3166-1 alpha-2 format. | 
**StateOrProvinceCode** | Pointer to **string** | The state or province code. | [optional] 
**City** | Pointer to **string** | The city. | [optional] 
**PostalCode** | Pointer to **string** | The postal code. | [optional] 

## Methods

### NewAddress

`func NewAddress(addressLine1 string, countryCode string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *Address) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *Address) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Address) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCountryCode

`func (o *Address) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Address) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Address) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetStateOrProvinceCode

`func (o *Address) GetStateOrProvinceCode() string`

GetStateOrProvinceCode returns the StateOrProvinceCode field if non-nil, zero value otherwise.

### GetStateOrProvinceCodeOk

`func (o *Address) GetStateOrProvinceCodeOk() (*string, bool)`

GetStateOrProvinceCodeOk returns a tuple with the StateOrProvinceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvinceCode

`func (o *Address) SetStateOrProvinceCode(v string)`

SetStateOrProvinceCode sets StateOrProvinceCode field to given value.

### HasStateOrProvinceCode

`func (o *Address) HasStateOrProvinceCode() bool`

HasStateOrProvinceCode returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


