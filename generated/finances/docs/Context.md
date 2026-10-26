# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreName** | Pointer to **string** | The store name associated with the transaction. | [optional] 
**OrderType** | Pointer to **string** | The order type of the transaction. | [optional] 
**Channel** | Pointer to **string** | The channel of the transaction. For example: &#x60;MFN&#x60;. | [optional] 
**Asin** | Pointer to **string** | The Amazon Standard Identification Number (ASIN) of the item. | [optional] 
**Sku** | Pointer to **string** | The Stock Keeping Unit (SKU) of the item. | [optional] 
**QuantityShipped** | Pointer to **int32** | The quantity of the item shipped. | [optional] 
**FulfillmentNetwork** | Pointer to **string** | The fulfillment network of the item. | [optional] 
**PaymentType** | Pointer to **string** | The type of payment. | [optional] 
**PaymentMethod** | Pointer to **string** | The method of payment. | [optional] 
**PaymentReference** | Pointer to **string** | The reference number of the payment. | [optional] 
**PaymentDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**DeferralReason** | Pointer to **string** | The deferral policy applied to the transaction.  **Examples:** &#x60;B2B&#x60; (invoiced orders), &#x60;DD7&#x60; (delivery date policy) | [optional] 
**MaturityDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**StartTime** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**EndTime** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**ContextType** | **string** | The type of context. | 

## Methods

### NewContext

`func NewContext(contextType string, ) *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreName

`func (o *Context) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *Context) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *Context) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *Context) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetOrderType

`func (o *Context) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *Context) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *Context) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *Context) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetChannel

`func (o *Context) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Context) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Context) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Context) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetAsin

`func (o *Context) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *Context) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *Context) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *Context) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetSku

`func (o *Context) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Context) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Context) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Context) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetQuantityShipped

`func (o *Context) GetQuantityShipped() int32`

GetQuantityShipped returns the QuantityShipped field if non-nil, zero value otherwise.

### GetQuantityShippedOk

`func (o *Context) GetQuantityShippedOk() (*int32, bool)`

GetQuantityShippedOk returns a tuple with the QuantityShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityShipped

`func (o *Context) SetQuantityShipped(v int32)`

SetQuantityShipped sets QuantityShipped field to given value.

### HasQuantityShipped

`func (o *Context) HasQuantityShipped() bool`

HasQuantityShipped returns a boolean if a field has been set.

### GetFulfillmentNetwork

`func (o *Context) GetFulfillmentNetwork() string`

GetFulfillmentNetwork returns the FulfillmentNetwork field if non-nil, zero value otherwise.

### GetFulfillmentNetworkOk

`func (o *Context) GetFulfillmentNetworkOk() (*string, bool)`

GetFulfillmentNetworkOk returns a tuple with the FulfillmentNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentNetwork

`func (o *Context) SetFulfillmentNetwork(v string)`

SetFulfillmentNetwork sets FulfillmentNetwork field to given value.

### HasFulfillmentNetwork

`func (o *Context) HasFulfillmentNetwork() bool`

HasFulfillmentNetwork returns a boolean if a field has been set.

### GetPaymentType

`func (o *Context) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *Context) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *Context) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *Context) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Context) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Context) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Context) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Context) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentReference

`func (o *Context) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *Context) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *Context) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *Context) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### GetPaymentDate

`func (o *Context) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Context) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Context) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *Context) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetDeferralReason

`func (o *Context) GetDeferralReason() string`

GetDeferralReason returns the DeferralReason field if non-nil, zero value otherwise.

### GetDeferralReasonOk

`func (o *Context) GetDeferralReasonOk() (*string, bool)`

GetDeferralReasonOk returns a tuple with the DeferralReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralReason

`func (o *Context) SetDeferralReason(v string)`

SetDeferralReason sets DeferralReason field to given value.

### HasDeferralReason

`func (o *Context) HasDeferralReason() bool`

HasDeferralReason returns a boolean if a field has been set.

### GetMaturityDate

`func (o *Context) GetMaturityDate() time.Time`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *Context) GetMaturityDateOk() (*time.Time, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *Context) SetMaturityDate(v time.Time)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *Context) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetStartTime

`func (o *Context) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Context) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Context) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Context) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Context) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Context) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Context) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Context) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetContextType

`func (o *Context) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *Context) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *Context) SetContextType(v string)`

SetContextType sets ContextType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


