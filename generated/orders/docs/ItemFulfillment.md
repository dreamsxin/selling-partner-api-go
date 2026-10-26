# ItemFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuantityFulfilled** | Pointer to **int32** | The number of units of this item that have been successfully processed and shipped. | [optional] 
**QuantityUnfulfilled** | Pointer to **int32** | The number of units of this item that remain to be processed and shipped. | [optional] 
**Picking** | Pointer to [**ItemPicking**](ItemPicking.md) |  | [optional] 
**Packing** | Pointer to [**ItemPacking**](ItemPacking.md) |  | [optional] 
**Shipping** | Pointer to [**ItemShipping**](ItemShipping.md) |  | [optional] 

## Methods

### NewItemFulfillment

`func NewItemFulfillment() *ItemFulfillment`

NewItemFulfillment instantiates a new ItemFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemFulfillmentWithDefaults

`func NewItemFulfillmentWithDefaults() *ItemFulfillment`

NewItemFulfillmentWithDefaults instantiates a new ItemFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantityFulfilled

`func (o *ItemFulfillment) GetQuantityFulfilled() int32`

GetQuantityFulfilled returns the QuantityFulfilled field if non-nil, zero value otherwise.

### GetQuantityFulfilledOk

`func (o *ItemFulfillment) GetQuantityFulfilledOk() (*int32, bool)`

GetQuantityFulfilledOk returns a tuple with the QuantityFulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFulfilled

`func (o *ItemFulfillment) SetQuantityFulfilled(v int32)`

SetQuantityFulfilled sets QuantityFulfilled field to given value.

### HasQuantityFulfilled

`func (o *ItemFulfillment) HasQuantityFulfilled() bool`

HasQuantityFulfilled returns a boolean if a field has been set.

### GetQuantityUnfulfilled

`func (o *ItemFulfillment) GetQuantityUnfulfilled() int32`

GetQuantityUnfulfilled returns the QuantityUnfulfilled field if non-nil, zero value otherwise.

### GetQuantityUnfulfilledOk

`func (o *ItemFulfillment) GetQuantityUnfulfilledOk() (*int32, bool)`

GetQuantityUnfulfilledOk returns a tuple with the QuantityUnfulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnfulfilled

`func (o *ItemFulfillment) SetQuantityUnfulfilled(v int32)`

SetQuantityUnfulfilled sets QuantityUnfulfilled field to given value.

### HasQuantityUnfulfilled

`func (o *ItemFulfillment) HasQuantityUnfulfilled() bool`

HasQuantityUnfulfilled returns a boolean if a field has been set.

### GetPicking

`func (o *ItemFulfillment) GetPicking() ItemPicking`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *ItemFulfillment) GetPickingOk() (*ItemPicking, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *ItemFulfillment) SetPicking(v ItemPicking)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *ItemFulfillment) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetPacking

`func (o *ItemFulfillment) GetPacking() ItemPacking`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *ItemFulfillment) GetPackingOk() (*ItemPacking, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *ItemFulfillment) SetPacking(v ItemPacking)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *ItemFulfillment) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetShipping

`func (o *ItemFulfillment) GetShipping() ItemShipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *ItemFulfillment) GetShippingOk() (*ItemShipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *ItemFulfillment) SetShipping(v ItemShipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *ItemFulfillment) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


