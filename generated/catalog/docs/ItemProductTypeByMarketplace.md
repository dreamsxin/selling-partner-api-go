# ItemProductTypeByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | Pointer to **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | [optional] 
**ProductType** | Pointer to **string** | Name of the product type that is associated with the Amazon catalog item. | [optional] 

## Methods

### NewItemProductTypeByMarketplace

`func NewItemProductTypeByMarketplace() *ItemProductTypeByMarketplace`

NewItemProductTypeByMarketplace instantiates a new ItemProductTypeByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemProductTypeByMarketplaceWithDefaults

`func NewItemProductTypeByMarketplaceWithDefaults() *ItemProductTypeByMarketplace`

NewItemProductTypeByMarketplaceWithDefaults instantiates a new ItemProductTypeByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemProductTypeByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemProductTypeByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemProductTypeByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *ItemProductTypeByMarketplace) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetProductType

`func (o *ItemProductTypeByMarketplace) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ItemProductTypeByMarketplace) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ItemProductTypeByMarketplace) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ItemProductTypeByMarketplace) HasProductType() bool`

HasProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


