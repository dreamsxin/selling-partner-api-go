# ItemRelationshipsByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**Relationships** | [**[]ItemRelationship**](ItemRelationship.md) | Relationships for the item. | 

## Methods

### NewItemRelationshipsByMarketplace

`func NewItemRelationshipsByMarketplace(marketplaceId string, relationships []ItemRelationship, ) *ItemRelationshipsByMarketplace`

NewItemRelationshipsByMarketplace instantiates a new ItemRelationshipsByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRelationshipsByMarketplaceWithDefaults

`func NewItemRelationshipsByMarketplaceWithDefaults() *ItemRelationshipsByMarketplace`

NewItemRelationshipsByMarketplaceWithDefaults instantiates a new ItemRelationshipsByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemRelationshipsByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemRelationshipsByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemRelationshipsByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetRelationships

`func (o *ItemRelationshipsByMarketplace) GetRelationships() []ItemRelationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ItemRelationshipsByMarketplace) GetRelationshipsOk() (*[]ItemRelationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ItemRelationshipsByMarketplace) SetRelationships(v []ItemRelationship)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


