# CreateInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerSku** | **string** | The seller SKU of the item. | 
**MarketplaceId** | **string** | The marketplaceId. | 
**ProductName** | **string** | The name of the item. | 

## Methods

### NewCreateInventoryItemRequest

`func NewCreateInventoryItemRequest(sellerSku string, marketplaceId string, productName string, ) *CreateInventoryItemRequest`

NewCreateInventoryItemRequest instantiates a new CreateInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryItemRequestWithDefaults

`func NewCreateInventoryItemRequestWithDefaults() *CreateInventoryItemRequest`

NewCreateInventoryItemRequestWithDefaults instantiates a new CreateInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellerSku

`func (o *CreateInventoryItemRequest) GetSellerSku() string`

GetSellerSku returns the SellerSku field if non-nil, zero value otherwise.

### GetSellerSkuOk

`func (o *CreateInventoryItemRequest) GetSellerSkuOk() (*string, bool)`

GetSellerSkuOk returns a tuple with the SellerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSku

`func (o *CreateInventoryItemRequest) SetSellerSku(v string)`

SetSellerSku sets SellerSku field to given value.


### GetMarketplaceId

`func (o *CreateInventoryItemRequest) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *CreateInventoryItemRequest) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *CreateInventoryItemRequest) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.


### GetProductName

`func (o *CreateInventoryItemRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CreateInventoryItemRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CreateInventoryItemRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


