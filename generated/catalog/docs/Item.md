# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | **string** | The unique identifier of an item in the Amazon catalog. | 
**Attributes** | Pointer to **map[string]interface{}** | A JSON object containing structured item attribute data that is keyed by attribute name. Catalog item attributes conform to the related Amazon product type definitions that you can get from the [Product Type Definitions API](https://developer-docs.amazon.com/sp-api/reference/product-type-definitions-v2020-09-01). | [optional] 
**Classifications** | Pointer to [**[]ItemBrowseClassificationsByMarketplace**](ItemBrowseClassificationsByMarketplace.md) | An array of classifications (browse nodes) that is associated with the item in the Amazon catalog, grouped by &#x60;marketplaceId&#x60;. | [optional] 
**Dimensions** | Pointer to [**[]ItemDimensionsByMarketplace**](ItemDimensionsByMarketplace.md) | An array of dimensions that are associated with the item in the Amazon catalog, grouped by &#x60;marketplaceId&#x60;. | [optional] 
**Identifiers** | Pointer to [**[]ItemIdentifiersByMarketplace**](ItemIdentifiersByMarketplace.md) | Identifiers associated with the item in the Amazon catalog, such as UPC and EAN identifiers. | [optional] 
**Images** | Pointer to [**[]ItemImagesByMarketplace**](ItemImagesByMarketplace.md) | The images for an item in the Amazon catalog. | [optional] 
**ProductTypes** | Pointer to [**[]ItemProductTypeByMarketplace**](ItemProductTypeByMarketplace.md) | Product types that are associated with the Amazon catalog item. | [optional] 
**Relationships** | Pointer to [**[]ItemRelationshipsByMarketplace**](ItemRelationshipsByMarketplace.md) | Relationships grouped by &#x60;marketplaceId&#x60; for an Amazon catalog item (for example, variations). | [optional] 
**SalesRanks** | Pointer to [**[]ItemSalesRanksByMarketplace**](ItemSalesRanksByMarketplace.md) | Sales ranks of an Amazon catalog item. | [optional] 
**Summaries** | Pointer to [**[]ItemSummaryByMarketplace**](ItemSummaryByMarketplace.md) | Summaries of Amazon catalog items. | [optional] 
**VendorDetails** | Pointer to [**[]ItemVendorDetailsByMarketplace**](ItemVendorDetailsByMarketplace.md) | The vendor details that are associated with an Amazon catalog item. Vendor details are only available to vendors. | [optional] 

## Methods

### NewItem

`func NewItem(asin string, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsin

`func (o *Item) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *Item) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *Item) SetAsin(v string)`

SetAsin sets Asin field to given value.


### GetAttributes

`func (o *Item) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Item) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Item) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Item) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetClassifications

`func (o *Item) GetClassifications() []ItemBrowseClassificationsByMarketplace`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *Item) GetClassificationsOk() (*[]ItemBrowseClassificationsByMarketplace, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *Item) SetClassifications(v []ItemBrowseClassificationsByMarketplace)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *Item) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetDimensions

`func (o *Item) GetDimensions() []ItemDimensionsByMarketplace`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Item) GetDimensionsOk() (*[]ItemDimensionsByMarketplace, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Item) SetDimensions(v []ItemDimensionsByMarketplace)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *Item) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetIdentifiers

`func (o *Item) GetIdentifiers() []ItemIdentifiersByMarketplace`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *Item) GetIdentifiersOk() (*[]ItemIdentifiersByMarketplace, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *Item) SetIdentifiers(v []ItemIdentifiersByMarketplace)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *Item) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetImages

`func (o *Item) GetImages() []ItemImagesByMarketplace`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Item) GetImagesOk() (*[]ItemImagesByMarketplace, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Item) SetImages(v []ItemImagesByMarketplace)`

SetImages sets Images field to given value.

### HasImages

`func (o *Item) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProductTypes

`func (o *Item) GetProductTypes() []ItemProductTypeByMarketplace`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *Item) GetProductTypesOk() (*[]ItemProductTypeByMarketplace, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *Item) SetProductTypes(v []ItemProductTypeByMarketplace)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *Item) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetRelationships

`func (o *Item) GetRelationships() []ItemRelationshipsByMarketplace`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Item) GetRelationshipsOk() (*[]ItemRelationshipsByMarketplace, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Item) SetRelationships(v []ItemRelationshipsByMarketplace)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Item) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSalesRanks

`func (o *Item) GetSalesRanks() []ItemSalesRanksByMarketplace`

GetSalesRanks returns the SalesRanks field if non-nil, zero value otherwise.

### GetSalesRanksOk

`func (o *Item) GetSalesRanksOk() (*[]ItemSalesRanksByMarketplace, bool)`

GetSalesRanksOk returns a tuple with the SalesRanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRanks

`func (o *Item) SetSalesRanks(v []ItemSalesRanksByMarketplace)`

SetSalesRanks sets SalesRanks field to given value.

### HasSalesRanks

`func (o *Item) HasSalesRanks() bool`

HasSalesRanks returns a boolean if a field has been set.

### GetSummaries

`func (o *Item) GetSummaries() []ItemSummaryByMarketplace`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *Item) GetSummariesOk() (*[]ItemSummaryByMarketplace, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *Item) SetSummaries(v []ItemSummaryByMarketplace)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *Item) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### GetVendorDetails

`func (o *Item) GetVendorDetails() []ItemVendorDetailsByMarketplace`

GetVendorDetails returns the VendorDetails field if non-nil, zero value otherwise.

### GetVendorDetailsOk

`func (o *Item) GetVendorDetailsOk() (*[]ItemVendorDetailsByMarketplace, bool)`

GetVendorDetailsOk returns a tuple with the VendorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorDetails

`func (o *Item) SetVendorDetails(v []ItemVendorDetailsByMarketplace)`

SetVendorDetails sets VendorDetails field to given value.

### HasVendorDetails

`func (o *Item) HasVendorDetails() bool`

HasVendorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


