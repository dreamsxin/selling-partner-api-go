# ItemBrowseClassificationsByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**Classifications** | Pointer to [**[]ItemBrowseClassification**](ItemBrowseClassification.md) | Classifications (browse nodes) that are associated with the item in the Amazon catalog. | [optional] 

## Methods

### NewItemBrowseClassificationsByMarketplace

`func NewItemBrowseClassificationsByMarketplace(marketplaceId string, ) *ItemBrowseClassificationsByMarketplace`

NewItemBrowseClassificationsByMarketplace instantiates a new ItemBrowseClassificationsByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemBrowseClassificationsByMarketplaceWithDefaults

`func NewItemBrowseClassificationsByMarketplaceWithDefaults() *ItemBrowseClassificationsByMarketplace`

NewItemBrowseClassificationsByMarketplaceWithDefaults instantiates a new ItemBrowseClassificationsByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemBrowseClassificationsByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemBrowseClassificationsByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemBrowseClassificationsByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetClassifications

`func (o *ItemBrowseClassificationsByMarketplace) GetClassifications() []ItemBrowseClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *ItemBrowseClassificationsByMarketplace) GetClassificationsOk() (*[]ItemBrowseClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *ItemBrowseClassificationsByMarketplace) SetClassifications(v []ItemBrowseClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *ItemBrowseClassificationsByMarketplace) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


