# OrderTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxRegistrations** | Pointer to [**[]OrderTaxRegistration**](OrderTaxRegistration.md) | A list of tax registrations associated with the order. | [optional] 
**TaxInvoicing** | Pointer to [**OrderTaxInvoicing**](OrderTaxInvoicing.md) |  | [optional] 

## Methods

### NewOrderTax

`func NewOrderTax() *OrderTax`

NewOrderTax instantiates a new OrderTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaxWithDefaults

`func NewOrderTaxWithDefaults() *OrderTax`

NewOrderTaxWithDefaults instantiates a new OrderTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxRegistrations

`func (o *OrderTax) GetTaxRegistrations() []OrderTaxRegistration`

GetTaxRegistrations returns the TaxRegistrations field if non-nil, zero value otherwise.

### GetTaxRegistrationsOk

`func (o *OrderTax) GetTaxRegistrationsOk() (*[]OrderTaxRegistration, bool)`

GetTaxRegistrationsOk returns a tuple with the TaxRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrations

`func (o *OrderTax) SetTaxRegistrations(v []OrderTaxRegistration)`

SetTaxRegistrations sets TaxRegistrations field to given value.

### HasTaxRegistrations

`func (o *OrderTax) HasTaxRegistrations() bool`

HasTaxRegistrations returns a boolean if a field has been set.

### GetTaxInvoicing

`func (o *OrderTax) GetTaxInvoicing() OrderTaxInvoicing`

GetTaxInvoicing returns the TaxInvoicing field if non-nil, zero value otherwise.

### GetTaxInvoicingOk

`func (o *OrderTax) GetTaxInvoicingOk() (*OrderTaxInvoicing, bool)`

GetTaxInvoicingOk returns a tuple with the TaxInvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxInvoicing

`func (o *OrderTax) SetTaxInvoicing(v OrderTaxInvoicing)`

SetTaxInvoicing sets TaxInvoicing field to given value.

### HasTaxInvoicing

`func (o *OrderTax) HasTaxInvoicing() bool`

HasTaxInvoicing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


