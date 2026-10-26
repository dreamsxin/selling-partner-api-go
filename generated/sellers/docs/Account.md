# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceParticipationList** | [**[]MarketplaceParticipation**](MarketplaceParticipation.md) | List of marketplace participations. | 
**BusinessType** | **string** | The type of business registered for the seller account. | 
**SellingPlan** | **string** | The selling plan details. | 
**Business** | Pointer to [**Business**](Business.md) |  | [optional] 
**PrimaryContact** | Pointer to [**PrimaryContact**](PrimaryContact.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount(marketplaceParticipationList []MarketplaceParticipation, businessType string, sellingPlan string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceParticipationList

`func (o *Account) GetMarketplaceParticipationList() []MarketplaceParticipation`

GetMarketplaceParticipationList returns the MarketplaceParticipationList field if non-nil, zero value otherwise.

### GetMarketplaceParticipationListOk

`func (o *Account) GetMarketplaceParticipationListOk() (*[]MarketplaceParticipation, bool)`

GetMarketplaceParticipationListOk returns a tuple with the MarketplaceParticipationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceParticipationList

`func (o *Account) SetMarketplaceParticipationList(v []MarketplaceParticipation)`

SetMarketplaceParticipationList sets MarketplaceParticipationList field to given value.


### GetBusinessType

`func (o *Account) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *Account) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *Account) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.


### GetSellingPlan

`func (o *Account) GetSellingPlan() string`

GetSellingPlan returns the SellingPlan field if non-nil, zero value otherwise.

### GetSellingPlanOk

`func (o *Account) GetSellingPlanOk() (*string, bool)`

GetSellingPlanOk returns a tuple with the SellingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPlan

`func (o *Account) SetSellingPlan(v string)`

SetSellingPlan sets SellingPlan field to given value.


### GetBusiness

`func (o *Account) GetBusiness() Business`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *Account) GetBusinessOk() (*Business, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *Account) SetBusiness(v Business)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *Account) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *Account) GetPrimaryContact() PrimaryContact`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *Account) GetPrimaryContactOk() (*PrimaryContact, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *Account) SetPrimaryContact(v PrimaryContact)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *Account) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


