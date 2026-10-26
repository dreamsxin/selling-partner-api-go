# ItemVendorDetailsByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**BrandCode** | Pointer to **string** | The brand code that is associated with an Amazon catalog item. | [optional] 
**ManufacturerCode** | Pointer to **string** | The manufacturer code that is associated with an Amazon catalog item. | [optional] 
**ManufacturerCodeParent** | Pointer to **string** | The parent vendor code of the manufacturer code. | [optional] 
**ProductCategory** | Pointer to [**ItemVendorDetailsCategory**](ItemVendorDetailsCategory.md) |  | [optional] 
**ProductGroup** | Pointer to **string** | The product group that is associated with an Amazon catalog item. | [optional] 
**ProductSubcategory** | Pointer to [**ItemVendorDetailsCategory**](ItemVendorDetailsCategory.md) |  | [optional] 
**ReplenishmentCategory** | Pointer to **string** | The replenishment category that is associated with an Amazon catalog item. | [optional] 

## Methods

### NewItemVendorDetailsByMarketplace

`func NewItemVendorDetailsByMarketplace(marketplaceId string, ) *ItemVendorDetailsByMarketplace`

NewItemVendorDetailsByMarketplace instantiates a new ItemVendorDetailsByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemVendorDetailsByMarketplaceWithDefaults

`func NewItemVendorDetailsByMarketplaceWithDefaults() *ItemVendorDetailsByMarketplace`

NewItemVendorDetailsByMarketplaceWithDefaults instantiates a new ItemVendorDetailsByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemVendorDetailsByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemVendorDetailsByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemVendorDetailsByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetBrandCode

`func (o *ItemVendorDetailsByMarketplace) GetBrandCode() string`

GetBrandCode returns the BrandCode field if non-nil, zero value otherwise.

### GetBrandCodeOk

`func (o *ItemVendorDetailsByMarketplace) GetBrandCodeOk() (*string, bool)`

GetBrandCodeOk returns a tuple with the BrandCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCode

`func (o *ItemVendorDetailsByMarketplace) SetBrandCode(v string)`

SetBrandCode sets BrandCode field to given value.

### HasBrandCode

`func (o *ItemVendorDetailsByMarketplace) HasBrandCode() bool`

HasBrandCode returns a boolean if a field has been set.

### GetManufacturerCode

`func (o *ItemVendorDetailsByMarketplace) GetManufacturerCode() string`

GetManufacturerCode returns the ManufacturerCode field if non-nil, zero value otherwise.

### GetManufacturerCodeOk

`func (o *ItemVendorDetailsByMarketplace) GetManufacturerCodeOk() (*string, bool)`

GetManufacturerCodeOk returns a tuple with the ManufacturerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerCode

`func (o *ItemVendorDetailsByMarketplace) SetManufacturerCode(v string)`

SetManufacturerCode sets ManufacturerCode field to given value.

### HasManufacturerCode

`func (o *ItemVendorDetailsByMarketplace) HasManufacturerCode() bool`

HasManufacturerCode returns a boolean if a field has been set.

### GetManufacturerCodeParent

`func (o *ItemVendorDetailsByMarketplace) GetManufacturerCodeParent() string`

GetManufacturerCodeParent returns the ManufacturerCodeParent field if non-nil, zero value otherwise.

### GetManufacturerCodeParentOk

`func (o *ItemVendorDetailsByMarketplace) GetManufacturerCodeParentOk() (*string, bool)`

GetManufacturerCodeParentOk returns a tuple with the ManufacturerCodeParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerCodeParent

`func (o *ItemVendorDetailsByMarketplace) SetManufacturerCodeParent(v string)`

SetManufacturerCodeParent sets ManufacturerCodeParent field to given value.

### HasManufacturerCodeParent

`func (o *ItemVendorDetailsByMarketplace) HasManufacturerCodeParent() bool`

HasManufacturerCodeParent returns a boolean if a field has been set.

### GetProductCategory

`func (o *ItemVendorDetailsByMarketplace) GetProductCategory() ItemVendorDetailsCategory`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ItemVendorDetailsByMarketplace) GetProductCategoryOk() (*ItemVendorDetailsCategory, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ItemVendorDetailsByMarketplace) SetProductCategory(v ItemVendorDetailsCategory)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ItemVendorDetailsByMarketplace) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### GetProductGroup

`func (o *ItemVendorDetailsByMarketplace) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *ItemVendorDetailsByMarketplace) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *ItemVendorDetailsByMarketplace) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *ItemVendorDetailsByMarketplace) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### GetProductSubcategory

`func (o *ItemVendorDetailsByMarketplace) GetProductSubcategory() ItemVendorDetailsCategory`

GetProductSubcategory returns the ProductSubcategory field if non-nil, zero value otherwise.

### GetProductSubcategoryOk

`func (o *ItemVendorDetailsByMarketplace) GetProductSubcategoryOk() (*ItemVendorDetailsCategory, bool)`

GetProductSubcategoryOk returns a tuple with the ProductSubcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubcategory

`func (o *ItemVendorDetailsByMarketplace) SetProductSubcategory(v ItemVendorDetailsCategory)`

SetProductSubcategory sets ProductSubcategory field to given value.

### HasProductSubcategory

`func (o *ItemVendorDetailsByMarketplace) HasProductSubcategory() bool`

HasProductSubcategory returns a boolean if a field has been set.

### GetReplenishmentCategory

`func (o *ItemVendorDetailsByMarketplace) GetReplenishmentCategory() string`

GetReplenishmentCategory returns the ReplenishmentCategory field if non-nil, zero value otherwise.

### GetReplenishmentCategoryOk

`func (o *ItemVendorDetailsByMarketplace) GetReplenishmentCategoryOk() (*string, bool)`

GetReplenishmentCategoryOk returns a tuple with the ReplenishmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplenishmentCategory

`func (o *ItemVendorDetailsByMarketplace) SetReplenishmentCategory(v string)`

SetReplenishmentCategory sets ReplenishmentCategory field to given value.

### HasReplenishmentCategory

`func (o *ItemVendorDetailsByMarketplace) HasReplenishmentCategory() bool`

HasReplenishmentCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


