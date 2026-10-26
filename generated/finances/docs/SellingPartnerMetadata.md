# SellingPartnerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellingPartnerId** | Pointer to **string** | The unique seller identifier. | [optional] 
**AccountType** | Pointer to **string** | The type of account in the transaction. | [optional] 
**MarketplaceId** | Pointer to **string** | The identifier of the marketplace where the transaction occurred. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for a marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | [optional] 

## Methods

### NewSellingPartnerMetadata

`func NewSellingPartnerMetadata() *SellingPartnerMetadata`

NewSellingPartnerMetadata instantiates a new SellingPartnerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellingPartnerMetadataWithDefaults

`func NewSellingPartnerMetadataWithDefaults() *SellingPartnerMetadata`

NewSellingPartnerMetadataWithDefaults instantiates a new SellingPartnerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellingPartnerId

`func (o *SellingPartnerMetadata) GetSellingPartnerId() string`

GetSellingPartnerId returns the SellingPartnerId field if non-nil, zero value otherwise.

### GetSellingPartnerIdOk

`func (o *SellingPartnerMetadata) GetSellingPartnerIdOk() (*string, bool)`

GetSellingPartnerIdOk returns a tuple with the SellingPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPartnerId

`func (o *SellingPartnerMetadata) SetSellingPartnerId(v string)`

SetSellingPartnerId sets SellingPartnerId field to given value.

### HasSellingPartnerId

`func (o *SellingPartnerMetadata) HasSellingPartnerId() bool`

HasSellingPartnerId returns a boolean if a field has been set.

### GetAccountType

`func (o *SellingPartnerMetadata) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SellingPartnerMetadata) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SellingPartnerMetadata) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SellingPartnerMetadata) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *SellingPartnerMetadata) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *SellingPartnerMetadata) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *SellingPartnerMetadata) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *SellingPartnerMetadata) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


