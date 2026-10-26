# OrderTaxRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** | The type of entity that the tax registration belongs to.  **Possible values**: - &#x60;BUYER&#x60; (Indicates that this is the buyer&#39;s tax registration information) - &#x60;MERCHANT&#x60; (Indicates that this is the merchant&#39;s tax registration information) - &#x60;MARKETPLACE&#x60; (Indicates that this is the marketplace&#39;s tax registration information) | [optional] 
**LegalName** | Pointer to **string** | The legal name associated with the tax registration. | [optional] 
**TaxRegistrationType** | Pointer to **string** | The type of the tax registration number.  **Possible values**: &#x60;BUSINESS&#x60;, &#x60;VAT&#x60;, &#x60;CST&#x60;, &#x60;CPF&#x60;, &#x60;CNPJ&#x60; | [optional] 
**TaxRegistrationNumber** | Pointer to **string** | The tax registration number that identifies the entity for tax purposes. | [optional] 
**TaxRegistrationAddress** | Pointer to [**CustomerAddress**](CustomerAddress.md) |  | [optional] 
**TaxRegistrationAttributes** | Pointer to [**[]TaxRegistrationAttribute**](TaxRegistrationAttribute.md) | Additional attributes related to the tax registration. | [optional] 

## Methods

### NewOrderTaxRegistration

`func NewOrderTaxRegistration() *OrderTaxRegistration`

NewOrderTaxRegistration instantiates a new OrderTaxRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaxRegistrationWithDefaults

`func NewOrderTaxRegistrationWithDefaults() *OrderTaxRegistration`

NewOrderTaxRegistrationWithDefaults instantiates a new OrderTaxRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *OrderTaxRegistration) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *OrderTaxRegistration) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *OrderTaxRegistration) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *OrderTaxRegistration) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetLegalName

`func (o *OrderTaxRegistration) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OrderTaxRegistration) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OrderTaxRegistration) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *OrderTaxRegistration) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxRegistrationType

`func (o *OrderTaxRegistration) GetTaxRegistrationType() string`

GetTaxRegistrationType returns the TaxRegistrationType field if non-nil, zero value otherwise.

### GetTaxRegistrationTypeOk

`func (o *OrderTaxRegistration) GetTaxRegistrationTypeOk() (*string, bool)`

GetTaxRegistrationTypeOk returns a tuple with the TaxRegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationType

`func (o *OrderTaxRegistration) SetTaxRegistrationType(v string)`

SetTaxRegistrationType sets TaxRegistrationType field to given value.

### HasTaxRegistrationType

`func (o *OrderTaxRegistration) HasTaxRegistrationType() bool`

HasTaxRegistrationType returns a boolean if a field has been set.

### GetTaxRegistrationNumber

`func (o *OrderTaxRegistration) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *OrderTaxRegistration) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *OrderTaxRegistration) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *OrderTaxRegistration) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### GetTaxRegistrationAddress

`func (o *OrderTaxRegistration) GetTaxRegistrationAddress() CustomerAddress`

GetTaxRegistrationAddress returns the TaxRegistrationAddress field if non-nil, zero value otherwise.

### GetTaxRegistrationAddressOk

`func (o *OrderTaxRegistration) GetTaxRegistrationAddressOk() (*CustomerAddress, bool)`

GetTaxRegistrationAddressOk returns a tuple with the TaxRegistrationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationAddress

`func (o *OrderTaxRegistration) SetTaxRegistrationAddress(v CustomerAddress)`

SetTaxRegistrationAddress sets TaxRegistrationAddress field to given value.

### HasTaxRegistrationAddress

`func (o *OrderTaxRegistration) HasTaxRegistrationAddress() bool`

HasTaxRegistrationAddress returns a boolean if a field has been set.

### GetTaxRegistrationAttributes

`func (o *OrderTaxRegistration) GetTaxRegistrationAttributes() []TaxRegistrationAttribute`

GetTaxRegistrationAttributes returns the TaxRegistrationAttributes field if non-nil, zero value otherwise.

### GetTaxRegistrationAttributesOk

`func (o *OrderTaxRegistration) GetTaxRegistrationAttributesOk() (*[]TaxRegistrationAttribute, bool)`

GetTaxRegistrationAttributesOk returns a tuple with the TaxRegistrationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationAttributes

`func (o *OrderTaxRegistration) SetTaxRegistrationAttributes(v []TaxRegistrationAttribute)`

SetTaxRegistrationAttributes sets TaxRegistrationAttributes field to given value.

### HasTaxRegistrationAttributes

`func (o *OrderTaxRegistration) HasTaxRegistrationAttributes() bool`

HasTaxRegistrationAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


