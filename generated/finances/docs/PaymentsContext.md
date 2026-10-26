# PaymentsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | Pointer to **string** | The type of payment. | [optional] 
**PaymentMethod** | Pointer to **string** | The method of payment. | [optional] 
**PaymentReference** | Pointer to **string** | The reference number of the payment. | [optional] 
**PaymentDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 

## Methods

### NewPaymentsContext

`func NewPaymentsContext() *PaymentsContext`

NewPaymentsContext instantiates a new PaymentsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsContextWithDefaults

`func NewPaymentsContextWithDefaults() *PaymentsContext`

NewPaymentsContextWithDefaults instantiates a new PaymentsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *PaymentsContext) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentsContext) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentsContext) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentsContext) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentsContext) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentsContext) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentsContext) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentsContext) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentReference

`func (o *PaymentsContext) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *PaymentsContext) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *PaymentsContext) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *PaymentsContext) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### GetPaymentDate

`func (o *PaymentsContext) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentsContext) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentsContext) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *PaymentsContext) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


