# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | An Amazon-defined order identifier. | 
**OrderAliases** | Pointer to [**[]Alias**](Alias.md) | Alternative identifiers that can be used to reference this order, such as seller-defined order numbers. | [optional] 
**CreatedTime** | **time.Time** | The time when the customer placed the order. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | 
**LastUpdatedTime** | **time.Time** | The most recent time when any aspect of this order was modified by Amazon or the seller. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. | 
**Programs** | Pointer to **[]string** | Special programs associated with this order that may affect fulfillment or customer experience.   **Possible values**: &#x60;AMAZON_BAZAAR&#x60;, &#x60;AMAZON_BUSINESS&#x60;, &#x60;AMAZON_EASY_SHIP&#x60;, &#x60;AMAZON_HAUL&#x60;, &#x60;DELIVERY_BY_AMAZON&#x60;, &#x60;FBM_SHIP_PLUS&#x60;, &#x60;INVOICE_BY_AMAZON&#x60;, &#x60;IN_STORE_PICK_UP&#x60;, &#x60;PREMIUM&#x60;, &#x60;PREORDER&#x60;, &#x60;PRIME&#x60; | [optional] 
**AssociatedOrders** | Pointer to [**[]AssociatedOrder**](AssociatedOrder.md) | Other orders that have a direct relationship to this order, such as replacement or exchange orders. | [optional] 
**SalesChannel** | [**SalesChannel**](SalesChannel.md) |  | 
**Buyer** | Pointer to [**Buyer**](Buyer.md) |  | [optional] 
**Recipient** | Pointer to [**Recipient**](Recipient.md) |  | [optional] 
**Proceeds** | Pointer to [**OrderProceeds**](OrderProceeds.md) |  | [optional] 
**Payment** | Pointer to [**OrderPayment**](OrderPayment.md) |  | [optional] 
**Tax** | Pointer to [**OrderTax**](OrderTax.md) |  | [optional] 
**Fulfillment** | Pointer to [**OrderFulfillment**](OrderFulfillment.md) |  | [optional] 
**OrderItems** | [**[]OrderItem**](OrderItem.md) | The list of all order items included in this order. | 
**Packages** | Pointer to [**[]OrderPackage**](OrderPackage.md) | Shipping packages created for this order, including tracking information. **Note:** Only available for merchant-fulfilled (FBM) orders. | [optional] 
**FulfillmentOrders** | Pointer to [**[]FulfillmentOrder**](FulfillmentOrder.md) | The list of fulfillment orders associated with this customer order. Each entry corresponds to one fulfillment unit created by Amazon for this order. **Note:** Only available for Easy Ship and Seller Flex. | [optional] 

## Methods

### NewOrder

`func NewOrder(orderId string, createdTime time.Time, lastUpdatedTime time.Time, salesChannel SalesChannel, orderItems []OrderItem, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Order) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetOrderAliases

`func (o *Order) GetOrderAliases() []Alias`

GetOrderAliases returns the OrderAliases field if non-nil, zero value otherwise.

### GetOrderAliasesOk

`func (o *Order) GetOrderAliasesOk() (*[]Alias, bool)`

GetOrderAliasesOk returns a tuple with the OrderAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAliases

`func (o *Order) SetOrderAliases(v []Alias)`

SetOrderAliases sets OrderAliases field to given value.

### HasOrderAliases

`func (o *Order) HasOrderAliases() bool`

HasOrderAliases returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Order) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Order) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Order) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastUpdatedTime

`func (o *Order) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Order) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Order) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPrograms

`func (o *Order) GetPrograms() []string`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *Order) GetProgramsOk() (*[]string, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *Order) SetPrograms(v []string)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *Order) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### GetAssociatedOrders

`func (o *Order) GetAssociatedOrders() []AssociatedOrder`

GetAssociatedOrders returns the AssociatedOrders field if non-nil, zero value otherwise.

### GetAssociatedOrdersOk

`func (o *Order) GetAssociatedOrdersOk() (*[]AssociatedOrder, bool)`

GetAssociatedOrdersOk returns a tuple with the AssociatedOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedOrders

`func (o *Order) SetAssociatedOrders(v []AssociatedOrder)`

SetAssociatedOrders sets AssociatedOrders field to given value.

### HasAssociatedOrders

`func (o *Order) HasAssociatedOrders() bool`

HasAssociatedOrders returns a boolean if a field has been set.

### GetSalesChannel

`func (o *Order) GetSalesChannel() SalesChannel`

GetSalesChannel returns the SalesChannel field if non-nil, zero value otherwise.

### GetSalesChannelOk

`func (o *Order) GetSalesChannelOk() (*SalesChannel, bool)`

GetSalesChannelOk returns a tuple with the SalesChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannel

`func (o *Order) SetSalesChannel(v SalesChannel)`

SetSalesChannel sets SalesChannel field to given value.


### GetBuyer

`func (o *Order) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Order) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Order) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Order) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetRecipient

`func (o *Order) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Order) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Order) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Order) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetProceeds

`func (o *Order) GetProceeds() OrderProceeds`

GetProceeds returns the Proceeds field if non-nil, zero value otherwise.

### GetProceedsOk

`func (o *Order) GetProceedsOk() (*OrderProceeds, bool)`

GetProceedsOk returns a tuple with the Proceeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceeds

`func (o *Order) SetProceeds(v OrderProceeds)`

SetProceeds sets Proceeds field to given value.

### HasProceeds

`func (o *Order) HasProceeds() bool`

HasProceeds returns a boolean if a field has been set.

### GetPayment

`func (o *Order) GetPayment() OrderPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Order) GetPaymentOk() (*OrderPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Order) SetPayment(v OrderPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Order) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetTax

`func (o *Order) GetTax() OrderTax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *Order) GetTaxOk() (*OrderTax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *Order) SetTax(v OrderTax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *Order) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetFulfillment

`func (o *Order) GetFulfillment() OrderFulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *Order) GetFulfillmentOk() (*OrderFulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *Order) SetFulfillment(v OrderFulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *Order) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetOrderItems

`func (o *Order) GetOrderItems() []OrderItem`

GetOrderItems returns the OrderItems field if non-nil, zero value otherwise.

### GetOrderItemsOk

`func (o *Order) GetOrderItemsOk() (*[]OrderItem, bool)`

GetOrderItemsOk returns a tuple with the OrderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItems

`func (o *Order) SetOrderItems(v []OrderItem)`

SetOrderItems sets OrderItems field to given value.


### GetPackages

`func (o *Order) GetPackages() []OrderPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *Order) GetPackagesOk() (*[]OrderPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *Order) SetPackages(v []OrderPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *Order) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetFulfillmentOrders

`func (o *Order) GetFulfillmentOrders() []FulfillmentOrder`

GetFulfillmentOrders returns the FulfillmentOrders field if non-nil, zero value otherwise.

### GetFulfillmentOrdersOk

`func (o *Order) GetFulfillmentOrdersOk() (*[]FulfillmentOrder, bool)`

GetFulfillmentOrdersOk returns a tuple with the FulfillmentOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentOrders

`func (o *Order) SetFulfillmentOrders(v []FulfillmentOrder)`

SetFulfillmentOrders sets FulfillmentOrders field to given value.

### HasFulfillmentOrders

`func (o *Order) HasFulfillmentOrders() bool`

HasFulfillmentOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


