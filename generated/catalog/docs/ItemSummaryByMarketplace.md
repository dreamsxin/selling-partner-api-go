# ItemSummaryByMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | **string** | Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids). | 
**AdultProduct** | Pointer to **bool** | When &#x60;true&#x60;, the Amazon catalog item is intended for an adult audience or is sexual in nature. | [optional] 
**Autographed** | Pointer to **bool** | When &#x60;true&#x60;, the Amazon catalog item is autographed. | [optional] 
**Brand** | Pointer to **string** | Name of the brand that is associated with the Amazon catalog item. | [optional] 
**BrowseClassification** | Pointer to [**ItemBrowseClassification**](ItemBrowseClassification.md) |  | [optional] 
**Color** | Pointer to **string** | The color that is associated with the Amazon catalog item. | [optional] 
**Contributors** | Pointer to [**[]ItemContributor**](ItemContributor.md) | Individual contributors to the creation of the item, such as the authors or actors. | [optional] 
**ItemClassification** | Pointer to **string** | Classification type that is associated with the Amazon catalog item. | [optional] 
**ItemName** | Pointer to **string** | The name that is associated with the Amazon catalog item. | [optional] 
**Manufacturer** | Pointer to **string** | The name of the manufacturer that is associated with the Amazon catalog item. | [optional] 
**Memorabilia** | Pointer to **bool** | When true, the item is classified as memorabilia. | [optional] 
**ModelNumber** | Pointer to **string** | The model number that is associated with the Amazon catalog item. | [optional] 
**PackageQuantity** | Pointer to **int32** | The quantity of the Amazon catalog item within one package. | [optional] 
**PartNumber** | Pointer to **string** | The part number that is associated with the Amazon catalog item. | [optional] 
**ReleaseDate** | Pointer to **string** | The earliest date on which the Amazon catalog item can be shipped to customers. | [optional] 
**Size** | Pointer to **string** | The name of the size of the Amazon catalog item. | [optional] 
**Style** | Pointer to **string** | The name of the style that is associated with the Amazon catalog item. | [optional] 
**TradeInEligible** | Pointer to **bool** | When true, the Amazon catalog item is eligible for trade-in. | [optional] 
**WebsiteDisplayGroup** | Pointer to **string** | The identifier of the website display group that is associated with the Amazon catalog item. | [optional] 
**WebsiteDisplayGroupName** | Pointer to **string** | The display name of the website display group that is associated with the Amazon catalog item. | [optional] 

## Methods

### NewItemSummaryByMarketplace

`func NewItemSummaryByMarketplace(marketplaceId string, ) *ItemSummaryByMarketplace`

NewItemSummaryByMarketplace instantiates a new ItemSummaryByMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSummaryByMarketplaceWithDefaults

`func NewItemSummaryByMarketplaceWithDefaults() *ItemSummaryByMarketplace`

NewItemSummaryByMarketplaceWithDefaults instantiates a new ItemSummaryByMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *ItemSummaryByMarketplace) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ItemSummaryByMarketplace) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ItemSummaryByMarketplace) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetAdultProduct

`func (o *ItemSummaryByMarketplace) GetAdultProduct() bool`

GetAdultProduct returns the AdultProduct field if non-nil, zero value otherwise.

### GetAdultProductOk

`func (o *ItemSummaryByMarketplace) GetAdultProductOk() (*bool, bool)`

GetAdultProductOk returns a tuple with the AdultProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultProduct

`func (o *ItemSummaryByMarketplace) SetAdultProduct(v bool)`

SetAdultProduct sets AdultProduct field to given value.

### HasAdultProduct

`func (o *ItemSummaryByMarketplace) HasAdultProduct() bool`

HasAdultProduct returns a boolean if a field has been set.

### GetAutographed

`func (o *ItemSummaryByMarketplace) GetAutographed() bool`

GetAutographed returns the Autographed field if non-nil, zero value otherwise.

### GetAutographedOk

`func (o *ItemSummaryByMarketplace) GetAutographedOk() (*bool, bool)`

GetAutographedOk returns a tuple with the Autographed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutographed

`func (o *ItemSummaryByMarketplace) SetAutographed(v bool)`

SetAutographed sets Autographed field to given value.

### HasAutographed

`func (o *ItemSummaryByMarketplace) HasAutographed() bool`

HasAutographed returns a boolean if a field has been set.

### GetBrand

`func (o *ItemSummaryByMarketplace) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ItemSummaryByMarketplace) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ItemSummaryByMarketplace) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ItemSummaryByMarketplace) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBrowseClassification

`func (o *ItemSummaryByMarketplace) GetBrowseClassification() ItemBrowseClassification`

GetBrowseClassification returns the BrowseClassification field if non-nil, zero value otherwise.

### GetBrowseClassificationOk

`func (o *ItemSummaryByMarketplace) GetBrowseClassificationOk() (*ItemBrowseClassification, bool)`

GetBrowseClassificationOk returns a tuple with the BrowseClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowseClassification

`func (o *ItemSummaryByMarketplace) SetBrowseClassification(v ItemBrowseClassification)`

SetBrowseClassification sets BrowseClassification field to given value.

### HasBrowseClassification

`func (o *ItemSummaryByMarketplace) HasBrowseClassification() bool`

HasBrowseClassification returns a boolean if a field has been set.

### GetColor

`func (o *ItemSummaryByMarketplace) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ItemSummaryByMarketplace) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ItemSummaryByMarketplace) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ItemSummaryByMarketplace) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetContributors

`func (o *ItemSummaryByMarketplace) GetContributors() []ItemContributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ItemSummaryByMarketplace) GetContributorsOk() (*[]ItemContributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ItemSummaryByMarketplace) SetContributors(v []ItemContributor)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ItemSummaryByMarketplace) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetItemClassification

`func (o *ItemSummaryByMarketplace) GetItemClassification() string`

GetItemClassification returns the ItemClassification field if non-nil, zero value otherwise.

### GetItemClassificationOk

`func (o *ItemSummaryByMarketplace) GetItemClassificationOk() (*string, bool)`

GetItemClassificationOk returns a tuple with the ItemClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemClassification

`func (o *ItemSummaryByMarketplace) SetItemClassification(v string)`

SetItemClassification sets ItemClassification field to given value.

### HasItemClassification

`func (o *ItemSummaryByMarketplace) HasItemClassification() bool`

HasItemClassification returns a boolean if a field has been set.

### GetItemName

`func (o *ItemSummaryByMarketplace) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ItemSummaryByMarketplace) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ItemSummaryByMarketplace) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *ItemSummaryByMarketplace) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetManufacturer

`func (o *ItemSummaryByMarketplace) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ItemSummaryByMarketplace) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ItemSummaryByMarketplace) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ItemSummaryByMarketplace) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetMemorabilia

`func (o *ItemSummaryByMarketplace) GetMemorabilia() bool`

GetMemorabilia returns the Memorabilia field if non-nil, zero value otherwise.

### GetMemorabiliaOk

`func (o *ItemSummaryByMarketplace) GetMemorabiliaOk() (*bool, bool)`

GetMemorabiliaOk returns a tuple with the Memorabilia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorabilia

`func (o *ItemSummaryByMarketplace) SetMemorabilia(v bool)`

SetMemorabilia sets Memorabilia field to given value.

### HasMemorabilia

`func (o *ItemSummaryByMarketplace) HasMemorabilia() bool`

HasMemorabilia returns a boolean if a field has been set.

### GetModelNumber

`func (o *ItemSummaryByMarketplace) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *ItemSummaryByMarketplace) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *ItemSummaryByMarketplace) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *ItemSummaryByMarketplace) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetPackageQuantity

`func (o *ItemSummaryByMarketplace) GetPackageQuantity() int32`

GetPackageQuantity returns the PackageQuantity field if non-nil, zero value otherwise.

### GetPackageQuantityOk

`func (o *ItemSummaryByMarketplace) GetPackageQuantityOk() (*int32, bool)`

GetPackageQuantityOk returns a tuple with the PackageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuantity

`func (o *ItemSummaryByMarketplace) SetPackageQuantity(v int32)`

SetPackageQuantity sets PackageQuantity field to given value.

### HasPackageQuantity

`func (o *ItemSummaryByMarketplace) HasPackageQuantity() bool`

HasPackageQuantity returns a boolean if a field has been set.

### GetPartNumber

`func (o *ItemSummaryByMarketplace) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ItemSummaryByMarketplace) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ItemSummaryByMarketplace) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ItemSummaryByMarketplace) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ItemSummaryByMarketplace) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ItemSummaryByMarketplace) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ItemSummaryByMarketplace) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ItemSummaryByMarketplace) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSize

`func (o *ItemSummaryByMarketplace) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ItemSummaryByMarketplace) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ItemSummaryByMarketplace) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ItemSummaryByMarketplace) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStyle

`func (o *ItemSummaryByMarketplace) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ItemSummaryByMarketplace) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ItemSummaryByMarketplace) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ItemSummaryByMarketplace) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetTradeInEligible

`func (o *ItemSummaryByMarketplace) GetTradeInEligible() bool`

GetTradeInEligible returns the TradeInEligible field if non-nil, zero value otherwise.

### GetTradeInEligibleOk

`func (o *ItemSummaryByMarketplace) GetTradeInEligibleOk() (*bool, bool)`

GetTradeInEligibleOk returns a tuple with the TradeInEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeInEligible

`func (o *ItemSummaryByMarketplace) SetTradeInEligible(v bool)`

SetTradeInEligible sets TradeInEligible field to given value.

### HasTradeInEligible

`func (o *ItemSummaryByMarketplace) HasTradeInEligible() bool`

HasTradeInEligible returns a boolean if a field has been set.

### GetWebsiteDisplayGroup

`func (o *ItemSummaryByMarketplace) GetWebsiteDisplayGroup() string`

GetWebsiteDisplayGroup returns the WebsiteDisplayGroup field if non-nil, zero value otherwise.

### GetWebsiteDisplayGroupOk

`func (o *ItemSummaryByMarketplace) GetWebsiteDisplayGroupOk() (*string, bool)`

GetWebsiteDisplayGroupOk returns a tuple with the WebsiteDisplayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteDisplayGroup

`func (o *ItemSummaryByMarketplace) SetWebsiteDisplayGroup(v string)`

SetWebsiteDisplayGroup sets WebsiteDisplayGroup field to given value.

### HasWebsiteDisplayGroup

`func (o *ItemSummaryByMarketplace) HasWebsiteDisplayGroup() bool`

HasWebsiteDisplayGroup returns a boolean if a field has been set.

### GetWebsiteDisplayGroupName

`func (o *ItemSummaryByMarketplace) GetWebsiteDisplayGroupName() string`

GetWebsiteDisplayGroupName returns the WebsiteDisplayGroupName field if non-nil, zero value otherwise.

### GetWebsiteDisplayGroupNameOk

`func (o *ItemSummaryByMarketplace) GetWebsiteDisplayGroupNameOk() (*string, bool)`

GetWebsiteDisplayGroupNameOk returns a tuple with the WebsiteDisplayGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteDisplayGroupName

`func (o *ItemSummaryByMarketplace) SetWebsiteDisplayGroupName(v string)`

SetWebsiteDisplayGroupName sets WebsiteDisplayGroupName field to given value.

### HasWebsiteDisplayGroupName

`func (o *ItemSummaryByMarketplace) HasWebsiteDisplayGroupName() bool`

HasWebsiteDisplayGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


