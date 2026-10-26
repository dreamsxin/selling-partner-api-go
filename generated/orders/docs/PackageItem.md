# PackageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** | Unique identifier of the order item included in this package. | 
**Quantity** | **int32** | Number of units of this item included in the package shipment. | 
**TransparencyCodes** | Pointer to **[]string** | The transparency codes associated with this item for product authentication. | [optional] 

## Methods

### NewPackageItem

`func NewPackageItem(orderItemId string, quantity int32, ) *PackageItem`

NewPackageItem instantiates a new PackageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageItemWithDefaults

`func NewPackageItemWithDefaults() *PackageItem`

NewPackageItemWithDefaults instantiates a new PackageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *PackageItem) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *PackageItem) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *PackageItem) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetQuantity

`func (o *PackageItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PackageItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PackageItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTransparencyCodes

`func (o *PackageItem) GetTransparencyCodes() []string`

GetTransparencyCodes returns the TransparencyCodes field if non-nil, zero value otherwise.

### GetTransparencyCodesOk

`func (o *PackageItem) GetTransparencyCodesOk() (*[]string, bool)`

GetTransparencyCodesOk returns a tuple with the TransparencyCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparencyCodes

`func (o *PackageItem) SetTransparencyCodes(v []string)`

SetTransparencyCodes sets TransparencyCodes field to given value.

### HasTransparencyCodes

`func (o *PackageItem) HasTransparencyCodes() bool`

HasTransparencyCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


