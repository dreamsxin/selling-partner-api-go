# ItemProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | Pointer to **string** | The Amazon Standard Identification Number (ASIN), which uniquely identifies a product (catalog item). | [optional] 
**Title** | Pointer to **string** | The product name or title as displayed to customers in the marketplace. | [optional] 
**SellerSku** | Pointer to **string** | The seller SKU of a product (catalog item). This is a unique number assigned by the seller when listing an item. | [optional] 
**Condition** | Pointer to [**ItemCondition**](ItemCondition.md) |  | [optional] 
**Price** | Pointer to [**ItemPrice**](ItemPrice.md) |  | [optional] 
**SerialNumbers** | Pointer to **[]string** | Unique serial numbers for products that require individual tracking, typically provided for FBA orders. | [optional] 
**Customization** | Pointer to [**ItemCustomization**](ItemCustomization.md) |  | [optional] 

## Methods

### NewItemProduct

`func NewItemProduct() *ItemProduct`

NewItemProduct instantiates a new ItemProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemProductWithDefaults

`func NewItemProductWithDefaults() *ItemProduct`

NewItemProductWithDefaults instantiates a new ItemProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsin

`func (o *ItemProduct) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ItemProduct) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ItemProduct) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ItemProduct) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetTitle

`func (o *ItemProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSellerSku

`func (o *ItemProduct) GetSellerSku() string`

GetSellerSku returns the SellerSku field if non-nil, zero value otherwise.

### GetSellerSkuOk

`func (o *ItemProduct) GetSellerSkuOk() (*string, bool)`

GetSellerSkuOk returns a tuple with the SellerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSku

`func (o *ItemProduct) SetSellerSku(v string)`

SetSellerSku sets SellerSku field to given value.

### HasSellerSku

`func (o *ItemProduct) HasSellerSku() bool`

HasSellerSku returns a boolean if a field has been set.

### GetCondition

`func (o *ItemProduct) GetCondition() ItemCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ItemProduct) GetConditionOk() (*ItemCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ItemProduct) SetCondition(v ItemCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ItemProduct) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetPrice

`func (o *ItemProduct) GetPrice() ItemPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemProduct) GetPriceOk() (*ItemPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemProduct) SetPrice(v ItemPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ItemProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *ItemProduct) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *ItemProduct) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *ItemProduct) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *ItemProduct) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetCustomization

`func (o *ItemProduct) GetCustomization() ItemCustomization`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *ItemProduct) GetCustomizationOk() (*ItemCustomization, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *ItemProduct) SetCustomization(v ItemCustomization)`

SetCustomization sets Customization field to given value.

### HasCustomization

`func (o *ItemProduct) HasCustomization() bool`

HasCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


