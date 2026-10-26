# OrderFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentStatus** | [**FulfillmentStatus**](FulfillmentStatus.md) |  | 
**FulfilledBy** | Pointer to **string** | Specifies whether Amazon or the merchant is responsible for fulfilling this order.  **Possible values**: &#x60;AMAZON&#x60;, &#x60;MERCHANT&#x60;. | [optional] 
**FulfillmentServiceLevel** | Pointer to **string** | The category of the shipping speed option selected by the customer at checkout.  **Possible values**: &#x60;EXPEDITED&#x60;, &#x60;FREE_ECONOMY&#x60;, &#x60;NEXT_DAY&#x60;, &#x60;PRIORITY&#x60;, &#x60;SAME_DAY&#x60;, &#x60;SECOND_DAY&#x60;, &#x60;SCHEDULED&#x60;, &#x60;STANDARD&#x60;. | [optional] 
**ShipByWindow** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 
**DeliverByWindow** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 
**LabelPrintingWindow** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 
**PromiseCalculationInputs** | Pointer to [**PromiseCalculationInputs**](PromiseCalculationInputs.md) |  | [optional] 

## Methods

### NewOrderFulfillment

`func NewOrderFulfillment(fulfillmentStatus FulfillmentStatus, ) *OrderFulfillment`

NewOrderFulfillment instantiates a new OrderFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFulfillmentWithDefaults

`func NewOrderFulfillmentWithDefaults() *OrderFulfillment`

NewOrderFulfillmentWithDefaults instantiates a new OrderFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentStatus

`func (o *OrderFulfillment) GetFulfillmentStatus() FulfillmentStatus`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *OrderFulfillment) GetFulfillmentStatusOk() (*FulfillmentStatus, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *OrderFulfillment) SetFulfillmentStatus(v FulfillmentStatus)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.


### GetFulfilledBy

`func (o *OrderFulfillment) GetFulfilledBy() string`

GetFulfilledBy returns the FulfilledBy field if non-nil, zero value otherwise.

### GetFulfilledByOk

`func (o *OrderFulfillment) GetFulfilledByOk() (*string, bool)`

GetFulfilledByOk returns a tuple with the FulfilledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledBy

`func (o *OrderFulfillment) SetFulfilledBy(v string)`

SetFulfilledBy sets FulfilledBy field to given value.

### HasFulfilledBy

`func (o *OrderFulfillment) HasFulfilledBy() bool`

HasFulfilledBy returns a boolean if a field has been set.

### GetFulfillmentServiceLevel

`func (o *OrderFulfillment) GetFulfillmentServiceLevel() string`

GetFulfillmentServiceLevel returns the FulfillmentServiceLevel field if non-nil, zero value otherwise.

### GetFulfillmentServiceLevelOk

`func (o *OrderFulfillment) GetFulfillmentServiceLevelOk() (*string, bool)`

GetFulfillmentServiceLevelOk returns a tuple with the FulfillmentServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentServiceLevel

`func (o *OrderFulfillment) SetFulfillmentServiceLevel(v string)`

SetFulfillmentServiceLevel sets FulfillmentServiceLevel field to given value.

### HasFulfillmentServiceLevel

`func (o *OrderFulfillment) HasFulfillmentServiceLevel() bool`

HasFulfillmentServiceLevel returns a boolean if a field has been set.

### GetShipByWindow

`func (o *OrderFulfillment) GetShipByWindow() DateTimeRange`

GetShipByWindow returns the ShipByWindow field if non-nil, zero value otherwise.

### GetShipByWindowOk

`func (o *OrderFulfillment) GetShipByWindowOk() (*DateTimeRange, bool)`

GetShipByWindowOk returns a tuple with the ShipByWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByWindow

`func (o *OrderFulfillment) SetShipByWindow(v DateTimeRange)`

SetShipByWindow sets ShipByWindow field to given value.

### HasShipByWindow

`func (o *OrderFulfillment) HasShipByWindow() bool`

HasShipByWindow returns a boolean if a field has been set.

### GetDeliverByWindow

`func (o *OrderFulfillment) GetDeliverByWindow() DateTimeRange`

GetDeliverByWindow returns the DeliverByWindow field if non-nil, zero value otherwise.

### GetDeliverByWindowOk

`func (o *OrderFulfillment) GetDeliverByWindowOk() (*DateTimeRange, bool)`

GetDeliverByWindowOk returns a tuple with the DeliverByWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverByWindow

`func (o *OrderFulfillment) SetDeliverByWindow(v DateTimeRange)`

SetDeliverByWindow sets DeliverByWindow field to given value.

### HasDeliverByWindow

`func (o *OrderFulfillment) HasDeliverByWindow() bool`

HasDeliverByWindow returns a boolean if a field has been set.

### GetLabelPrintingWindow

`func (o *OrderFulfillment) GetLabelPrintingWindow() DateTimeRange`

GetLabelPrintingWindow returns the LabelPrintingWindow field if non-nil, zero value otherwise.

### GetLabelPrintingWindowOk

`func (o *OrderFulfillment) GetLabelPrintingWindowOk() (*DateTimeRange, bool)`

GetLabelPrintingWindowOk returns a tuple with the LabelPrintingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPrintingWindow

`func (o *OrderFulfillment) SetLabelPrintingWindow(v DateTimeRange)`

SetLabelPrintingWindow sets LabelPrintingWindow field to given value.

### HasLabelPrintingWindow

`func (o *OrderFulfillment) HasLabelPrintingWindow() bool`

HasLabelPrintingWindow returns a boolean if a field has been set.

### GetPromiseCalculationInputs

`func (o *OrderFulfillment) GetPromiseCalculationInputs() PromiseCalculationInputs`

GetPromiseCalculationInputs returns the PromiseCalculationInputs field if non-nil, zero value otherwise.

### GetPromiseCalculationInputsOk

`func (o *OrderFulfillment) GetPromiseCalculationInputsOk() (*PromiseCalculationInputs, bool)`

GetPromiseCalculationInputsOk returns a tuple with the PromiseCalculationInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromiseCalculationInputs

`func (o *OrderFulfillment) SetPromiseCalculationInputs(v PromiseCalculationInputs)`

SetPromiseCalculationInputs sets PromiseCalculationInputs field to given value.

### HasPromiseCalculationInputs

`func (o *OrderFulfillment) HasPromiseCalculationInputs() bool`

HasPromiseCalculationInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


