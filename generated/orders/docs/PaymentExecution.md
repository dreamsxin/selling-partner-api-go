# PaymentExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | Pointer to **string** | The payment method used for this payment execution (for example, CashOnDelivery, ConvenienceStore, CreditCard, Invoice, Pix, and so on). | [optional] 
**PaymentAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**AcquirerId** | Pointer to **string** | The unique identifier of the payment processor or acquiring bank that authorizes the payment.   **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the &#x60;paymentMethod&#x60; is &#x60;CreditCard&#x60; or &#x60;Pix&#x60;. | [optional] 
**CardBrand** | Pointer to **string** | The card network or brand used in the payment transaction (for example, Visa or Mastercard).  **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the &#x60;paymentMethod&#x60; is &#x60;CreditCard&#x60;. | [optional] 
**AuthorizationCode** | Pointer to **string** | The unique code that confirms the payment authorization.  **Note**: This attribute is only available for orders in the Brazil (BR) marketplace when the &#x60;paymentMethod&#x60; is &#x60;CreditCard&#x60; or &#x60;Pix&#x60;. | [optional] 

## Methods

### NewPaymentExecution

`func NewPaymentExecution() *PaymentExecution`

NewPaymentExecution instantiates a new PaymentExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentExecutionWithDefaults

`func NewPaymentExecutionWithDefaults() *PaymentExecution`

NewPaymentExecutionWithDefaults instantiates a new PaymentExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *PaymentExecution) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentExecution) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentExecution) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentExecution) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentExecution) GetPaymentAmount() Money`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentExecution) GetPaymentAmountOk() (*Money, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentExecution) SetPaymentAmount(v Money)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentExecution) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetAcquirerId

`func (o *PaymentExecution) GetAcquirerId() string`

GetAcquirerId returns the AcquirerId field if non-nil, zero value otherwise.

### GetAcquirerIdOk

`func (o *PaymentExecution) GetAcquirerIdOk() (*string, bool)`

GetAcquirerIdOk returns a tuple with the AcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerId

`func (o *PaymentExecution) SetAcquirerId(v string)`

SetAcquirerId sets AcquirerId field to given value.

### HasAcquirerId

`func (o *PaymentExecution) HasAcquirerId() bool`

HasAcquirerId returns a boolean if a field has been set.

### GetCardBrand

`func (o *PaymentExecution) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PaymentExecution) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PaymentExecution) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *PaymentExecution) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetAuthorizationCode

`func (o *PaymentExecution) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *PaymentExecution) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *PaymentExecution) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *PaymentExecution) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


