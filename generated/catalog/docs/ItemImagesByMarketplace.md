# ItemImagesByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**Images** | [**[]ItemImage**](ItemImage.md) | Images for an item in the Amazon catalog, grouped by &#x60;marketplaceId&#x60;. | 

## Methods

### NewItemImagesByMarketplace

`func NewItemImagesByMarketplace(marketplaceId string, images []ItemImage, ) *ItemImagesByMarketplace`

NewItemImagesByMarketplace instantiates a new ItemImagesByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImagesByMarketplaceWithDefaults

`func NewItemImagesByMarketplaceWithDefaults() *ItemImagesByMarketplace`

NewItemImagesByMarketplaceWithDefaults instantiates a new ItemImagesByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemImagesByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemImagesByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemImagesByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetImages

`func (o *ItemImagesByMarketplace) GetImages() []ItemImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ItemImagesByMarketplace) GetImagesOk() (*[]ItemImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ItemImagesByMarketplace) SetImages(v []ItemImage)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


