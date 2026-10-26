# CustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full name of the person who will receive the delivery at this address. | [optional] 
**CompanyName** | Pointer to **string** | The name of the business or organization at this address. | [optional] 
**AddressLine1** | Pointer to **string** | The primary street address. | [optional] 
**AddressLine2** | Pointer to **string** | Additional address information. | [optional] 
**AddressLine3** | Pointer to **string** | Additional address information. | [optional] 
**City** | Pointer to **string** | The city of the address. | [optional] 
**DistrictOrCounty** | Pointer to **string** | The district or county of the address. | [optional] 
**StateOrRegion** | Pointer to **string** | The state, province, or region of the address. | [optional] 
**Municipality** | Pointer to **string** | The municipality of the address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code, ZIP code, or equivalent mailing code of the address. | [optional] 
**CountryCode** | Pointer to **string** | The two-letter country code identifying the country of the address, in ISO 3166-1 format. | [optional] 
**Phone** | Pointer to **string** | The contact phone number for delivery coordination and customer communication.  **Note**: The buyer phone number will be suppressed in some cases, including but not limited to:  - All FBA (Fulfillment by Amazon) orders. - Shipped FBM (Fulfilled by the merchant) orders when current date is past the latest delivery date. | [optional] 
**ExtendedFields** | Pointer to [**AddressExtendedFields**](AddressExtendedFields.md) |  | [optional] 
**AddressType** | Pointer to **string** | The type of location.  **Possible values**: &#x60;RESIDENTIAL&#x60;, &#x60;COMMERCIAL&#x60;, &#x60;PICKUP_POINT&#x60; | [optional] 

## Methods

### NewCustomerAddress

`func NewCustomerAddress() *CustomerAddress`

NewCustomerAddress instantiates a new CustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressWithDefaults

`func NewCustomerAddressWithDefaults() *CustomerAddress`

NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerAddress) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerAddress) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerAddress) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerAddress) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CustomerAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomerAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomerAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomerAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CustomerAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomerAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomerAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomerAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *CustomerAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *CustomerAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *CustomerAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *CustomerAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *CustomerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrictOrCounty

`func (o *CustomerAddress) GetDistrictOrCounty() string`

GetDistrictOrCounty returns the DistrictOrCounty field if non-nil, zero value otherwise.

### GetDistrictOrCountyOk

`func (o *CustomerAddress) GetDistrictOrCountyOk() (*string, bool)`

GetDistrictOrCountyOk returns a tuple with the DistrictOrCounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictOrCounty

`func (o *CustomerAddress) SetDistrictOrCounty(v string)`

SetDistrictOrCounty sets DistrictOrCounty field to given value.

### HasDistrictOrCounty

`func (o *CustomerAddress) HasDistrictOrCounty() bool`

HasDistrictOrCounty returns a boolean if a field has been set.

### GetStateOrRegion

`func (o *CustomerAddress) GetStateOrRegion() string`

GetStateOrRegion returns the StateOrRegion field if non-nil, zero value otherwise.

### GetStateOrRegionOk

`func (o *CustomerAddress) GetStateOrRegionOk() (*string, bool)`

GetStateOrRegionOk returns a tuple with the StateOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrRegion

`func (o *CustomerAddress) SetStateOrRegion(v string)`

SetStateOrRegion sets StateOrRegion field to given value.

### HasStateOrRegion

`func (o *CustomerAddress) HasStateOrRegion() bool`

HasStateOrRegion returns a boolean if a field has been set.

### GetMunicipality

`func (o *CustomerAddress) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *CustomerAddress) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *CustomerAddress) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *CustomerAddress) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomerAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *CustomerAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CustomerAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CustomerAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CustomerAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetExtendedFields

`func (o *CustomerAddress) GetExtendedFields() AddressExtendedFields`

GetExtendedFields returns the ExtendedFields field if non-nil, zero value otherwise.

### GetExtendedFieldsOk

`func (o *CustomerAddress) GetExtendedFieldsOk() (*AddressExtendedFields, bool)`

GetExtendedFieldsOk returns a tuple with the ExtendedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFields

`func (o *CustomerAddress) SetExtendedFields(v AddressExtendedFields)`

SetExtendedFields sets ExtendedFields field to given value.

### HasExtendedFields

`func (o *CustomerAddress) HasExtendedFields() bool`

HasExtendedFields returns a boolean if a field has been set.

### GetAddressType

`func (o *CustomerAddress) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *CustomerAddress) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *CustomerAddress) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *CustomerAddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


