# MerchantAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the business or facility at this address. | [optional] 
**AddressLine1** | Pointer to **string** | The primary street address of the merchant&#39;s location. | [optional] 
**AddressLine2** | Pointer to **string** | Additional address information. | [optional] 
**AddressLine3** | Pointer to **string** | Additional address information. | [optional] 
**City** | Pointer to **string** | The city of the address. | [optional] 
**DistrictOrCounty** | Pointer to **string** | The district or county of the address. | [optional] 
**StateOrRegion** | Pointer to **string** | The state, province, or region of the address. | [optional] 
**Municipality** | Pointer to **string** | The municipality of the address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the address. | [optional] 
**CountryCode** | Pointer to **string** | The two-letter country code for the address, in ISO 3166-1 format. | [optional] 

## Methods

### NewMerchantAddress

`func NewMerchantAddress() *MerchantAddress`

NewMerchantAddress instantiates a new MerchantAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAddressWithDefaults

`func NewMerchantAddressWithDefaults() *MerchantAddress`

NewMerchantAddressWithDefaults instantiates a new MerchantAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *MerchantAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *MerchantAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *MerchantAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *MerchantAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *MerchantAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *MerchantAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *MerchantAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *MerchantAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *MerchantAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *MerchantAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *MerchantAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *MerchantAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *MerchantAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MerchantAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MerchantAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MerchantAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrictOrCounty

`func (o *MerchantAddress) GetDistrictOrCounty() string`

GetDistrictOrCounty returns the DistrictOrCounty field if non-nil, zero value otherwise.

### GetDistrictOrCountyOk

`func (o *MerchantAddress) GetDistrictOrCountyOk() (*string, bool)`

GetDistrictOrCountyOk returns a tuple with the DistrictOrCounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictOrCounty

`func (o *MerchantAddress) SetDistrictOrCounty(v string)`

SetDistrictOrCounty sets DistrictOrCounty field to given value.

### HasDistrictOrCounty

`func (o *MerchantAddress) HasDistrictOrCounty() bool`

HasDistrictOrCounty returns a boolean if a field has been set.

### GetStateOrRegion

`func (o *MerchantAddress) GetStateOrRegion() string`

GetStateOrRegion returns the StateOrRegion field if non-nil, zero value otherwise.

### GetStateOrRegionOk

`func (o *MerchantAddress) GetStateOrRegionOk() (*string, bool)`

GetStateOrRegionOk returns a tuple with the StateOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrRegion

`func (o *MerchantAddress) SetStateOrRegion(v string)`

SetStateOrRegion sets StateOrRegion field to given value.

### HasStateOrRegion

`func (o *MerchantAddress) HasStateOrRegion() bool`

HasStateOrRegion returns a boolean if a field has been set.

### GetMunicipality

`func (o *MerchantAddress) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *MerchantAddress) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *MerchantAddress) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *MerchantAddress) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### GetPostalCode

`func (o *MerchantAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MerchantAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MerchantAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MerchantAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *MerchantAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MerchantAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MerchantAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MerchantAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


