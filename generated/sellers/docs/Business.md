# Business

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The registered business name. | 
**RegisteredBusinessAddress** | [**Address**](Address.md) |  | 
**CompanyRegistrationNumber** | Pointer to **string** | The seller&#39;s company registration number, if applicable. This field will be absent for individual sellers and sole proprietorships. | [optional] 
**CompanyTaxIdentificationNumber** | Pointer to **string** | The seller&#39;s company tax identification number, if applicable. This field will be present for certain business types only, such as sole proprietorships. | [optional] 
**NonLatinName** | Pointer to **string** | The non-Latin script version of the registered business name, if applicable. | [optional] 

## Methods

### NewBusiness

`func NewBusiness(name string, registeredBusinessAddress Address, ) *Business`

NewBusiness instantiates a new Business object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessWithDefaults

`func NewBusinessWithDefaults() *Business`

NewBusinessWithDefaults instantiates a new Business object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Business) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Business) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Business) SetName(v string)`

SetName sets Name field to given value.


### GetRegisteredBusinessAddress

`func (o *Business) GetRegisteredBusinessAddress() Address`

GetRegisteredBusinessAddress returns the RegisteredBusinessAddress field if non-nil, zero value otherwise.

### GetRegisteredBusinessAddressOk

`func (o *Business) GetRegisteredBusinessAddressOk() (*Address, bool)`

GetRegisteredBusinessAddressOk returns a tuple with the RegisteredBusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredBusinessAddress

`func (o *Business) SetRegisteredBusinessAddress(v Address)`

SetRegisteredBusinessAddress sets RegisteredBusinessAddress field to given value.


### GetCompanyRegistrationNumber

`func (o *Business) GetCompanyRegistrationNumber() string`

GetCompanyRegistrationNumber returns the CompanyRegistrationNumber field if non-nil, zero value otherwise.

### GetCompanyRegistrationNumberOk

`func (o *Business) GetCompanyRegistrationNumberOk() (*string, bool)`

GetCompanyRegistrationNumberOk returns a tuple with the CompanyRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRegistrationNumber

`func (o *Business) SetCompanyRegistrationNumber(v string)`

SetCompanyRegistrationNumber sets CompanyRegistrationNumber field to given value.

### HasCompanyRegistrationNumber

`func (o *Business) HasCompanyRegistrationNumber() bool`

HasCompanyRegistrationNumber returns a boolean if a field has been set.

### GetCompanyTaxIdentificationNumber

`func (o *Business) GetCompanyTaxIdentificationNumber() string`

GetCompanyTaxIdentificationNumber returns the CompanyTaxIdentificationNumber field if non-nil, zero value otherwise.

### GetCompanyTaxIdentificationNumberOk

`func (o *Business) GetCompanyTaxIdentificationNumberOk() (*string, bool)`

GetCompanyTaxIdentificationNumberOk returns a tuple with the CompanyTaxIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTaxIdentificationNumber

`func (o *Business) SetCompanyTaxIdentificationNumber(v string)`

SetCompanyTaxIdentificationNumber sets CompanyTaxIdentificationNumber field to given value.

### HasCompanyTaxIdentificationNumber

`func (o *Business) HasCompanyTaxIdentificationNumber() bool`

HasCompanyTaxIdentificationNumber returns a boolean if a field has been set.

### GetNonLatinName

`func (o *Business) GetNonLatinName() string`

GetNonLatinName returns the NonLatinName field if non-nil, zero value otherwise.

### GetNonLatinNameOk

`func (o *Business) GetNonLatinNameOk() (*string, bool)`

GetNonLatinNameOk returns a tuple with the NonLatinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonLatinName

`func (o *Business) SetNonLatinName(v string)`

SetNonLatinName sets NonLatinName field to given value.

### HasNonLatinName

`func (o *Business) HasNonLatinName() bool`

HasNonLatinName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


