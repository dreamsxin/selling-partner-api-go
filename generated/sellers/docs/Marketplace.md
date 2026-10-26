# Marketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The encrypted marketplace value. | 
**Name** | **string** | The marketplace name. | 
**CountryCode** | **string** | The ISO 3166-1 alpha-2 format country code of the marketplace. | 
**DefaultCurrencyCode** | **string** | The ISO 4217 format currency code of the marketplace. | 
**DefaultLanguageCode** | **string** | The ISO 639-1 format language code of the marketplace. | 
**DomainName** | **string** | The domain name of the marketplace. | 

## Methods

### NewMarketplace

`func NewMarketplace(id string, name string, countryCode string, defaultCurrencyCode string, defaultLanguageCode string, domainName string, ) *Marketplace`

NewMarketplace instantiates a new Marketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceWithDefaults

`func NewMarketplaceWithDefaults() *Marketplace`

NewMarketplaceWithDefaults instantiates a new Marketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Marketplace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Marketplace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Marketplace) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Marketplace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Marketplace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Marketplace) SetName(v string)`

SetName sets Name field to given value.


### GetCountryCode

`func (o *Marketplace) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Marketplace) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Marketplace) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetDefaultCurrencyCode

`func (o *Marketplace) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *Marketplace) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *Marketplace) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.


### GetDefaultLanguageCode

`func (o *Marketplace) GetDefaultLanguageCode() string`

GetDefaultLanguageCode returns the DefaultLanguageCode field if non-nil, zero value otherwise.

### GetDefaultLanguageCodeOk

`func (o *Marketplace) GetDefaultLanguageCodeOk() (*string, bool)`

GetDefaultLanguageCodeOk returns a tuple with the DefaultLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageCode

`func (o *Marketplace) SetDefaultLanguageCode(v string)`

SetDefaultLanguageCode sets DefaultLanguageCode field to given value.


### GetDomainName

`func (o *Marketplace) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Marketplace) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Marketplace) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


