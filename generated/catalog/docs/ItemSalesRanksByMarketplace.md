# ItemSalesRanksByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**ClassificationRanks** | Pointer to [**[]ItemClassificationSalesRank**](ItemClassificationSalesRank.md) | Sales ranks of an Amazon catalog item for a &#x60;marketplaceId&#x60;, grouped by classification. | [optional] 
**DisplayGroupRanks** | Pointer to [**[]ItemDisplayGroupSalesRank**](ItemDisplayGroupSalesRank.md) | Sales ranks of an Amazon catalog item for a &#x60;marketplaceId&#x60;, grouped by website display group. | [optional] 

## Methods

### NewItemSalesRanksByMarketplace

`func NewItemSalesRanksByMarketplace(marketplaceId string, ) *ItemSalesRanksByMarketplace`

NewItemSalesRanksByMarketplace instantiates a new ItemSalesRanksByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSalesRanksByMarketplaceWithDefaults

`func NewItemSalesRanksByMarketplaceWithDefaults() *ItemSalesRanksByMarketplace`

NewItemSalesRanksByMarketplaceWithDefaults instantiates a new ItemSalesRanksByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemSalesRanksByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemSalesRanksByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemSalesRanksByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetClassificationRanks

`func (o *ItemSalesRanksByMarketplace) GetClassificationRanks() []ItemClassificationSalesRank`

GetClassificationRanks returns the ClassificationRanks field if non-nil, zero value otherwise.

### GetClassificationRanksOk

`func (o *ItemSalesRanksByMarketplace) GetClassificationRanksOk() (*[]ItemClassificationSalesRank, bool)`

GetClassificationRanksOk returns a tuple with the ClassificationRanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationRanks

`func (o *ItemSalesRanksByMarketplace) SetClassificationRanks(v []ItemClassificationSalesRank)`

SetClassificationRanks sets ClassificationRanks field to given value.

### HasClassificationRanks

`func (o *ItemSalesRanksByMarketplace) HasClassificationRanks() bool`

HasClassificationRanks returns a boolean if a field has been set.

### GetDisplayGroupRanks

`func (o *ItemSalesRanksByMarketplace) GetDisplayGroupRanks() []ItemDisplayGroupSalesRank`

GetDisplayGroupRanks returns the DisplayGroupRanks field if non-nil, zero value otherwise.

### GetDisplayGroupRanksOk

`func (o *ItemSalesRanksByMarketplace) GetDisplayGroupRanksOk() (*[]ItemDisplayGroupSalesRank, bool)`

GetDisplayGroupRanksOk returns a tuple with the DisplayGroupRanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayGroupRanks

`func (o *ItemSalesRanksByMarketplace) SetDisplayGroupRanks(v []ItemDisplayGroupSalesRank)`

SetDisplayGroupRanks sets DisplayGroupRanks field to given value.

### HasDisplayGroupRanks

`func (o *ItemSalesRanksByMarketplace) HasDisplayGroupRanks() bool`

HasDisplayGroupRanks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


