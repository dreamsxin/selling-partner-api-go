# ItemDimensionsByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**Item** | Pointer to [**Dimensions**](Dimensions.md) |  | [optional] 
**Package** | Pointer to [**Dimensions**](Dimensions.md) |  | [optional] 

## Methods

### NewItemDimensionsByMarketplace

`func NewItemDimensionsByMarketplace(marketplaceId string, ) *ItemDimensionsByMarketplace`

NewItemDimensionsByMarketplace instantiates a new ItemDimensionsByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemDimensionsByMarketplaceWithDefaults

`func NewItemDimensionsByMarketplaceWithDefaults() *ItemDimensionsByMarketplace`

NewItemDimensionsByMarketplaceWithDefaults instantiates a new ItemDimensionsByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemDimensionsByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemDimensionsByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemDimensionsByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetItem

`func (o *ItemDimensionsByMarketplace) GetItem() Dimensions`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemDimensionsByMarketplace) GetItemOk() (*Dimensions, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemDimensionsByMarketplace) SetItem(v Dimensions)`

SetItem sets Item field to given value.

### HasItem

`func (o *ItemDimensionsByMarketplace) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetPackage

`func (o *ItemDimensionsByMarketplace) GetPackage() Dimensions`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ItemDimensionsByMarketplace) GetPackageOk() (*Dimensions, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ItemDimensionsByMarketplace) SetPackage(v Dimensions)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ItemDimensionsByMarketplace) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


