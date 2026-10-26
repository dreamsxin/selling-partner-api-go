# ProductContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | Pointer to **string** | The Amazon Standard Identification Number (ASIN) of the item. | [optional] 
**Sku** | Pointer to **string** | The Stock Keeping Unit (SKU) of the item. | [optional] 
**QuantityShipped** | Pointer to **int32** | The quantity of the item shipped. | [optional] 
**FulfillmentNetwork** | Pointer to **string** | The fulfillment network of the item. | [optional] 

## Methods

### NewProductContext

`func NewProductContext() *ProductContext`

NewProductContext instantiates a new ProductContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductContextWithDefaults

`func NewProductContextWithDefaults() *ProductContext`

NewProductContextWithDefaults instantiates a new ProductContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsin

`func (o *ProductContext) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ProductContext) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ProductContext) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ProductContext) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetSku

`func (o *ProductContext) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductContext) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductContext) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductContext) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetQuantityShipped

`func (o *ProductContext) GetQuantityShipped() int32`

GetQuantityShipped returns the QuantityShipped field if non-nil, zero value otherwise.

### GetQuantityShippedOk

`func (o *ProductContext) GetQuantityShippedOk() (*int32, bool)`

GetQuantityShippedOk returns a tuple with the QuantityShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityShipped

`func (o *ProductContext) SetQuantityShipped(v int32)`

SetQuantityShipped sets QuantityShipped field to given value.

### HasQuantityShipped

`func (o *ProductContext) HasQuantityShipped() bool`

HasQuantityShipped returns a boolean if a field has been set.

### GetFulfillmentNetwork

`func (o *ProductContext) GetFulfillmentNetwork() string`

GetFulfillmentNetwork returns the FulfillmentNetwork field if non-nil, zero value otherwise.

### GetFulfillmentNetworkOk

`func (o *ProductContext) GetFulfillmentNetworkOk() (*string, bool)`

GetFulfillmentNetworkOk returns a tuple with the FulfillmentNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentNetwork

`func (o *ProductContext) SetFulfillmentNetwork(v string)`

SetFulfillmentNetwork sets FulfillmentNetwork field to given value.

### HasFulfillmentNetwork

`func (o *ProductContext) HasFulfillmentNetwork() bool`

HasFulfillmentNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


