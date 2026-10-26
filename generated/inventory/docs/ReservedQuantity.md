# ReservedQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalReservedQuantity** | Pointer to **int32** | The total number of units in Amazon&#39;s fulfillment network that are currently being picked, packed, and shipped; or are sidelined for measurement, sampling, or other internal processes. | [optional] 
**PendingCustomerOrderQuantity** | Pointer to **int32** | The number of units reserved for customer orders. | [optional] 
**PendingTransshipmentQuantity** | Pointer to **int32** | The number of units being transferred from one fulfillment center to another. | [optional] 
**FcProcessingQuantity** | Pointer to **int32** | The number of units that have been sidelined at the fulfillment center for additional processing. | [optional] 

## Methods

### NewReservedQuantity

`func NewReservedQuantity() *ReservedQuantity`

NewReservedQuantity instantiates a new ReservedQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservedQuantityWithDefaults

`func NewReservedQuantityWithDefaults() *ReservedQuantity`

NewReservedQuantityWithDefaults instantiates a new ReservedQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalReservedQuantity

`func (o *ReservedQuantity) GetTotalReservedQuantity() int32`

GetTotalReservedQuantity returns the TotalReservedQuantity field if non-nil, zero value otherwise.

### GetTotalReservedQuantityOk

`func (o *ReservedQuantity) GetTotalReservedQuantityOk() (*int32, bool)`

GetTotalReservedQuantityOk returns a tuple with the TotalReservedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservedQuantity

`func (o *ReservedQuantity) SetTotalReservedQuantity(v int32)`

SetTotalReservedQuantity sets TotalReservedQuantity field to given value.

### HasTotalReservedQuantity

`func (o *ReservedQuantity) HasTotalReservedQuantity() bool`

HasTotalReservedQuantity returns a boolean if a field has been set.

### GetPendingCustomerOrderQuantity

`func (o *ReservedQuantity) GetPendingCustomerOrderQuantity() int32`

GetPendingCustomerOrderQuantity returns the PendingCustomerOrderQuantity field if non-nil, zero value otherwise.

### GetPendingCustomerOrderQuantityOk

`func (o *ReservedQuantity) GetPendingCustomerOrderQuantityOk() (*int32, bool)`

GetPendingCustomerOrderQuantityOk returns a tuple with the PendingCustomerOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCustomerOrderQuantity

`func (o *ReservedQuantity) SetPendingCustomerOrderQuantity(v int32)`

SetPendingCustomerOrderQuantity sets PendingCustomerOrderQuantity field to given value.

### HasPendingCustomerOrderQuantity

`func (o *ReservedQuantity) HasPendingCustomerOrderQuantity() bool`

HasPendingCustomerOrderQuantity returns a boolean if a field has been set.

### GetPendingTransshipmentQuantity

`func (o *ReservedQuantity) GetPendingTransshipmentQuantity() int32`

GetPendingTransshipmentQuantity returns the PendingTransshipmentQuantity field if non-nil, zero value otherwise.

### GetPendingTransshipmentQuantityOk

`func (o *ReservedQuantity) GetPendingTransshipmentQuantityOk() (*int32, bool)`

GetPendingTransshipmentQuantityOk returns a tuple with the PendingTransshipmentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTransshipmentQuantity

`func (o *ReservedQuantity) SetPendingTransshipmentQuantity(v int32)`

SetPendingTransshipmentQuantity sets PendingTransshipmentQuantity field to given value.

### HasPendingTransshipmentQuantity

`func (o *ReservedQuantity) HasPendingTransshipmentQuantity() bool`

HasPendingTransshipmentQuantity returns a boolean if a field has been set.

### GetFcProcessingQuantity

`func (o *ReservedQuantity) GetFcProcessingQuantity() int32`

GetFcProcessingQuantity returns the FcProcessingQuantity field if non-nil, zero value otherwise.

### GetFcProcessingQuantityOk

`func (o *ReservedQuantity) GetFcProcessingQuantityOk() (*int32, bool)`

GetFcProcessingQuantityOk returns a tuple with the FcProcessingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcProcessingQuantity

`func (o *ReservedQuantity) SetFcProcessingQuantity(v int32)`

SetFcProcessingQuantity sets FcProcessingQuantity field to given value.

### HasFcProcessingQuantity

`func (o *ReservedQuantity) HasFcProcessingQuantity() bool`

HasFcProcessingQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


