# OrderChangeTypeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderChangeTypes** | Pointer to [**[]OrderChangeTypeEnum**](OrderChangeTypeEnum.md) | A list of order change types you can subscribe to (for example, &#x60;BuyerRequestedChange&#x60;). To receive notifications of all change types, do not provide this list. | [optional] 

## Methods

### NewOrderChangeTypeFilter

`func NewOrderChangeTypeFilter() *OrderChangeTypeFilter`

NewOrderChangeTypeFilter instantiates a new OrderChangeTypeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderChangeTypeFilterWithDefaults

`func NewOrderChangeTypeFilterWithDefaults() *OrderChangeTypeFilter`

NewOrderChangeTypeFilterWithDefaults instantiates a new OrderChangeTypeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderChangeTypes

`func (o *OrderChangeTypeFilter) GetOrderChangeTypes() []OrderChangeTypeEnum`

GetOrderChangeTypes returns the OrderChangeTypes field if non-nil, zero value otherwise.

### GetOrderChangeTypesOk

`func (o *OrderChangeTypeFilter) GetOrderChangeTypesOk() (*[]OrderChangeTypeEnum, bool)`

GetOrderChangeTypesOk returns a tuple with the OrderChangeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderChangeTypes

`func (o *OrderChangeTypeFilter) SetOrderChangeTypes(v []OrderChangeTypeEnum)`

SetOrderChangeTypes sets OrderChangeTypes field to given value.

### HasOrderChangeTypes

`func (o *OrderChangeTypeFilter) HasOrderChangeTypes() bool`

HasOrderChangeTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


