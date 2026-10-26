# OrderPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentExecutions** | Pointer to [**[]PaymentExecution**](PaymentExecution.md) | A list of payment executions for the order. | [optional] 

## Methods

### NewOrderPayment

`func NewOrderPayment() *OrderPayment`

NewOrderPayment instantiates a new OrderPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPaymentWithDefaults

`func NewOrderPaymentWithDefaults() *OrderPayment`

NewOrderPaymentWithDefaults instantiates a new OrderPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentExecutions

`func (o *OrderPayment) GetPaymentExecutions() []PaymentExecution`

GetPaymentExecutions returns the PaymentExecutions field if non-nil, zero value otherwise.

### GetPaymentExecutionsOk

`func (o *OrderPayment) GetPaymentExecutionsOk() (*[]PaymentExecution, bool)`

GetPaymentExecutionsOk returns a tuple with the PaymentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentExecutions

`func (o *OrderPayment) SetPaymentExecutions(v []PaymentExecution)`

SetPaymentExecutions sets PaymentExecutions field to given value.

### HasPaymentExecutions

`func (o *OrderPayment) HasPaymentExecutions() bool`

HasPaymentExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


