# PartnerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerId** | Pointer to **string** | The unique seller identifier. | [optional] 
**AccountType** | Pointer to **string** | The type of account. | [optional] 
**MarketplaceId** | Pointer to **string** | The identifier of the marketplace where the transaction occurred. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for a marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | [optional] 

## Methods

### NewPartnerMetadata

`func NewPartnerMetadata() *PartnerMetadata`

NewPartnerMetadata instantiates a new PartnerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerMetadataWithDefaults

`func NewPartnerMetadataWithDefaults() *PartnerMetadata`

NewPartnerMetadataWithDefaults instantiates a new PartnerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerId

`func (o *PartnerMetadata) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PartnerMetadata) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PartnerMetadata) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *PartnerMetadata) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetAccountType

`func (o *PartnerMetadata) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PartnerMetadata) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PartnerMetadata) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PartnerMetadata) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *PartnerMetadata) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *PartnerMetadata) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *PartnerMetadata) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *PartnerMetadata) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


