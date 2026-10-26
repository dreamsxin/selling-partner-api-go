# ItemIdentifiersByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).identifier. | 
**Identifiers** | [**[]ItemIdentifier**](ItemIdentifier.md) | Identifiers associated with the item in the Amazon catalog for the indicated &#x60;marketplaceId&#x60;. | 

## Methods

### NewItemIdentifiersByMarketplace

`func NewItemIdentifiersByMarketplace(marketplaceId string, identifiers []ItemIdentifier, ) *ItemIdentifiersByMarketplace`

NewItemIdentifiersByMarketplace instantiates a new ItemIdentifiersByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemIdentifiersByMarketplaceWithDefaults

`func NewItemIdentifiersByMarketplaceWithDefaults() *ItemIdentifiersByMarketplace`

NewItemIdentifiersByMarketplaceWithDefaults instantiates a new ItemIdentifiersByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemIdentifiersByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemIdentifiersByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemIdentifiersByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetIdentifiers

`func (o *ItemIdentifiersByMarketplace) GetIdentifiers() []ItemIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ItemIdentifiersByMarketplace) GetIdentifiersOk() (*[]ItemIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ItemIdentifiersByMarketplace) SetIdentifiers(v []ItemIdentifier)`

SetIdentifiers sets Identifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


