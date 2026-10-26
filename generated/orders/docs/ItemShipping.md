# ItemShipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledDeliveryWindow** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 
**ShippingConstraints** | Pointer to [**ItemShippingConstraints**](ItemShippingConstraints.md) |  | [optional] 
**InternationalShipping** | Pointer to [**ItemInternationalShipping**](ItemInternationalShipping.md) |  | [optional] 

## Methods

### NewItemShipping

`func NewItemShipping() *ItemShipping`

NewItemShipping instantiates a new ItemShipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemShippingWithDefaults

`func NewItemShippingWithDefaults() *ItemShipping`

NewItemShippingWithDefaults instantiates a new ItemShipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledDeliveryWindow

`func (o *ItemShipping) GetScheduledDeliveryWindow() DateTimeRange`

GetScheduledDeliveryWindow returns the ScheduledDeliveryWindow field if non-nil, zero value otherwise.

### GetScheduledDeliveryWindowOk

`func (o *ItemShipping) GetScheduledDeliveryWindowOk() (*DateTimeRange, bool)`

GetScheduledDeliveryWindowOk returns a tuple with the ScheduledDeliveryWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryWindow

`func (o *ItemShipping) SetScheduledDeliveryWindow(v DateTimeRange)`

SetScheduledDeliveryWindow sets ScheduledDeliveryWindow field to given value.

### HasScheduledDeliveryWindow

`func (o *ItemShipping) HasScheduledDeliveryWindow() bool`

HasScheduledDeliveryWindow returns a boolean if a field has been set.

### GetShippingConstraints

`func (o *ItemShipping) GetShippingConstraints() ItemShippingConstraints`

GetShippingConstraints returns the ShippingConstraints field if non-nil, zero value otherwise.

### GetShippingConstraintsOk

`func (o *ItemShipping) GetShippingConstraintsOk() (*ItemShippingConstraints, bool)`

GetShippingConstraintsOk returns a tuple with the ShippingConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingConstraints

`func (o *ItemShipping) SetShippingConstraints(v ItemShippingConstraints)`

SetShippingConstraints sets ShippingConstraints field to given value.

### HasShippingConstraints

`func (o *ItemShipping) HasShippingConstraints() bool`

HasShippingConstraints returns a boolean if a field has been set.

### GetInternationalShipping

`func (o *ItemShipping) GetInternationalShipping() ItemInternationalShipping`

GetInternationalShipping returns the InternationalShipping field if non-nil, zero value otherwise.

### GetInternationalShippingOk

`func (o *ItemShipping) GetInternationalShippingOk() (*ItemInternationalShipping, bool)`

GetInternationalShippingOk returns a tuple with the InternationalShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalShipping

`func (o *ItemShipping) SetInternationalShipping(v ItemInternationalShipping)`

SetInternationalShipping sets InternationalShipping field to given value.

### HasInternationalShipping

`func (o *ItemShipping) HasInternationalShipping() bool`

HasInternationalShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


