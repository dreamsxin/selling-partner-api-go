# InventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerSku** | **string** | The seller SKU of the item. | 
**MarketplaceId** | **string** | The marketplaceId. | 
**Quantity** | **int32** | The quantity of item to add. | 

## Methods

### NewInventoryItem

`func NewInventoryItem(sellerSku string, marketplaceId string, quantity int32, ) *InventoryItem`

NewInventoryItem instantiates a new InventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemWithDefaults

`func NewInventoryItemWithDefaults() *InventoryItem`

NewInventoryItemWithDefaults instantiates a new InventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellerSku

`func (o *InventoryItem) GetSellerSku() string`

GetSellerSku returns the SellerSku field if non-nil, zero value otherwise.

### GetSellerSkuOk

`func (o *InventoryItem) GetSellerSkuOk() (*string, bool)`

GetSellerSkuOk returns a tuple with the SellerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSku

`func (o *InventoryItem) SetSellerSku(v string)`

SetSellerSku sets SellerSku field to given value.


### GetMarketplaceId

`func (o *InventoryItem) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *InventoryItem) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *InventoryItem) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetQuantity

`func (o *InventoryItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InventoryItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InventoryItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


