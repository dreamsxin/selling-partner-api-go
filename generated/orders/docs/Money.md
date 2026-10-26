# Money

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | A decimal number with no loss of precision. Follows RFC 7159 for number representation. | 
**CurrencyCode** | **string** | The three-letter currency code that identifies the currency type, following ISO 4217 international standards. | 

## Methods

### NewMoney

`func NewMoney(amount string, currencyCode string, ) *Money`

NewMoney instantiates a new Money object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyWithDefaults

`func NewMoneyWithDefaults() *Money`

NewMoneyWithDefaults instantiates a new Money object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Money) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Money) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Money) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *Money) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Money) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Money) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


