# UnfulfillableQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUnfulfillableQuantity** | Pointer to **int32** | The total number of units in Amazon&#39;s fulfillment network in unsellable condition. | [optional] 
**CustomerDamagedQuantity** | Pointer to **int32** | The number of units in customer damaged disposition. | [optional] 
**WarehouseDamagedQuantity** | Pointer to **int32** | The number of units in warehouse damaged disposition. | [optional] 
**DistributorDamagedQuantity** | Pointer to **int32** | The number of units in distributor damaged disposition. | [optional] 
**CarrierDamagedQuantity** | Pointer to **int32** | The number of units in carrier damaged disposition. | [optional] 
**DefectiveQuantity** | Pointer to **int32** | The number of units in defective disposition. | [optional] 
**ExpiredQuantity** | Pointer to **int32** | The number of units in expired disposition. | [optional] 

## Methods

### NewUnfulfillableQuantity

`func NewUnfulfillableQuantity() *UnfulfillableQuantity`

NewUnfulfillableQuantity instantiates a new UnfulfillableQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnfulfillableQuantityWithDefaults

`func NewUnfulfillableQuantityWithDefaults() *UnfulfillableQuantity`

NewUnfulfillableQuantityWithDefaults instantiates a new UnfulfillableQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUnfulfillableQuantity

`func (o *UnfulfillableQuantity) GetTotalUnfulfillableQuantity() int32`

GetTotalUnfulfillableQuantity returns the TotalUnfulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalUnfulfillableQuantityOk

`func (o *UnfulfillableQuantity) GetTotalUnfulfillableQuantityOk() (*int32, bool)`

GetTotalUnfulfillableQuantityOk returns a tuple with the TotalUnfulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnfulfillableQuantity

`func (o *UnfulfillableQuantity) SetTotalUnfulfillableQuantity(v int32)`

SetTotalUnfulfillableQuantity sets TotalUnfulfillableQuantity field to given value.

### HasTotalUnfulfillableQuantity

`func (o *UnfulfillableQuantity) HasTotalUnfulfillableQuantity() bool`

HasTotalUnfulfillableQuantity returns a boolean if a field has been set.

### GetCustomerDamagedQuantity

`func (o *UnfulfillableQuantity) GetCustomerDamagedQuantity() int32`

GetCustomerDamagedQuantity returns the CustomerDamagedQuantity field if non-nil, zero value otherwise.

### GetCustomerDamagedQuantityOk

`func (o *UnfulfillableQuantity) GetCustomerDamagedQuantityOk() (*int32, bool)`

GetCustomerDamagedQuantityOk returns a tuple with the CustomerDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDamagedQuantity

`func (o *UnfulfillableQuantity) SetCustomerDamagedQuantity(v int32)`

SetCustomerDamagedQuantity sets CustomerDamagedQuantity field to given value.

### HasCustomerDamagedQuantity

`func (o *UnfulfillableQuantity) HasCustomerDamagedQuantity() bool`

HasCustomerDamagedQuantity returns a boolean if a field has been set.

### GetWarehouseDamagedQuantity

`func (o *UnfulfillableQuantity) GetWarehouseDamagedQuantity() int32`

GetWarehouseDamagedQuantity returns the WarehouseDamagedQuantity field if non-nil, zero value otherwise.

### GetWarehouseDamagedQuantityOk

`func (o *UnfulfillableQuantity) GetWarehouseDamagedQuantityOk() (*int32, bool)`

GetWarehouseDamagedQuantityOk returns a tuple with the WarehouseDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseDamagedQuantity

`func (o *UnfulfillableQuantity) SetWarehouseDamagedQuantity(v int32)`

SetWarehouseDamagedQuantity sets WarehouseDamagedQuantity field to given value.

### HasWarehouseDamagedQuantity

`func (o *UnfulfillableQuantity) HasWarehouseDamagedQuantity() bool`

HasWarehouseDamagedQuantity returns a boolean if a field has been set.

### GetDistributorDamagedQuantity

`func (o *UnfulfillableQuantity) GetDistributorDamagedQuantity() int32`

GetDistributorDamagedQuantity returns the DistributorDamagedQuantity field if non-nil, zero value otherwise.

### GetDistributorDamagedQuantityOk

`func (o *UnfulfillableQuantity) GetDistributorDamagedQuantityOk() (*int32, bool)`

GetDistributorDamagedQuantityOk returns a tuple with the DistributorDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributorDamagedQuantity

`func (o *UnfulfillableQuantity) SetDistributorDamagedQuantity(v int32)`

SetDistributorDamagedQuantity sets DistributorDamagedQuantity field to given value.

### HasDistributorDamagedQuantity

`func (o *UnfulfillableQuantity) HasDistributorDamagedQuantity() bool`

HasDistributorDamagedQuantity returns a boolean if a field has been set.

### GetCarrierDamagedQuantity

`func (o *UnfulfillableQuantity) GetCarrierDamagedQuantity() int32`

GetCarrierDamagedQuantity returns the CarrierDamagedQuantity field if non-nil, zero value otherwise.

### GetCarrierDamagedQuantityOk

`func (o *UnfulfillableQuantity) GetCarrierDamagedQuantityOk() (*int32, bool)`

GetCarrierDamagedQuantityOk returns a tuple with the CarrierDamagedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierDamagedQuantity

`func (o *UnfulfillableQuantity) SetCarrierDamagedQuantity(v int32)`

SetCarrierDamagedQuantity sets CarrierDamagedQuantity field to given value.

### HasCarrierDamagedQuantity

`func (o *UnfulfillableQuantity) HasCarrierDamagedQuantity() bool`

HasCarrierDamagedQuantity returns a boolean if a field has been set.

### GetDefectiveQuantity

`func (o *UnfulfillableQuantity) GetDefectiveQuantity() int32`

GetDefectiveQuantity returns the DefectiveQuantity field if non-nil, zero value otherwise.

### GetDefectiveQuantityOk

`func (o *UnfulfillableQuantity) GetDefectiveQuantityOk() (*int32, bool)`

GetDefectiveQuantityOk returns a tuple with the DefectiveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectiveQuantity

`func (o *UnfulfillableQuantity) SetDefectiveQuantity(v int32)`

SetDefectiveQuantity sets DefectiveQuantity field to given value.

### HasDefectiveQuantity

`func (o *UnfulfillableQuantity) HasDefectiveQuantity() bool`

HasDefectiveQuantity returns a boolean if a field has been set.

### GetExpiredQuantity

`func (o *UnfulfillableQuantity) GetExpiredQuantity() int32`

GetExpiredQuantity returns the ExpiredQuantity field if non-nil, zero value otherwise.

### GetExpiredQuantityOk

`func (o *UnfulfillableQuantity) GetExpiredQuantityOk() (*int32, bool)`

GetExpiredQuantityOk returns a tuple with the ExpiredQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredQuantity

`func (o *UnfulfillableQuantity) SetExpiredQuantity(v int32)`

SetExpiredQuantity sets ExpiredQuantity field to given value.

### HasExpiredQuantity

`func (o *UnfulfillableQuantity) HasExpiredQuantity() bool`

HasExpiredQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


