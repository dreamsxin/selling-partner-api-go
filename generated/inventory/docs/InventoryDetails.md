# InventoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillableQuantity** | Pointer to **int32** | The item quantity that can be picked, packed, and shipped. | [optional] 
**InboundWorkingQuantity** | Pointer to **int32** | The number of units in an inbound shipment for which you have notified Amazon. | [optional] 
**InboundShippedQuantity** | Pointer to **int32** | The number of units in an inbound shipment that you have notified Amazon about and have provided a tracking number. | [optional] 
**InboundReceivingQuantity** | Pointer to **int32** | The number of units that have not yet been received at an Amazon fulfillment center for processing, but are part of an inbound shipment with some units that have already been received and processed. | [optional] 
**ReservedQuantity** | Pointer to [**ReservedQuantity**](ReservedQuantity.md) |  | [optional] 
**ResearchingQuantity** | Pointer to [**ResearchingQuantity**](ResearchingQuantity.md) |  | [optional] 
**UnfulfillableQuantity** | Pointer to [**UnfulfillableQuantity**](UnfulfillableQuantity.md) |  | [optional] 

## Methods

### NewInventoryDetails

`func NewInventoryDetails() *InventoryDetails`

NewInventoryDetails instantiates a new InventoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDetailsWithDefaults

`func NewInventoryDetailsWithDefaults() *InventoryDetails`

NewInventoryDetailsWithDefaults instantiates a new InventoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillableQuantity

`func (o *InventoryDetails) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *InventoryDetails) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *InventoryDetails) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *InventoryDetails) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetInboundWorkingQuantity

`func (o *InventoryDetails) GetInboundWorkingQuantity() int32`

GetInboundWorkingQuantity returns the InboundWorkingQuantity field if non-nil, zero value otherwise.

### GetInboundWorkingQuantityOk

`func (o *InventoryDetails) GetInboundWorkingQuantityOk() (*int32, bool)`

GetInboundWorkingQuantityOk returns a tuple with the InboundWorkingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundWorkingQuantity

`func (o *InventoryDetails) SetInboundWorkingQuantity(v int32)`

SetInboundWorkingQuantity sets InboundWorkingQuantity field to given value.

### HasInboundWorkingQuantity

`func (o *InventoryDetails) HasInboundWorkingQuantity() bool`

HasInboundWorkingQuantity returns a boolean if a field has been set.

### GetInboundShippedQuantity

`func (o *InventoryDetails) GetInboundShippedQuantity() int32`

GetInboundShippedQuantity returns the InboundShippedQuantity field if non-nil, zero value otherwise.

### GetInboundShippedQuantityOk

`func (o *InventoryDetails) GetInboundShippedQuantityOk() (*int32, bool)`

GetInboundShippedQuantityOk returns a tuple with the InboundShippedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundShippedQuantity

`func (o *InventoryDetails) SetInboundShippedQuantity(v int32)`

SetInboundShippedQuantity sets InboundShippedQuantity field to given value.

### HasInboundShippedQuantity

`func (o *InventoryDetails) HasInboundShippedQuantity() bool`

HasInboundShippedQuantity returns a boolean if a field has been set.

### GetInboundReceivingQuantity

`func (o *InventoryDetails) GetInboundReceivingQuantity() int32`

GetInboundReceivingQuantity returns the InboundReceivingQuantity field if non-nil, zero value otherwise.

### GetInboundReceivingQuantityOk

`func (o *InventoryDetails) GetInboundReceivingQuantityOk() (*int32, bool)`

GetInboundReceivingQuantityOk returns a tuple with the InboundReceivingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundReceivingQuantity

`func (o *InventoryDetails) SetInboundReceivingQuantity(v int32)`

SetInboundReceivingQuantity sets InboundReceivingQuantity field to given value.

### HasInboundReceivingQuantity

`func (o *InventoryDetails) HasInboundReceivingQuantity() bool`

HasInboundReceivingQuantity returns a boolean if a field has been set.

### GetReservedQuantity

`func (o *InventoryDetails) GetReservedQuantity() ReservedQuantity`

GetReservedQuantity returns the ReservedQuantity field if non-nil, zero value otherwise.

### GetReservedQuantityOk

`func (o *InventoryDetails) GetReservedQuantityOk() (*ReservedQuantity, bool)`

GetReservedQuantityOk returns a tuple with the ReservedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedQuantity

`func (o *InventoryDetails) SetReservedQuantity(v ReservedQuantity)`

SetReservedQuantity sets ReservedQuantity field to given value.

### HasReservedQuantity

`func (o *InventoryDetails) HasReservedQuantity() bool`

HasReservedQuantity returns a boolean if a field has been set.

### GetResearchingQuantity

`func (o *InventoryDetails) GetResearchingQuantity() ResearchingQuantity`

GetResearchingQuantity returns the ResearchingQuantity field if non-nil, zero value otherwise.

### GetResearchingQuantityOk

`func (o *InventoryDetails) GetResearchingQuantityOk() (*ResearchingQuantity, bool)`

GetResearchingQuantityOk returns a tuple with the ResearchingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchingQuantity

`func (o *InventoryDetails) SetResearchingQuantity(v ResearchingQuantity)`

SetResearchingQuantity sets ResearchingQuantity field to given value.

### HasResearchingQuantity

`func (o *InventoryDetails) HasResearchingQuantity() bool`

HasResearchingQuantity returns a boolean if a field has been set.

### GetUnfulfillableQuantity

`func (o *InventoryDetails) GetUnfulfillableQuantity() UnfulfillableQuantity`

GetUnfulfillableQuantity returns the UnfulfillableQuantity field if non-nil, zero value otherwise.

### GetUnfulfillableQuantityOk

`func (o *InventoryDetails) GetUnfulfillableQuantityOk() (*UnfulfillableQuantity, bool)`

GetUnfulfillableQuantityOk returns a tuple with the UnfulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfulfillableQuantity

`func (o *InventoryDetails) SetUnfulfillableQuantity(v UnfulfillableQuantity)`

SetUnfulfillableQuantity sets UnfulfillableQuantity field to given value.

### HasUnfulfillableQuantity

`func (o *InventoryDetails) HasUnfulfillableQuantity() bool`

HasUnfulfillableQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


