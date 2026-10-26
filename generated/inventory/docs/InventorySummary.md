# InventorySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | Pointer to **string** | The Amazon Standard Identification Number (ASIN) of an item. | [optional] 
**FnSku** | Pointer to **string** | Amazon&#39;s fulfillment network SKU identifier. | [optional] 
**SellerSku** | Pointer to **string** | The seller SKU of the item. | [optional] 
**Condition** | Pointer to **string** | The condition of the item as described by the seller (for example, New Item). | [optional] 
**InventoryDetails** | Pointer to [**InventoryDetails**](InventoryDetails.md) |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time that any quantity was last updated. | [optional] 
**ProductName** | Pointer to **string** | The localized language product title of the item within the specific marketplace. | [optional] 
**TotalQuantity** | Pointer to **int32** | The total number of units in an inbound shipment or in Amazon fulfillment centers. | [optional] 
**Stores** | Pointer to **[]string** | A list of seller-enrolled stores that apply to this seller SKU. | [optional] 

## Methods

### NewInventorySummary

`func NewInventorySummary() *InventorySummary`

NewInventorySummary instantiates a new InventorySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventorySummaryWithDefaults

`func NewInventorySummaryWithDefaults() *InventorySummary`

NewInventorySummaryWithDefaults instantiates a new InventorySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsin

`func (o *InventorySummary) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *InventorySummary) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *InventorySummary) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *InventorySummary) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetFnSku

`func (o *InventorySummary) GetFnSku() string`

GetFnSku returns the FnSku field if non-nil, zero value otherwise.

### GetFnSkuOk

`func (o *InventorySummary) GetFnSkuOk() (*string, bool)`

GetFnSkuOk returns a tuple with the FnSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnSku

`func (o *InventorySummary) SetFnSku(v string)`

SetFnSku sets FnSku field to given value.

### HasFnSku

`func (o *InventorySummary) HasFnSku() bool`

HasFnSku returns a boolean if a field has been set.

### GetSellerSku

`func (o *InventorySummary) GetSellerSku() string`

GetSellerSku returns the SellerSku field if non-nil, zero value otherwise.

### GetSellerSkuOk

`func (o *InventorySummary) GetSellerSkuOk() (*string, bool)`

GetSellerSkuOk returns a tuple with the SellerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSku

`func (o *InventorySummary) SetSellerSku(v string)`

SetSellerSku sets SellerSku field to given value.

### HasSellerSku

`func (o *InventorySummary) HasSellerSku() bool`

HasSellerSku returns a boolean if a field has been set.

### GetCondition

`func (o *InventorySummary) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *InventorySummary) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *InventorySummary) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *InventorySummary) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetInventoryDetails

`func (o *InventorySummary) GetInventoryDetails() InventoryDetails`

GetInventoryDetails returns the InventoryDetails field if non-nil, zero value otherwise.

### GetInventoryDetailsOk

`func (o *InventorySummary) GetInventoryDetailsOk() (*InventoryDetails, bool)`

GetInventoryDetailsOk returns a tuple with the InventoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDetails

`func (o *InventorySummary) SetInventoryDetails(v InventoryDetails)`

SetInventoryDetails sets InventoryDetails field to given value.

### HasInventoryDetails

`func (o *InventorySummary) HasInventoryDetails() bool`

HasInventoryDetails returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *InventorySummary) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *InventorySummary) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *InventorySummary) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *InventorySummary) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetProductName

`func (o *InventorySummary) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *InventorySummary) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *InventorySummary) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *InventorySummary) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *InventorySummary) GetTotalQuantity() int32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *InventorySummary) GetTotalQuantityOk() (*int32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *InventorySummary) SetTotalQuantity(v int32)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *InventorySummary) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetStores

`func (o *InventorySummary) GetStores() []string`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *InventorySummary) GetStoresOk() (*[]string, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *InventorySummary) SetStores(v []string)`

SetStores sets Stores field to given value.

### HasStores

`func (o *InventorySummary) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


