# OrderTaxInvoicing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerInvoicePreference** | Pointer to **string** | The buyer&#39;s invoicing preference, which indicates whether the seller should issue an individual or a business invoice to the buyer.    **Note**: This attribute is only available in the Turkey marketplace.   **Possible values**: - &#x60;INDIVIDUAL&#x60; (Issues an individual invoice to the buyer) - &#x60;BUSINESS&#x60; (Issues a business invoice to the buyer) | [optional] 
**InvoiceStatus** | Pointer to **string** | The status of the invoice. Only available for Easy Ship orders and orders in the Brazil marketplace.  **Possible values**: - &#x60;NOT_REQUIRED&#x60; (The order does not require an electronic invoice to be uploaded) - &#x60;NOT_FOUND&#x60; (The order requires an electronic invoice but it is not uploaded) - &#x60;PROCESSING&#x60; (The required electronic invoice was uploaded and is processing) - &#x60;ERRORED&#x60; (The uploaded electronic invoice was not accepted) - &#x60;ACCEPTED&#x60; (The uploaded electronic invoice was accepted) | [optional] 

## Methods

### NewOrderTaxInvoicing

`func NewOrderTaxInvoicing() *OrderTaxInvoicing`

NewOrderTaxInvoicing instantiates a new OrderTaxInvoicing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaxInvoicingWithDefaults

`func NewOrderTaxInvoicingWithDefaults() *OrderTaxInvoicing`

NewOrderTaxInvoicingWithDefaults instantiates a new OrderTaxInvoicing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerInvoicePreference

`func (o *OrderTaxInvoicing) GetBuyerInvoicePreference() string`

GetBuyerInvoicePreference returns the BuyerInvoicePreference field if non-nil, zero value otherwise.

### GetBuyerInvoicePreferenceOk

`func (o *OrderTaxInvoicing) GetBuyerInvoicePreferenceOk() (*string, bool)`

GetBuyerInvoicePreferenceOk returns a tuple with the BuyerInvoicePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerInvoicePreference

`func (o *OrderTaxInvoicing) SetBuyerInvoicePreference(v string)`

SetBuyerInvoicePreference sets BuyerInvoicePreference field to given value.

### HasBuyerInvoicePreference

`func (o *OrderTaxInvoicing) HasBuyerInvoicePreference() bool`

HasBuyerInvoicePreference returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *OrderTaxInvoicing) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *OrderTaxInvoicing) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *OrderTaxInvoicing) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *OrderTaxInvoicing) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


